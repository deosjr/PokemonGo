package model

import ()

type Pokemon struct {
	Level     int
	Species   POKEMON
	Name      string
	XP        XP
	currentHP int

	Stats      Stats
	iv         Stats
	statStages Stats

	Moves []*Move
	//NonVolatileCondition Condition
	//VolatileConditions []VolatileCondition
}

type XP struct {
	PreviousXPLevelReq int
	CurrentXP          int
	NextXPLevelReq     int
}

func GetPokemon(level int, pType POKEMON) *Pokemon {
	ivs := generateIVs()
	species := GetSpeciesByID(pType)
	stats := calculateStats(ivs, species.Stats, level)

	return &Pokemon{
		Level:      level,
		Species:    pType,
		Name:       species.Name,
		Stats:      stats,
		iv:         ivs,
		currentHP:  stats.HP(),
		statStages: Stats{make([]int, 6)},
		Moves:      make([]*Move, 4),
	}
}

func (p *Pokemon) getSpecies() Species {
	return GetSpeciesByID(p.Species)
}

func (p *Pokemon) takeDamage(damage int) {
	hp := p.currentHP - damage
	if hp < 1 {
		p.currentHP = 1
		return
	}
	p.currentHP = hp
}
