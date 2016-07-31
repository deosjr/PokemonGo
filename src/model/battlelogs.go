package model

import (
	"fmt"
	"strings"
)

type BattleLog interface {
	replay(*Battle) string
}

type TextLog struct {
	text string
}

func (l TextLog) replay(b *Battle) string {
	return l.text
}

type DamageLog struct {
	index  int
	damage int
}

func (l DamageLog) replay(b *Battle) string {
	target, _ := b.pokemonAtIndex(l.index)
	target.TakeDamage(l.damage)
	return fmt.Sprintf("DEBUG: %s took %d damage!", target.Name, l.damage)
}

func (b *Battle) logf(f string, s ...interface{}) {
	b.Logs = append(b.Logs, TextLog{fmt.Sprintf(f, s...)})
}

func (b *Battle) logDamage(index, damage int) {
	b.Logs = append(b.Logs, DamageLog{index, damage})
}

func (b *Battle) logDamageWithMessages(name string, index, dmg int, t, crit float64) {
	b.logDamage(index, dmg)
	if crit > 1 {
		b.logf("Critical hit!")
	}
	switch {
	case t == 0:
		b.logf("It doesn't affect %s.", name)
	case t > 1:
		b.logf("It's super effective!")
	case t < 1:
		b.logf("It's not very effective..")
	}
}

type StatStageChangeLog struct {
	index   int
	changes Stats
}

func (l StatStageChangeLog) replay(b *Battle) string {
	target, _ := b.pokemonAtIndex(l.index)
	effectiveChanges, _ := target.ChangeStatStages(l.changes)
	return fmt.Sprintf("DEBUG: %s changed stats: %v! New stages: %v", target.Name, effectiveChanges, target.statStages)
}

func (b *Battle) logStatStageChanges(name string, index int, changes Stats, maxed [6]bool) {
	b.Logs = append(b.Logs, StatStageChangeLog{index, changes})
	statNames := []string{"attack", "defense", "special attack", "special defense", "speed"}
	sharply := func(n int) string {
		if n > 1 || n < -1 {
			return " sharply"
		}
		return ""
	}
	for i, v := range changes.stats() {
		if i == 0 {
			continue
		}
		statName := statNames[i-1]
		if v > 0 {
			if maxed[i] {
				b.logf("%s's %s won't go any higher!", name, statName)
				continue
			}
			b.logf("%s's %s%s rose!", name, statName, sharply(v))
			continue
		}
		if v < 0 {
			b.logf("%s's %s%s fell!", name, statName, sharply(v))
		}
	}
}

func (b *Battle) initialLog() {
	// TODO: Add Pokemon / Trainer data to log
	// so we can load an entire battle from logdump
	//b.logf("P1: %+v", b.pokemon[0])
	//b.logf("P2: %+v", b.pokemon[1])
}

func (b *Battle) String() string {
	loglines := make([]string, len(b.Logs))
	for i, log := range b.Logs {
		loglines[i] = log.replay(b)
	}
	return strings.Join(loglines, "\n")
}

//TODO: func loadBattleFromLog(string) *Battle
