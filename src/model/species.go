package model

import (
	"strings"
)

var nameToSpecies = make(map[string]Species)

type Species struct {
	Name  string
	ID    int
	Stats Stats

	Types          []TYPE
	LearnableMoves map[int][]MOVE

	GrowthRate string
	BaseXP     int

	Pokedex string
	Color   string
	Height  float64
	Weight  float64

	// To be useful
	BattlerPlayerY  int
	BattlerEnemyY   int
	BattlerAltitude int
	GenderRate      string
	StepsToHatch    int
	Habitat         string
	Kind            string
	Abilities       string
	HiddenAbility   string
	Rareness        int
	Happiness       int

	// To be transformed into useful data structures
	EffortPoints string
	// Optional values
	Evolutions       string
	Compatibility    string
	EggMoves         string
	RegionalNumbers  string
	WildItemCommon   string
	WildItemUncommon string
	WildItemRare     string
	FormNames        string
	Incense          string
}

func initSpeciesData() {
	for _, p := range pokemonData {
		s := strings.Replace(p.Name, " ", "", -1)
		s = strings.Replace(s, "-", "", -1)
		s = strings.Replace(s, "'", "", -1)
		s = strings.Replace(s, ".", "", -1)
		nameToSpecies[strings.ToUpper(s)] = p
	}
}

func GetSpecies(name string) Species {
	key := strings.ToUpper(name)
	return nameToSpecies[key]
}

func GetSpeciesByID(pType POKEMON) Species {
	return pokemonData[pType]
}
