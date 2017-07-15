package model

import (
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
		move   MOVE
		logs   []battleLog
	}{
		{
			source: testPokemon{10, BULBASAUR},
			target: testPokemon{10, CHARMANDER},
			move:   TACKLE,
			logs:   []battleLog{damageLog{Index: 1}},
		},
		{
			source: testPokemon{10, CHARMANDER},
			target: testPokemon{10, BULBASAUR},
			move:   EMBER,
			logs:   []battleLog{damageLog{Index: 1}},
		},
		{
			source: testPokemon{100, PONYTA},
			target: testPokemon{10, CHARMANDER},
			move:   FLAMETHROWER,
			logs:   []battleLog{damageLog{Index: 1}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   GROWL,
			logs:   []battleLog{statStageChangeLog{Index: 1, Changes: Stats{attack: -1}}},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		source.Moves[0] = &Move{Data: GetMoveDataByID(tt.move)}
		attemptedMove := attemptedMove{
			Source:      source,
			SourceIndex: 0,
			TargetIndex: 1,
			Move:        source.Moves[0],
		}
		battle := NewSingleBattle(source, target)
		HandleMove(battle, attemptedMove)

		evaluateLogs(t, i, battle.Log().Logs()[1], tt.logs)
	}
}

func TestMovePriority(t *testing.T) {
	for i, tt := range []struct {
		source     testPokemon
		target     testPokemon
		sourceMove MOVE
		targetMove MOVE
		logs       []battleLog
	}{
		{
			source:     testPokemon{10, PIKACHU},
			target:     testPokemon{10, RATTATA},
			sourceMove: SPARK,
			targetMove: QUICKATTACK,
			logs:       []battleLog{damageLog{Index: 0}, damageLog{Index: 1}},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		source.Moves[0] = &Move{Data: GetMoveDataByID(tt.sourceMove)}
		target.Moves[0] = &Move{Data: GetMoveDataByID(tt.targetMove)}
		battle := NewSingleBattle(source, target)
		HandleTurn(battle, []Command{{0, 1, 0}, {1, 0, 0}})

		evaluateLogs(t, i, battle.Log().Logs()[1], tt.logs)
	}
}

func evaluateLogs(t *testing.T, n int, gotLogs, wantLogs []battleLog) {
	got, want := filterLogs(gotLogs), wantLogs
	if len(got) != len(want) {
		t.Fatalf("%d: got %v want %v", n, got, want)
	}
	for i, g := range got {
		wantLog := want[i]
		switch gotLog := g.(type) {
		case damageLog:
			wantDamageLog, ok := wantLog.(damageLog)
			if !ok {
				t.Errorf("%d: got %+v want %+v", n, gotLog, wantLog)
			}
			if gotLog.Index != wantDamageLog.Index {
				t.Errorf("%d: got %+v want %+v (exact damage doesn't matter)", n, gotLog, wantLog)
			}
		default:
			if gotLog != wantLog {
				t.Errorf("%d: got %+v want %+v", n, gotLog, wantLog)
			}
		}
	}
}

func filterLogs(logs []battleLog) []battleLog {
	l := []battleLog{}
	for _, log := range logs {
		switch log.(type) {
		case textLog:
			continue
		default:
			l = append(l, log)
		}
	}
	return l
}
