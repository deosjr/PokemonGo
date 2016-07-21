package model

import (
	"strings"
)

type MoveData struct {
	Name            string
	functionCode    string
	Power           int
	Type            TYPE
	Category        DMG_CATEGORY
	Accuracy        int
	PP              int
	AddEffectChance int
	Target          TARGET
	Priority        int
	Flags           string
	Description     string
}

var nameToMove = make(map[string]*MoveData)

func initMoveData() {
	for _, m := range moveData {
		s := strings.Replace(m.Name, " ", "", -1)
		s = strings.Replace(s, "-", "", -1)
		nameToMove[strings.ToUpper(s)] = m
	}
}

func GetMoveData(name string) *MoveData {
	return nameToMove[strings.ToUpper(name)]
}

type Move struct {
	MoveData  *MoveData
	CurrentPP int
	TotalPP   int
}

func GetMove(move *MoveData, pp int) *Move {
	return &Move{
		MoveData:  move,
		CurrentPP: pp,
		TotalPP:   pp,
	}
}
