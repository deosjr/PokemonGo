package model

import "fmt"

type Logger struct {
	turn int
	logs map[int][]battleLog
}

func NewLogger() *Logger {
	return &Logger{
		turn: 1,
		logs: make(map[int][]battleLog),
	}
	// TODO: Add Pokemon / Trainer data to log
	// so we can load an entire battle from logdump
	//b.logf("P1: %+v", b.pokemon[0])
	//b.logf("P2: %+v", b.pokemon[1])
}

func (l *Logger) nextTurn() {
	l.turn++
}

func (l *Logger) Turn() int {
	return l.turn
}

type battleLog interface {
	replay(Battle) string
}

func (l *Logger) addToLogs(bl battleLog) {
	if list, ok := l.logs[l.turn]; ok {
		l.logs[l.turn] = append(list, bl)
		return
	}
	l.logs[l.turn] = []battleLog{bl}
}

type textLog struct {
	Text string `json:"text"`
}

func (l textLog) replay(b Battle) string {
	return l.Text
}

type damageLog struct {
	Index  int `json:"index"`
	Damage int `json:"damage"`
}

func (l damageLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.Index)
	target.TakeDamage(l.Damage)
	return fmt.Sprintf("Debug: %s took %d damage!", target.Name, l.Damage)
}

func (l *Logger) logf(f string, s ...interface{}) {
	l.addToLogs(textLog{fmt.Sprintf(f, s...)})
}

func (l *Logger) logDamage(index, damage int) {
	l.addToLogs(damageLog{index, damage})
}

func (l *Logger) logDamageWithMessages(name string, index, dmg int, t, crit float64) {
	l.logDamage(index, dmg)
	if crit > 1 {
		l.logf("Critical hit!")
	}
	switch {
	case t == 0:
		l.logf("It doesn't affect %s.", name)
	case t > 1:
		l.logf("It's super effective!")
	case t < 1:
		l.logf("It's not very effective..")
	}
}

type statStageLog struct {
	Index   int   `json:"index"`
	Changes Stats `json:"statStageChanges"`
}

func (l statStageLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.Index)
	effectiveChanges, _ := target.ChangeStatStages(l.Changes)
	return fmt.Sprintf("Debug: %s changed stats: %v!", target.Name, effectiveChanges)
}

func sharply(n int) string {
	if n > 1 {
		return " sharply"
	}
	if n < -1 {
		return " harshly"
	}
	return ""
}

var statNames = []string{"attack", "defense", "special attack", "special defense", "speed"}

func (l *Logger) logStatStages(name string, index int, changes Stats, maxed [6]bool) {
	l.addToLogs(statStageLog{index, changes})
	for i, v := range changes.stats() {
		if i == 0 {
			continue
		}
		statName := statNames[i-1]
		if v > 0 {
			if maxed[i] {
				l.logf("%s's %s won't go any higher!", name, statName)
				continue
			}
			l.logf("%s's %s%s rose!", name, statName, sharply(v))
			continue
		}
		if v < 0 {
			if maxed[i] {
				l.logf("%s's %s won't go any lower!", name, statName)
				continue
			}
			l.logf("%s's %s%s fell!", name, statName, sharply(v))
		}
	}
}

// ugly catch-all log for edge cases
// might change a lot in the future
type genericUpdateLog struct {
	Index      int   `json:"index"`
	StatStages Stats `json:"statStages"`
}

func (l genericUpdateLog) replay(b Battle) string {
	return fmt.Sprintf("Debug: TODO!")
}

func (l *Logger) Logs() map[int][]battleLog {
	return l.logs
}

//TODO: func loadBattleFromLog(string) *Battle
