package model

import (
	"fmt"
	"testing"
)

type testPokemon struct {
	level   int
	species pokemon
}

// test whether battle flow through a single move is correct
// for more specific move tests, check move_test.go
func TestHandleMoveSingleBattle(t *testing.T) {
	for i, tt := range []struct {
		source testPokemon
		target testPokemon
		move   move
		logs   []battleLog
	}{
		{
			source: testPokemon{10, BULBASAUR},
			target: testPokemon{10, CHARMANDER},
			move:   TACKLE,
			logs:   []battleLog{DamageLog{Index: 1}},
		},
		{
			source: testPokemon{10, SQUIRTLE},
			target: testPokemon{10, CHARMANDER},
			move:   WATERGUN,
			logs:   []battleLog{DamageLog{Index: 1}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   GROWL,
			logs:   []battleLog{StatStageLog{Index: 1, Changes: Stats{attack: -1}}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, SQUIRTLE},
			move:   FLAMECHARGE,
			logs: []battleLog{
				DamageLog{Index: 1},
				StatStageLog{Index: 0, Changes: Stats{speed: +1}},
			},
		},
		{
			source: testPokemon{5, PIKACHU},
			target: testPokemon{5, RATTATA},
			move:   THUNDERWAVE,
			logs: []battleLog{
				NonVolatileConditionLog{Index: 1, Condition: Paralysis{}},
			},
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

		evaluateNonTextLogs(t, battle.Log().Logs()[1], tt.logs, fmt.Sprintf("%d)", i))
	}
}

func TestMovePriority(t *testing.T) {
	for i, tt := range []struct {
		source     testPokemon
		target     testPokemon
		sourceMove move
		targetMove move
		logs       []battleLog
	}{
		{
			source:     testPokemon{10, RATTATA},
			target:     testPokemon{10, RATTATA},
			sourceMove: TACKLE,
			targetMove: QUICKATTACK,
			logs:       []battleLog{DamageLog{Index: 0}, DamageLog{Index: 1}},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		source.Moves[0] = &Move{Data: GetMoveDataByID(tt.sourceMove)}
		target.Moves[0] = &Move{Data: GetMoveDataByID(tt.targetMove)}
		battle := NewSingleBattle(source, target)
		HandleTurn(battle, []Command{{0, 1, 0}, {1, 0, 0}})

		evaluateNonTextLogs(t, battle.Log().Logs()[1], tt.logs, fmt.Sprintf("%d)", i))
	}
}

func TestExactDamage(t *testing.T) {
	for i, tt := range []struct {
		source testPokemon
		target testPokemon
		move   move
		damage int
	}{
		{
			source: testPokemon{10, MAGNEMITE},
			target: testPokemon{10, RATTATA},
			move:   SONICBOOM,
			damage: 20,
		},
		{
			source: testPokemon{10, MAGNEMITE},
			target: testPokemon{10, GASTLY},
			move:   SONICBOOM,
			damage: 0,
		},
		{
			source: testPokemon{10, HAUNTER},
			target: testPokemon{12, GOLEM},
			move:   NIGHTSHADE,
			damage: 10,
		},
		{
			source: testPokemon{10, HAUNTER},
			target: testPokemon{12, RATTATA},
			move:   NIGHTSHADE,
			damage: 0,
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		md := GetMoveDataByID(tt.move)
		md.Accuracy = 100 // we only want to test cases that hit
		source.Moves[0] = &Move{Data: md}
		attemptedMove := attemptedMove{
			Source:      source,
			SourceIndex: 0,
			TargetIndex: 1,
			Move:        source.Moves[0],
		}
		battle := NewSingleBattle(source, target)
		HandleMove(battle, attemptedMove)

		log := battle.Log().Logs()[1][1]
		dLog, ok := log.(DamageLog)
		if !ok {
			if tt.damage == 0 {
				continue
			}
			t.Fatalf("%d): expected a DamageLog got %v", i, log)
		}
		if dLog.Damage != tt.damage {
			t.Errorf("%d): got %d want %d", i, dLog.Damage, tt.damage)
		}
	}
}

func TestEntireBattle(t *testing.T) {
Test:
	for k, tt := range map[string]struct {
		source         testPokemon
		target         testPokemon
		sourceMoves    []move
		targetMoves    []move
		sourceMoveFunc func() int
		targetMoveFunc func() int
		numTurns       int
		alwaysHit      bool
		exactTargetHP  int
	}{
		"overkill: kills in one turn": {
			source:         testPokemon{15, HAUNTER},
			target:         testPokemon{100, RATICATE},
			sourceMoves:    []move{NIGHTSHADE},
			targetMoves:    []move{BITE},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       1,
		},
		"poison kills in 8 turns (1/8th per turn)": {
			source:         testPokemon{15, BUTTERFREE},
			target:         testPokemon{100, MAGIKARP},
			sourceMoves:    []move{POISONPOWDER},
			targetMoves:    []move{SPLASH},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       8,
			exactTargetHP:  80,
			alwaysHit:      true,
		},
		"bad poison kills in 6 turns (1 + 2 + 3 + 4 + 5 + 6 = 21/16)": {
			source:         testPokemon{15, HAUNTER},
			target:         testPokemon{100, MAGIKARP},
			sourceMoves:    []move{TOXIC},
			targetMoves:    []move{SPLASH},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       6,
			exactTargetHP:  160,
		},
		"dragon rage deals exactly 40 dmg": {
			source:         testPokemon{50, DRAGONITE},
			target:         testPokemon{1, MAGIKARP},
			sourceMoves:    []move{DRAGONRAGE},
			targetMoves:    []move{SPLASH},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       2,
			exactTargetHP:  80,
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		if tt.exactTargetHP != 0 {
			target.Stats.hp = tt.exactTargetHP
			target.currentHP = tt.exactTargetHP
		}
		for i := 0; i < len(tt.sourceMoves); i++ {
			m := &Move{Data: GetMoveDataByID(tt.sourceMoves[i])}
			if tt.alwaysHit {
				m.Data.Accuracy = 100
			}
			source.Moves[i] = m
		}
		for i := 0; i < len(tt.targetMoves); i++ {
			m := &Move{Data: GetMoveDataByID(tt.targetMoves[i])}
			if tt.alwaysHit {
				m.Data.Accuracy = 100
			}
			target.Moves[i] = m
		}
		battle := NewSingleBattle(source, target)
		for !battle.IsOver() {
			sourceCommand := Command{
				SourceIndex: 0,
				TargetIndex: 1,
				MoveIndex:   tt.sourceMoveFunc(),
			}
			targetCommand := Command{
				SourceIndex: 1,
				TargetIndex: 0,
				MoveIndex:   tt.targetMoveFunc(),
			}
			commands := []Command{sourceCommand, targetCommand}
			err := HandleTurn(battle, commands)
			if err != nil {
				t.Errorf("%s: %v", k, err)
				break
			}
			// numTurns+1 since turncounter is upped at end of HandleTurn
			if battle.Log().turn > tt.numTurns+1 {
				testPrintLogs(t, battle.Log().Logs())
				t.Errorf("%s: battle taking too long", k)
				continue Test
			}
		}
		if battle.Log().turn != tt.numTurns+1 {
			testPrintLogs(t, battle.Log().Logs())
			t.Errorf("%s): expected battle to take %d but took %d turns instead", k, tt.numTurns, battle.Log().turn-1)
		}
	}
}

func evaluateNonTextLogs(t *testing.T, gotLogs, wantLogs []battleLog, testpref string) {
	got := filterLogs(gotLogs)
	evaluateLogs(t, got, wantLogs, testpref)
}

func evaluateLogs(t *testing.T, got, want []battleLog, testpref string) {
	if len(got) != len(want) {
		t.Fatalf("%s: got %v want %v", testpref, got, want)
	}
	for i, g := range got {
		wantLog := want[i]
		switch gotLog := g.(type) {
		case DamageLog:
			wantDamageLog, ok := wantLog.(DamageLog)
			if !ok {
				t.Errorf("%s: got %+v want %+v", testpref, gotLog, wantLog)
			}
			if gotLog.Index != wantDamageLog.Index {
				t.Errorf("%s: got %+v want %+v (exact damage doesn't matter)", testpref, gotLog, wantLog)
			}
		default:
			if gotLog != wantLog {
				t.Errorf("%s: got %+v want %+v", testpref, gotLog, wantLog)
			}
		}
	}
}

func filterLogs(logs []battleLog) []battleLog {
	l := []battleLog{}
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

func testPrintLogs(t *testing.T, logs map[int][]battleLog) {
	for i := 1; i <= len(logs); i++ {
		t.Errorf("%#v", logs[i])
	}
}
