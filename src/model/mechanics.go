package model

import (
	"math"
	"math/rand"
	"sort"
)

//TODO
func attackMisses() bool { return false }

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

func changeStatStages(log *Logger, p *Pokemon, index int, changes Stats) {
	stats, effectiveChanges, maxed := p.ChangeStatStages(changes)
	log.logStatStages(p.Name, index, stats, effectiveChanges, maxed)
}

func fixedDamage(damage int, typ pType, target *Pokemon) (int, float64, float64) {
	t := typeEffectiveness(typ, target.getSpecies().Types)
	if t == 0 {
		return 0, 0, 1
	}
	return damage, 1, 1
}
