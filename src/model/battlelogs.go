package model

import (
	"fmt"
	"strings"
)

type BattleLog interface {
	replay(Battle) string
}

func (b *battle) addToLogs(l BattleLog) {
	if list, ok := b.logs[b.turn]; ok {
		b.logs[b.turn] = append(list, l)
		return
	}
	b.logs[b.turn] = []BattleLog{l}
}

type TextLog struct {
	text string
}

func (l TextLog) replay(b Battle) string {
	return l.text
}

type DamageLog struct {
	index  int
	damage int
}

func (l DamageLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.index)
	target.TakeDamage(l.damage)
	return fmt.Sprintf("DEBUG: %s took %d damage!", target.Name, l.damage)
}

func (b *battle) logf(f string, s ...interface{}) {
	b.addToLogs(TextLog{fmt.Sprintf(f, s...)})
}

func (b *battle) logDamage(index, damage int) {
	b.addToLogs(DamageLog{index, damage})
}

func (b *battle) logDamageWithMessages(name string, index, dmg int, t, crit float64) {
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

func (l StatStageChangeLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.index)
	effectiveChanges, _ := target.ChangeStatStages(l.changes)
	return fmt.Sprintf("DEBUG: %s changed stats: %v!", target.Name, effectiveChanges)
}

func (b *battle) logStatStageChanges(name string, index int, changes Stats, maxed [6]bool) {
	b.addToLogs(StatStageChangeLog{index, changes})
	statNames := []string{"attack", "defense", "special attack", "special defense", "speed"}
	sharply := func(n int) string {
		if n > 1 {
			return " sharply"
		}
		if n < -1 {
			return " harshly"
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

func (b *battle) initialLog() {
	// TODO: Add Pokemon / Trainer data to log
	// so we can load an entire battle from logdump
	//b.logf("P1: %+v", b.pokemon[0])
	//b.logf("P2: %+v", b.pokemon[1])
}

func LogsToString(logs map[int][]BattleLog) string {
	loglines := []string{}
	turn := 1
	for {
		s := TurnToString(logs, turn)
		if s == "" {
			break
		}
		loglines = append(loglines, s)
		turn++
	}
	return strings.Join(loglines, "\n")
}

func TurnToString(logs map[int][]BattleLog, turn int) string {
	loglines := []string{}
	list, ok := logs[turn]
	if !ok {
		return ""
	}
	for _, log := range list {
		switch log.(type) {
		case TextLog:
			loglines = append(loglines, fmt.Sprintf("%d): %s", turn, log.replay(nil)))
		}
	}
	return strings.Join(loglines, "\n")
}

func (b *battle) Logs() map[int][]BattleLog {
	return b.logs
}

//TODO: func loadBattleFromLog(string) *Battle
