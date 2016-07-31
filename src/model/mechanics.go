package model

import (
	"math"
	"math/rand"
)

//TODO
func attackMisses() bool { return false }

func dealDamage(source, target *Pokemon, moveData MoveData) (dmg int, t, crit float64) {
	dmg, t, crit = determineDamage(source, target, moveData)
	return dmg, t, crit
}

func determineDamage(source, target *Pokemon, m MoveData) (dmg int, t, crit float64) {
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
	return int(math.Max(1.0, round(formula*mod))), typeEffectiveness, critical
}

func moveHasEffect(effectChance int) bool {
	if effectChance == 0 {
		return true
	}
	return rand.Intn(100) < effectChance
}

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}
