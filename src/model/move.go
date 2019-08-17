package model

import (
	"strings"
)

var nameToMove = make(map[string]MoveData)

type MoveData struct {
	Name            string
	Power           int
	Type            pType
	Category        damageCategory
	Accuracy        int
	PP              int
	AddEffectChance int
	Target          target
	Priority        int
	Flags           string
	Description     string

	functionCode   string //TODO: replace with other properties
	applicable     func(source, target *Pokemon) bool
	effect         func(logger *Logger, source, target *Pokemon, sourceIndex, targetIndex, dmgTaken int)
	damageFunction func(source, target *Pokemon) (dmg int, t, crit float64)
	canNotMiss     func(source *Pokemon) bool
}

func initMoveData() {
	for _, m := range moveData {
		nameToMove[strings.ToUpper(m.Name)] = m
	}
}

func GetMoveData(name string) (MoveData, bool) {
	md, ok := nameToMove[strings.ToUpper(name)]
	return md, ok
}

func GetMoveDataByID(m move) MoveData {
	return moveData[m]
}

type Move struct {
	Data      MoveData
	CurrentPP int
	TotalPP   int
}

func GetMoveByName(name string) (*Move, bool) {
	md, ok := GetMoveData(name)
	if !ok {
		return nil, false
	}
	return GetMove(md, md.PP), true
}

func GetMove(move MoveData, pp int) *Move {
	return &Move{
		Data:      move,
		CurrentPP: pp,
		TotalPP:   pp,
	}
}
