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
		logs   []BattleLog
	}{
		{
			source: testPokemon{10, BULBASAUR},
			target: testPokemon{10, CHARMANDER},
			move:   TACKLE,
			logs:   []BattleLog{DamageLog{index: 1}},
		},
		{
			source: testPokemon{10, CHARMANDER},
			target: testPokemon{10, BULBASAUR},
			move:   EMBER,
			logs:   []BattleLog{DamageLog{index: 1}},
		},
		{
			source: testPokemon{100, PONYTA},
			target: testPokemon{10, CHARMANDER},
			move:   FLAMETHROWER,
			logs:   []BattleLog{DamageLog{index: 1}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   GROWL,
			logs:   []BattleLog{StatStageChangeLog{1, Stats{attack: -1}}},
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

		evaluateLogs(t, i, battle.Logs()[1], tt.logs)
		t.Log(LogsToString(battle.Logs()))
	}
}

func TestMovePriority(t *testing.T) {
	for i, tt := range []struct {
		source     testPokemon
		target     testPokemon
		sourceMove MOVE
		targetMove MOVE
		logs       []BattleLog
	}{
		{
			source:     testPokemon{10, PIKACHU},
			target:     testPokemon{10, RATTATA},
			sourceMove: SPARK,
			targetMove: QUICKATTACK,
			logs:       []BattleLog{DamageLog{index: 0}, DamageLog{index: 1}},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		source.Moves[0] = &Move{Data: GetMoveDataByID(tt.sourceMove)}
		target.Moves[0] = &Move{Data: GetMoveDataByID(tt.targetMove)}
		battle := NewSingleBattle(source, target)
		HandleTurn(battle, []Command{{0, 1, 0}, {1, 0, 0}})

		evaluateLogs(t, i, battle.Logs()[1], tt.logs)
		// NOTE: REplays, so effects are applied twice!
		t.Log(LogsToString(battle.Logs()))
	}
}

func evaluateLogs(t *testing.T, n int, gotLogs, wantLogs []BattleLog) {
	got, want := filterLogs(gotLogs), wantLogs
	if len(got) != len(want) {
		t.Fatalf("%d: got %v want %v", n, got, want)
	}
	for i, g := range got {
		wantLog := want[i]
		switch gotLog := g.(type) {
		case DamageLog:
			wantDamageLog, ok := wantLog.(DamageLog)
			if !ok {
				t.Errorf("%d: got %+v want %+v", n, gotLog, wantLog)
			}
			if gotLog.index != wantDamageLog.index {
				t.Errorf("%d: got %+v want %+v (exact damage doesn't matter)", n, gotLog, wantLog)
			}
		default:
			if gotLog != wantLog {
				t.Errorf("%d: got %+v want %+v", n, gotLog, wantLog)
			}
		}
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
