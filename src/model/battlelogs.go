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

func (log *Logger) nextTurn() {
	log.turn++
}

func (log *Logger) Turn() int {
	return log.turn
}

type battleLog interface {
	replay(Battle) string
}

func (log *Logger) add(bl battleLog) {
	if list, ok := log.logs[log.turn]; ok {
		log.logs[log.turn] = append(list, bl)
		return
	}
	log.logs[log.turn] = []battleLog{bl}
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

func (log *Logger) f(f string, s ...interface{}) {
	log.add(textLog{fmt.Sprintf(f, s...)})
}

func (log *Logger) damage(index, damage int) {
	log.add(damageLog{index, damage})
}

func (log *Logger) damageWithMessages(name string, index, dmg int, t, crit float64) {
	log.damage(index, dmg)
	if crit > 1 {
		log.f("Critical hit!")
	}
	switch {
	case t == 0:
		log.f("It doesn't affect %s.", name)
	case t > 1:
		log.f("It's super effective!")
	case t < 1:
		log.f("It's not very effective..")
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

func (log *Logger) statStages(name string, index int, changes Stats, maxed [6]bool) {
	log.add(statStageLog{index, changes})
	for i, v := range changes.stats() {
		if i == 0 {
			continue
		}
		statName := statNames[i-1]
		if v > 0 {
			if maxed[i] {
				log.f("%s's %s won't go any higher!", name, statName)
				continue
			}
			log.f("%s's %s%s rose!", name, statName, sharply(v))
			continue
		}
		if v < 0 {
			if maxed[i] {
				log.f("%s's %s won't go any lower!", name, statName)
				continue
			}
			log.f("%s's %s%s fell!", name, statName, sharply(v))
		}
	}
}

type volatileConditionLog struct {
	Index     int               `json:"index"`
	Condition VolatileCondition `json:"condition"`
}

func (l volatileConditionLog) replay(b Battle) string {
	return "TODO"
}

func (log *Logger) volatileCondition() {

}

type nonVolatileConditionLog struct {
	Index     int                  `json:"index"`
	Condition NonVolatileCondition `json:"condition"`
}

func (l nonVolatileConditionLog) replay(b Battle) string {
	return "TODO"
}

func (log *Logger) nonVolatileCondition(name string, index int, success bool, condition NonVolatileCondition) {
	if !success {
		log.f("%s is already %s", name, condition.name())
		return
	}
	log.add(nonVolatileConditionLog{Index: index, Condition: condition})
	log.f(condition.initMessage(), name)
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

func (log *Logger) Logs() map[int][]battleLog {
	return log.logs
}

//TODO: func loadBattleFromLog(string) *Battle
