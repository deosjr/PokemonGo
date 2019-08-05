package model

import (
	"testing"
)

type testPokemon struct {
	level   int
	species pokemon
}

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
			logs:   []battleLog{damageLog{Index: 1}},
		},
		{
			source: testPokemon{10, SQUIRTLE},
			target: testPokemon{10, CHARMANDER},
			move:   WATERGUN,
			logs:   []battleLog{damageLog{Index: 1}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   GROWL,
			logs:   []battleLog{statStageLog{Index: 1, Changes: Stats{attack: -1}}},
		},
		{
			source: testPokemon{5, CHARMANDER},
			target: testPokemon{5, BULBASAUR},
			move:   FLAMECHARGE,
			logs: []battleLog{
				damageLog{Index: 1},
				statStageLog{Index: 0, Changes: Stats{speed: +1}},
			},
		},
		{
			source: testPokemon{5, PIKACHU},
			target: testPokemon{5, RATTATA},
			move:   THUNDERWAVE,
			logs: []battleLog{
				nonVolatileConditionLog{Index: 1, Condition: Paralysis{}},
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

		evaluateLogs(t, i, battle.Log().Logs()[1], tt.logs)
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
		dLog, ok := log.(damageLog)
		if !ok {
			t.Fatalf("%d): expected a damageLog got %v", i, log)
		}
		if dLog.Damage != tt.damage {
			t.Errorf("%d): got %d want %d", i, dLog.Damage, tt.damage)
		}
	}
}

func TestEntireBattle(t *testing.T) {
	for i, tt := range []struct {
		source         testPokemon
		target         testPokemon
		sourceMoves    []move
		targetMoves    []move
		sourceMoveFunc func() int
		targetMoveFunc func() int
		numTurns       int
		alwaysHit      bool
	}{
		{
			source:         testPokemon{15, HAUNTER},
			target:         testPokemon{100, RATICATE},
			sourceMoves:    []move{NIGHTSHADE},
			targetMoves:    []move{BITE},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       1,
		},
		{
			source:         testPokemon{15, HAUNTER},
			target:         testPokemon{100, MAGIKARP},
			sourceMoves:    []move{TOXIC},
			targetMoves:    []move{SPLASH},
			sourceMoveFunc: func() int { return 0 },
			targetMoveFunc: func() int { return 0 },
			numTurns:       6,
			alwaysHit:      true,
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
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
		for !battle.isOver() {
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
				t.Errorf("%d): %v", i, err)
				break
			}
			// numTurns+1 since turncounter is upped at end of HandleTurn
			if battle.Log().turn > tt.numTurns+1 {
				t.Errorf("%#v", battle.Log().Logs())
				t.Errorf("%d): battle taking too long", i)
				break
			}
		}

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
