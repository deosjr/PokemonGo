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
	targetIndex int
	damage      int
}

func (l DamageLog) replay(b *Battle) string {
	source, _ := pokemonAtIndex(l.targetIndex, b.side1Count, b.maxIndex, b.pokemon)
	source.takeDamage(l.damage)
	return fmt.Sprintf("%s took %d damage!", source.Name, l.damage)
}

func (b *Battle) logf(f string, s ...interface{}) {
	b.Logs = append(b.Logs, TextLog{fmt.Sprintf(f, s...)})
}

func (b *Battle) logDamage(targetIndex, damage int) {
	b.Logs = append(b.Logs, DamageLog{targetIndex, damage})
}

func (b *Battle) initialLog() {
	// TODO: Add Pokemon / Trainer data to log
	// so we can load an entire battle from logdump
	//b.logf("P1: %+v", b.pokemon[0])
	//b.logf("P2: %+v", b.pokemon[1])
}

func (b *Battle) logDamageMessages(name string, dmg int, t, crit float64) {
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

func (b *Battle) String() string {
	loglines := make([]string, len(b.Logs))
	for i, log := range b.Logs {
		loglines[i] = log.replay(b)
	}
	return strings.Join(loglines, "\n")
}

//TODO: func loadBattleFromLog(string) *Battle
