package model

import (
	"fmt"
	"math/rand"
)

type Stats struct {
	stats []int
}

func (s Stats) HP() int {
	return s.stats[0]
}
func (s Stats) Attack() int {
	return s.stats[1]
}
func (s Stats) Defense() int {
	return s.stats[2]
}
func (s Stats) SPAttack() int {
	return s.stats[3]
}
func (s Stats) SPDefense() int {
	return s.stats[4]
}
func (s Stats) Speed() int {
	return s.stats[5]
}

func GetStats(values []int) (Stats, error) {
	if len(values) != 6 {
		return Stats{}, fmt.Errorf("Length stats is %d instead of 6", len(values))
	}
	return Stats{values}, nil
}

func calculateStats(ivs, base Stats, level int) Stats {
	stats := make([]int, 6)
	stats[0] = calculateHP(ivs.stats[0], base.stats[0], level)
	for i := 1; i < 6; i++ {
		stats[i] = calculateStat(ivs.stats[i], base.stats[i], level)
	}
	return Stats{stats}
}

func calculateHP(iv, base, level int) int {
	temp := iv + (2 * base) // + EV/4
	temp = ((temp * level) / 100) + 5
	return temp // x Nature ?
}

func calculateStat(iv, base, level int) int {
	temp := iv + (2 * base) + 100 // + EV/4
	temp = ((temp * level) / 100) + 10
	return temp // x Nature ?
}

func generateIVs() Stats {
	return Stats{[]int{rand.Intn(32), rand.Intn(32), rand.Intn(32), rand.Intn(32), rand.Intn(32), rand.Intn(32)}}
}

func attackStat(stats, stages Stats, c DMG_CATEGORY) float64 {
	var attack int
	var stage int
	if c == PHYSICAL {
		attack = stats.Attack()
		stage = stages.Attack()
	} else if c == SPECIAL {
		attack = stats.SPAttack()
		stage = stats.SPAttack()
	}
	return modifyStat(attack, stage)
}

func defenseStat(stats, stages Stats, c DMG_CATEGORY) float64 {
	var defense int
	var stage int
	if c == PHYSICAL {
		defense = stats.Defense()
		stage = stages.Defense()
	} else if c == SPECIAL {
		defense = stats.SPDefense()
		stage = stats.SPDefense()
	}
	return modifyStat(defense, stage)
}

func modifyStat(stat, stage int) float64 {
	if stage >= 0 {
		return float64(stat) * (float64(stage)*0.5 + 1)
	}
	return float64(stat) * (1 / ((-float64(stage) * 0.5) + 1))
}
