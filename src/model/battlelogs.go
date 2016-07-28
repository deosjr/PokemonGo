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

func (b *Battle) String() string {
	loglines := make([]string, len(b.Logs))
	for i, log := range b.Logs {
		loglines[i] = log.replay(b)
	}
	return fmt.Sprintf(`
P1: %+v 

P2: %+v

%s`,
		b.pokemon[0], b.pokemon[1], strings.Join(loglines, "\n"))
}
