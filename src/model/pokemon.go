package model

import ()

type Pokemon struct {
	Level   int
	Species POKEMON
	Name    string
	XP      XP

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

	return &Pokemon{
		Level:      level,
		Species:    pType,
		Name:       species.Name,
		Stats:      calculateStats(ivs, species.Stats, level),
		iv:         ivs,
		statStages: Stats{make([]int, 6)},
		Moves:      make([]*Move, 4),
	}
}

func (p *Pokemon) getSpecies() Species {
	return GetSpeciesByID(p.Species)
}
