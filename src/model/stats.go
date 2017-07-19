package model

type Stats struct {
	hp        int
	attack    int
	defense   int
	spattack  int
	spdefense int
	speed     int
}

var emptyStages Stats

func GetStats(values [6]int) Stats {
	hp, attack, defense, spattack, spdefense, speed := values[0], values[1], values[2], values[3], values[4], values[5]
	return Stats{hp, attack, defense, spattack, spdefense, speed}
}

func (s Stats) stats() [6]int {
	return [6]int{s.hp, s.attack, s.defense, s.spattack, s.spdefense, s.speed}
}

func (s Stats) updateStages(changes Stats) (Stats, Stats, [6]bool) {
	newStages := [6]int{}
	effective := [6]int{}
	maxed := [6]bool{}
	stats, changeStats := s.stats(), changes.stats()
	for i := 0; i < 6; i++ {
		newStages[i], effective[i], maxed[i] = validStatStage(stats[i], changeStats[i])
	}
	return GetStats(newStages), GetStats(effective), maxed
}

func validStatStage(stage, change int) (changed, effective int, maxed bool) {
	switch {
	case stage == 6 && change > 0:
		return 6, 0, true
	case stage == -6 && change < 0:
		return -6, 0, true
	default:
		totalChanged := stage + change
		effective = totalChanged - stage
		switch {
		case totalChanged > 6:
			return 6, effective, false
		case totalChanged < -6:
			return -6, effective, false
		default:
			return totalChanged, effective, false
		}
	}
}

func calculateStats(ivs, base Stats, level int) Stats {
	stats := [6]int{}
	ivStats, baseStats := ivs.stats(), base.stats()
	stats[0] = calculateHP(ivs.hp, base.hp, level)
	for i := 1; i < 6; i++ {
		stats[i] = calculateStat(ivStats[i], baseStats[i], level)
	}
	return GetStats(stats)
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
	return GetStats([6]int{random.Intn(32), random.Intn(32), random.Intn(32), random.Intn(32), random.Intn(32), random.Intn(32)})
}

func modifyStat(stat, stage int) float64 {
	if stage >= 0 {
		return float64(stat) * (float64(stage)*0.5 + 1)
	}
	return float64(stat) * (1 / ((-float64(stage) * 0.5) + 1))
}

func accuracyOrEvasionToMod(stage int) float64 {
	if stage >= 0 {
		return float64(stage+3) / 3.0
	}
	return 3 / float64(3-stage)
}
