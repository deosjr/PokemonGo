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

type TextLog struct {
	Text string `json:"text"`
}

func (l TextLog) replay(b Battle) string {
	return l.Text
}

type DamageLog struct {
	Index  int `json:"index"`
	Damage int `json:"damage"`
}

func (l DamageLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.Index)
	target.TakeDamage(l.Damage)
	return fmt.Sprintf("Debug: %s took %d damage!", target.Name, l.Damage)
}

func (log *Logger) f(f string, s ...interface{}) {
	log.add(TextLog{fmt.Sprintf(f, s...)})
}

func (log *Logger) damage(index, damage int) {
	log.add(DamageLog{index, damage})
}
func (log *Logger) heal(index, amount int) {
	log.add(DamageLog{index, -amount})
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

type StatStageLog struct {
	Index    int   `json:"index"`
	Changes  Stats `json:"statStageChanges"`
	Accuracy int   `json:"accuracy"`
	Evasion  int   `json:"evasion"`
}

func (l StatStageLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.Index)
	target.ChangeStatStages(l.Changes)
	target.ChangeAccuracy(l.Accuracy)
	target.ChangeEvasion(l.Evasion)
	return fmt.Sprintf("Debug: %s changed stats: %v %d %d!", target.Name, target.Stats, target.accuracyStage, target.evasionStage)
}

func sharply(n int) string {
	if n > 3 {
		n = 3
	}
	if n < -3 {
		n = -3
	}
	switch n {
	case 1:
		return "rose"
	case 2:
		return "sharply rose"
	case 3:
		return "rose drastically"
	case -1:
		return "fell"
	case -2:
		return "sharply fell"
	case -3:
		return "severely fell"
	default:
		return ""
	}
}

var statNames = []string{"attack", "defense", "special attack", "special defense", "speed"}

func (log *Logger) statStages(name string, index int, changes Stats, maxed [6]bool) {
	log.add(StatStageLog{Index: index, Changes: changes})
	for i, v := range changes.stats() {
		if i == 0 {
			// ignore HP
			continue
		}
		log.statStage(name, index, statNames[i-1], v, maxed[i])
	}
}
func (log *Logger) statStage(name string, index int, statName string, change int, maxed bool) {
	if change > 0 {
		if maxed {
			log.f("%s's %s won't go any higher!", name, statName)
			return
		}
		log.f("%s's %s %s!", name, statName, sharply(change))
		return
	}
	if change < 0 {
		if maxed {
			log.f("%s's %s won't go any lower!", name, statName)
			return
		}
		log.f("%s's %s %s!", name, statName, sharply(change))
	}
}

type VolatileConditionLog struct {
	Index     int                    `json:"index"`
	Label     VolatileConditionLabel `json:"label"`
	Condition VolatileCondition      `json:"condition"`
}

func (l VolatileConditionLog) replay(b Battle) string {
	return "TODO"
}

func (log *Logger) volatileCondition(name string, index int, success bool, condition VolatileCondition) {
	if !success {
		log.f(condition.failMessage(), name)
		return
	}
	log.add(VolatileConditionLog{Index: index, Label: condition.getLabel(), Condition: condition})
	log.f(condition.initMessage(), name)
}

type NonVolatileConditionLog struct {
	Index     int                  `json:"index"`
	Condition NonVolatileCondition `json:"condition"`
}

func (l NonVolatileConditionLog) replay(b Battle) string {
	return "TODO"
}

func (log *Logger) nonVolatileCondition(name string, index int, success bool, condition NonVolatileCondition) {
	if !success {
		log.f("%s is already %s", name, condition.name())
		return
	}
	log.add(NonVolatileConditionLog{Index: index, Condition: condition})
	log.f(condition.initMessage(), name)
}

// ugly catch-all log for edge cases
// might change a lot in the future
type GenericUpdateLog struct {
	Index      int   `json:"index"`
	StatStages Stats `json:"statStages"`
}

func (l GenericUpdateLog) replay(b Battle) string {
	target, _ := b.pokemonAtIndex(l.Index)
	target.statStages = l.StatStages
	return fmt.Sprintf("Debug: TODO!")
}

func (log *Logger) Logs() map[int][]battleLog {
	return log.logs
}

//TODO: func loadBattleFromLog(string) *Battle
