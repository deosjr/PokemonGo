package model

import (
	"strings"
)

var nameToMove = make(map[string]MoveData)

type MoveData struct {
	Name             string
	functionCode     string //TODO: replace with other properties
	statStageChanges Stats
	Power            int
	Type             pType
	Category         damageCategory
	Accuracy         int
	PP               int
	AddEffectChance  int
	Target           target
	Priority         int
	Flags            string
	Description      string
}

func initMoveData() {
	for _, m := range moveData {
		s := strings.Replace(m.Name, " ", "", -1)
		s = strings.Replace(s, "-", "", -1)
		nameToMove[strings.ToUpper(s)] = m
	}
}

func GetMoveData(name string) MoveData {
	return nameToMove[strings.ToUpper(name)]
}

func GetMoveDataByID(m move) MoveData {
	return moveData[m]
}

type Move struct {
	Data      MoveData
	CurrentPP int
	TotalPP   int
}

func GetMove(move MoveData, pp int) *Move {
	return &Move{
		Data:      move,
		CurrentPP: pp,
		TotalPP:   pp,
	}
}
