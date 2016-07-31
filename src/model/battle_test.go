package model

import (
	"math/rand"
	"testing"
)

type testPokemon struct {
	level   int
	species POKEMON
}

func TestHandleMoveSingleBattle(t *testing.T) {
	for i, tt := range []struct {
		source testPokemon
		target testPokemon
		move   *Move
		logs   []BattleLog
	}{
		{
			source: testPokemon{10, BULBASAUR},
			target: testPokemon{10, CHARMANDER},
			move:   &Move{Data: GetMoveDataByID(TACKLE)},
			logs:   []BattleLog{DamageLog{1, 8}},
		},
		{
			source: testPokemon{10, CHARMANDER},
			target: testPokemon{10, BULBASAUR},
			move:   &Move{Data: GetMoveDataByID(EMBER)},
			logs:   []BattleLog{DamageLog{1, 20}},
		},
		{
			source: testPokemon{100, PONYTA},
			target: testPokemon{10, CHARMANDER},
			move:   &Move{Data: GetMoveDataByID(FLAMETHROWER)},
			logs:   []BattleLog{DamageLog{1, 4766}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   &Move{Data: GetMoveDataByID(GROWL)},
			logs:   []BattleLog{StatStageChangeLog{1, Stats{attack: -1}}},
		},
	} {
		rand.Seed(42)
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		source.Moves[0] = tt.move
		attemptedMove := attemptedMove{
			Source:      source,
			SourceIndex: 0,
			TargetIndex: 1,
			Move:        tt.move,
		}
		battle := NewSingleBattle(source, target)
		battle.HandleMove(attemptedMove)

		got, want := filterLogs(battle.Logs), tt.logs
		if len(got) != len(want) {
			t.Fatalf("%d: got %v want %v", i, got, want)
		}
		for j, g := range got {
			if g != want[j] {
				t.Errorf("%d: got %+v want %+v", i, g, want[j])
			}
		}

		// NOTE: REplays, so effects are applied twice!
		//t.Log(battle.String())
	}
}

func filterLogs(logs []BattleLog) []BattleLog {
	l := []BattleLog{}
	for _, log := range logs {
		switch log.(type) {
		case TextLog:
			continue
		default:
			l = append(l, log)
		}
	}
	return l
}
