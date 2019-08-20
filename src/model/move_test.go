package model

import (
	"testing"
)

func TestSingleMoveSpecifics(t *testing.T) {
	for k, tt := range map[string]struct {
		source       testPokemon
		target       testPokemon
		move         move
		alwaysHit    bool
		alwaysEffect bool
		testTextLogs bool
		logs         []battleLog
	}{
		"Thunder Wave is the only status move that checks type immunity": {
			source: testPokemon{5, PIKACHU},
			target: testPokemon{5, GEODUDE},
			move:   THUNDERWAVE,
			logs:   []battleLog{},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		m := &Move{Data: GetMoveDataByID(tt.move)}
		if tt.alwaysHit {
			m.Data.Accuracy = 100
		}
		if tt.alwaysEffect {
			m.Data.AddEffectChance = 100
		}
		source.Moves[0] = m
		attemptedMove := attemptedMove{
			Source:      source,
			SourceIndex: 0,
			TargetIndex: 1,
			Move:        source.Moves[0],
		}
		battle := NewSingleBattle(source, target)
		HandleMove(battle, attemptedMove)

		if tt.testTextLogs {
			evaluateLogs(t, battle.Log().Logs()[1], tt.logs, k)
		} else {
			evaluateNonTextLogs(t, battle.Log().Logs()[1], tt.logs, k)
		}
	}
}

func TestOneTurnMoveInteractions(t *testing.T) {
	for k, tt := range map[string]struct {
		source       testPokemon
		target       testPokemon
		sourceMove   move
		targetMove   move
		alwaysHit    bool
		alwaysEffect bool
		testTextLogs bool
		logs         []battleLog
	}{
		"Zen Headbutt will flinch the Steelix": {
			source:       testPokemon{59, METAGROSS},
			target:       testPokemon{60, STEELIX},
			sourceMove:   ZENHEADBUTT,
			targetMove:   EARTHQUAKE,
			alwaysHit:    true,
			alwaysEffect: true,
			testTextLogs: true,
			logs: []battleLog{
				TextLog{"Metagross used Zen Headbutt!"},
				DamageLog{Index: 1},
				TextLog{"It's not very effective.."},
				TextLog{"Steelix flinched!"},
			},
		},
		"Psychic type does not affect Dark type, so Umbreon should not flinch": {
			source:       testPokemon{100, METAGROSS},
			target:       testPokemon{50, UMBREON},
			sourceMove:   ZENHEADBUTT,
			targetMove:   BITE,
			alwaysHit:    true,
			alwaysEffect: true,
			testTextLogs: true,
			logs: []battleLog{
				TextLog{"Metagross used Zen Headbutt!"},
				TextLog{"It doesn't affect Umbreon."},
				TextLog{"Umbreon used Bite!"},
				DamageLog{Index: 0},
			},
		},
	} {
		source := GetPokemon(tt.source.level, tt.source.species)
		target := GetPokemon(tt.target.level, tt.target.species)
		sm := &Move{Data: GetMoveDataByID(tt.sourceMove)}
		tm := &Move{Data: GetMoveDataByID(tt.targetMove)}
		if tt.alwaysHit {
			sm.Data.Accuracy = 100
			tm.Data.Accuracy = 100
		}
		if tt.alwaysEffect {
			sm.Data.AddEffectChance = 100
			tm.Data.AddEffectChance = 100
		}
		source.Moves[0] = sm
		target.Moves[0] = tm
		battle := NewSingleBattle(source, target)
		HandleTurn(battle, []Command{{0, 1, 0}, {1, 0, 0}})

		if tt.testTextLogs {
			evaluateLogs(t, battle.Log().Logs()[1], tt.logs, k)
		} else {
			evaluateNonTextLogs(t, battle.Log().Logs()[1], tt.logs, k)
		}

	}
}
