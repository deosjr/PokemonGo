package model

import (
    "strings"
    "strconv"
)

var pokemonData []*Species
var nameToSpecies = make(map[string]*Species)

type Species struct {
    Name string
    ID int
    Stats Stats

    Types []*Type
    //Evolutions []Evolution
    //LearnableMoves map[int][]Move

    GrowthRate string
    BaseXP int
    //EP int

    Pokedex string
    Color string
    Height string
    Weight string

    BattlerPlayerY int
    BattlerEnemyY int
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
            values[i], _ = strconv.Atoi(s)
        }
        s.Stats, _ = GetStats(values)
    case "Type1", "Type2":
        t := GetTypeByName(value)
        s.Types = append(s.Types, t)
    case "GrowthRate":
        s.GrowthRate = value
    case "BaseXP":
        s.BaseXP, _ = strconv.Atoi(value)
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

func GetSpecies(name string) *Species {
    key := strings.ToUpper(name)
    return nameToSpecies[key]
}