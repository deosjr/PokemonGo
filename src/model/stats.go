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
	return Stats{[]int{rand.Intn(33), rand.Intn(33), rand.Intn(33), rand.Intn(33), rand.Intn(33), rand.Intn(33)}}
}
