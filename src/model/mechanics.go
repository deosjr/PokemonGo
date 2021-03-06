package model

import (
	"math"
	"sort"
)

func determineHit(moveData MoveData, source, target *Pokemon) (miss bool) {
	accuracy := source.Accuracy()
	evasion := target.Evasion()
	if moveData.Accuracy == 0 {
		return false
	}
	if moveData.canNotMiss != nil && moveData.canNotMiss(source) {
		return false
	}
	p := float64(moveData.Accuracy) / 100 * (accuracy / evasion)
	return random.Float64() > p
}

func dealDamage(source, target *Pokemon, moveData MoveData) (dmg int, t, crit float64) {
	if moveData.damageFunction != nil {
		return moveData.damageFunction(source, target)
	}
	dmg, t, crit = determineDamage(source, target, moveData)
	return dmg, t, crit
}

func determineDamage(source, target *Pokemon, m MoveData) (dmg int, t, crit float64) {
	attack := attackStat(source, m.Category)
	defense := defenseStat(target, m.Category)
	level := float64(source.Level)
	power := float64(m.Power)
	stab := getSTAB(m.Type, source.getSpecies().Types)
	typeEffectiveness := typeEffectiveness(m.Type, target.getSpecies().Types)
	critical := 1.0 //TODO
	other := 1.0    //TODO
	r := 1.0 - 0.15*random.Float64()
	mod := stab * typeEffectiveness * critical * other * r
	formula := math.Floor(math.Floor(math.Floor(2*level/5+2)*(attack/defense)*power)/50) + 2
	return int(math.Floor(formula * mod)), typeEffectiveness, critical
}

func moveHasEffect(effectChance int) bool {
	if effectChance == 0 {
		return true
	}
	return random.Intn(100) < effectChance
}

func attackStat(p *Pokemon, c damageCategory) float64 {
	var attack int
	var stage int
	if c == physical {
		attack = p.Stats.attack
		stage = p.statStages.attack
	} else if c == special {
		attack = p.Stats.spattack
		stage = p.statStages.spattack
	}
	return modifyStat(attack, stage)
}

func defenseStat(p *Pokemon, c damageCategory) float64 {
	var defense int
	var stage int
	if c == physical {
		defense = p.Stats.defense
		stage = p.statStages.defense
	} else if c == special {
		defense = p.Stats.spdefense
		stage = p.statStages.spdefense
	}
	return modifyStat(defense, stage)
}

// radixsort: first sort on user speed, then on priority
func sortMoves(unsorted []attemptedMove) []attemptedMove {
	moves := []attemptedMove{}
	speed := make(map[int][]int)
	priorities := make(map[int][]int)
	for i, m := range unsorted {
		s := m.Source.Speed()
		if list, ok := speed[s]; ok {
			speed[s] = append(list, i)
			continue
		}
		speed[s] = []int{i}
	}
	speedKeys := []int{}
	for k, _ := range speed {
		speedKeys = append(speedKeys, k)
	}
	sort.Ints(speedKeys)
	for _, k := range speedKeys {
		for _, i := range speed[k] {
			prio := unsorted[i].Move.Data.Priority
			if list, ok := priorities[prio]; ok {
				priorities[prio] = append(list, i)
				continue
			}
			priorities[prio] = []int{i}
		}
	}
	for p := 6; p >= -7; p-- {
		plist := priorities[p]
		// move backwards: speed is sorted low->high
		for n := len(plist) - 1; n >= 0; n-- {
			move := unsorted[plist[n]]
			moves = append(moves, move)
		}
	}
	return moves
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
