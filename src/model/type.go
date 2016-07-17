package model

import (
	"strings"
)

type TYPE int

const (
	NORMAL TYPE = iota
	FIGHTING
	FLYING
	POISON
	GROUND
	ROCK
	BUG
	GHOST
	STEEL
	FIRE
	WATER
	GRASS
	ELECTRIC
	PSYCHIC
	ICE
	DRAGON
	DARK
)

type Type struct {
	Name        string
	Weaknesses  []TYPE
	Resistances []TYPE
	Immunities  []TYPE
}

var typeData = []*Type{
	{
		Name:       "Normal",
		Weaknesses: []TYPE{FIGHTING},
		Immunities: []TYPE{GHOST},
	},
	{
		Name:        "Fighting",
		Weaknesses:  []TYPE{FLYING, PSYCHIC},
		Resistances: []TYPE{ROCK, BUG, DARK},
	},
	{
		Name:        "Flying",
		Weaknesses:  []TYPE{ROCK, ELECTRIC, ICE},
		Resistances: []TYPE{FIGHTING, BUG, GRASS},
		Immunities:  []TYPE{GROUND},
	},
	{
		Name:        "Poison",
		Weaknesses:  []TYPE{GROUND, PSYCHIC},
		Resistances: []TYPE{FIGHTING, POISON, BUG, GRASS},
	},
	{
		Name:        "Ground",
		Weaknesses:  []TYPE{WATER, GRASS, ICE},
		Resistances: []TYPE{POISON, ROCK},
		Immunities:  []TYPE{ELECTRIC},
	},
	{
		Name:        "Rock",
		Weaknesses:  []TYPE{FIGHTING, GROUND, STEEL, WATER, GRASS},
		Resistances: []TYPE{NORMAL, FLYING, POISON, FIRE},
	},
	{
		Name:        "Bug",
		Weaknesses:  []TYPE{FLYING, ROCK, FIRE},
		Resistances: []TYPE{FIGHTING, GROUND, GRASS},
	},
	{
		Name:        "Ghost",
		Weaknesses:  []TYPE{GHOST, DARK},
		Resistances: []TYPE{POISON, BUG},
		Immunities:  []TYPE{NORMAL, FIGHTING},
	},
	{
		Name:        "Steel",
		Weaknesses:  []TYPE{FIGHTING, GROUND, FIRE},
		Resistances: []TYPE{NORMAL, FLYING, ROCK, BUG, GHOST, STEEL, GRASS, PSYCHIC, ICE, DRAGON, DARK},
		Immunities:  []TYPE{POISON},
	},
	{
		Name:        "Fire",
		Weaknesses:  []TYPE{GROUND, ROCK, WATER},
		Resistances: []TYPE{BUG, STEEL, FIRE, GRASS, ICE},
	},
	{
		Name:        "Water",
		Weaknesses:  []TYPE{GRASS, ELECTRIC},
		Resistances: []TYPE{STEEL, FIRE, WATER, ICE},
	},
	{
		Name:        "Grass",
		Weaknesses:  []TYPE{FLYING, POISON, BUG, FIRE, ICE},
		Resistances: []TYPE{GROUND, WATER, GRASS, ELECTRIC},
	},
	{
		Name:        "Electric",
		Weaknesses:  []TYPE{GROUND},
		Resistances: []TYPE{FLYING, STEEL, ELECTRIC},
	},
	{
		Name:        "Psychic",
		Weaknesses:  []TYPE{BUG, GHOST, DARK},
		Resistances: []TYPE{FIGHTING, PSYCHIC},
	},
	{
		Name:        "Ice",
		Weaknesses:  []TYPE{FIGHTING, ROCK, STEEL, FIRE},
		Resistances: []TYPE{ICE},
	},
	{
		Name:        "Dragon",
		Weaknesses:  []TYPE{ICE, DRAGON},
		Resistances: []TYPE{FIRE, WATER, GRASS, ELECTRIC},
	},
	{
		Name:        "Dark",
		Weaknesses:  []TYPE{FIGHTING, BUG},
		Resistances: []TYPE{GHOST, DARK},
		Immunities:  []TYPE{PSYCHIC},
	},
}

var nameToType = make(map[string]*Type)

func initTypeMap() {
	for _, t := range typeData {
		nameToType[strings.ToUpper(t.Name)] = t
	}
}

func GetType(name string) *Type {
	return nameToType[strings.ToUpper(name)]
}
