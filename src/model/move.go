package model

import (
	"strings"
)

type MoveData struct {
	Name            string
	functionCode    string
	Power           int
	Type            *Type
	Category        string
	Accuracy        int
	PP              int
	AddEffectChance int
}

var nameToMove = make(map[string]*MoveData)

func LoadMove(features []string) {
	move := &MoveData{
		Name:            features[2],
		functionCode:    features[3],
		Power:           atoiTrusted(features[4]),
		Type:            GetType(features[5]),
		Category:        features[6],
		Accuracy:        atoiTrusted(features[7]),
		PP:              atoiTrusted(features[8]),
		AddEffectChance: atoiTrusted(features[9]),
	}
	nameToMove[features[1]] = move
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