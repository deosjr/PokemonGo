package model

import (
	"strings"
)

var pokemonData []*Species
var nameToSpecies = make(map[string]*Species)

type Species struct {
	Name  string
	ID    int
	Stats Stats

	Types []*Type
	//Evolutions []Evolution
	LearnableMoves map[int][]*MoveData

	GrowthRate string
	BaseXP     int
	//EP int

	Pokedex string
	Color   string
	Height  string
	Weight  string

	BattlerPlayerY  int
	BattlerEnemyY   int
	BattlerAltitude int
}

func (s *Species) update(key, value string) {
	switch key {
	case "Name":
		s.Name = value
	case "InternalName":
		nameToSpecies[value] = s
	case "BaseStats":
		stringValues := strings.Split(value, ",")
		values := make([]int, 6)
		for i, s := range stringValues {
			values[i] = atoiTrusted(s)
		}
		s.Stats, _ = GetStats(values)
	case "Type1", "Type2":
		t := GetType(value)
		s.Types = append(s.Types, t)
	case "Moves":
		s.setMoves(value)
	case "GrowthRate":
		s.GrowthRate = value
	case "BaseXP":
		s.BaseXP = atoiTrusted(value)
	case "Pokedex":
		s.Pokedex = value
	case "Color":
		s.Color = value
	case "Height":
		s.Height = value
	case "Weight":
		s.Weight = value
	}
}

func (s *Species) setMoves(data string) {
	s.LearnableMoves = make(map[int][]*MoveData)
	m := strings.Split(data, ",")
	for i := 0; i < len(m); i += 2 {
		level := atoiTrusted(m[i])
		md := GetMoveData(m[i+1])
		if moves, ok := s.LearnableMoves[level]; ok {
			s.LearnableMoves[level] = append(moves, md)
			continue
		}
		s.LearnableMoves[level] = []*MoveData{md}
	}
}

func GetSpecies(name string) *Species {
	key := strings.ToUpper(name)
	return nameToSpecies[key]
}
