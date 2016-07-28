package model

import (
	"fmt"
	"math"
	"math/rand"
)

func HandleMove(m MoveData, source, target *Pokemon) {
	miss := attackMisses()

	if !(miss || m.Category == STATUS) {
		dealDamage(source, target, m)
	}
	// TODO: move functions
}

//TODO
func attackMisses() bool { return false }

func dealDamage(source, target *Pokemon, m MoveData) {
	dmg := determineDamage(source, target, m)
	fmt.Printf("%s took %d damage!\n", target.Name, dmg)
}

func determineDamage(source, target *Pokemon, m MoveData) int {
	attack := attackStat(source.Stats, source.statStages, m.Category)
	defense := defenseStat(target.Stats, target.statStages, m.Category)
	level := float64(source.Level)
	power := float64(m.Power)
	stab := getSTAB(m.Type, source.getSpecies().Types)
	typeEffectiveness := typeEffectiveness(m.Type, target.getSpecies().Types)
	critical := 1.0 //TODO
	other := 1.0    //TODO
	random := 1.0 - 0.15*rand.Float64()
	mod := stab * typeEffectiveness * critical * other * random
	formula := ((2*level+10)/250)*(attack/defense)*power + 2
	return int(math.Max(1.0, round(formula*mod)))
}

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}
