package model

type pType byte

const (
	NORMAL pType = iota
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

var typeData = []Type{
	{
		Name:       "Normal",
		Weaknesses: []pType{FIGHTING},
		Immunities: []pType{GHOST},
	},
	{
		Name:        "Fighting",
		Weaknesses:  []pType{FLYING, PSYCHIC},
		Resistances: []pType{ROCK, BUG, DARK},
	},
	{
		Name:        "Flying",
		Weaknesses:  []pType{ROCK, ELECTRIC, ICE},
		Resistances: []pType{FIGHTING, BUG, GRASS},
		Immunities:  []pType{GROUND},
	},
	{
		Name:        "Poison",
		Weaknesses:  []pType{GROUND, PSYCHIC},
		Resistances: []pType{FIGHTING, POISON, BUG, GRASS},
	},
	{
		Name:        "Ground",
		Weaknesses:  []pType{WATER, GRASS, ICE},
		Resistances: []pType{POISON, ROCK},
		Immunities:  []pType{ELECTRIC},
	},
	{
		Name:        "Rock",
		Weaknesses:  []pType{FIGHTING, GROUND, STEEL, WATER, GRASS},
		Resistances: []pType{NORMAL, FLYING, POISON, FIRE},
	},
	{
		Name:        "Bug",
		Weaknesses:  []pType{FLYING, ROCK, FIRE},
		Resistances: []pType{FIGHTING, GROUND, GRASS},
	},
	{
		Name:        "Ghost",
		Weaknesses:  []pType{GHOST, DARK},
		Resistances: []pType{POISON, BUG},
		Immunities:  []pType{NORMAL, FIGHTING},
	},
	{
		Name:        "Steel",
		Weaknesses:  []pType{FIGHTING, GROUND, FIRE},
		Resistances: []pType{NORMAL, FLYING, ROCK, BUG, GHOST, STEEL, GRASS, PSYCHIC, ICE, DRAGON, DARK},
		Immunities:  []pType{POISON},
	},
	{
		Name:        "Fire",
		Weaknesses:  []pType{GROUND, ROCK, WATER},
		Resistances: []pType{BUG, STEEL, FIRE, GRASS, ICE},
	},
	{
		Name:        "Water",
		Weaknesses:  []pType{GRASS, ELECTRIC},
		Resistances: []pType{STEEL, FIRE, WATER, ICE},
	},
	{
		Name:        "Grass",
		Weaknesses:  []pType{FLYING, POISON, BUG, FIRE, ICE},
		Resistances: []pType{GROUND, WATER, GRASS, ELECTRIC},
	},
	{
		Name:        "Electric",
		Weaknesses:  []pType{GROUND},
		Resistances: []pType{FLYING, STEEL, ELECTRIC},
	},
	{
		Name:        "Psychic",
		Weaknesses:  []pType{BUG, GHOST, DARK},
		Resistances: []pType{FIGHTING, PSYCHIC},
	},
	{
		Name:        "Ice",
		Weaknesses:  []pType{FIGHTING, ROCK, STEEL, FIRE},
		Resistances: []pType{ICE},
	},
	{
		Name:        "Dragon",
		Weaknesses:  []pType{ICE, DRAGON},
		Resistances: []pType{FIRE, WATER, GRASS, ELECTRIC},
	},
	{
		Name:        "Dark",
		Weaknesses:  []pType{FIGHTING, BUG},
		Resistances: []pType{GHOST, DARK},
		Immunities:  []pType{PSYCHIC},
	},
}
