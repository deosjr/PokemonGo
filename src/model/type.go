package model

import (
	"strings"
)

type Type struct {
	Name        string
	Weaknesses  []pType
	Resistances []pType
	Immunities  []pType
}

var nameToType = make(map[string]Type)

func initTypeData() {
	for _, t := range typeData {
		nameToType[strings.ToUpper(t.Name)] = t
	}
}

func GetType(name string) Type {
	return nameToType[strings.ToUpper(name)]
}

func GetTypeByID(t pType) Type {
	return typeData[t]
}

func getSTAB(attackType pType, sourceTypes []pType) float64 {
	if typeContains(attackType, sourceTypes) {
		return 1.5
	}
	return 1.0
}

func typeEffectiveness(attackType pType, targetTypes []pType) (t float64) {
	t = 1
	for _, tt := range targetTypes {
		targetType := GetTypeByID(tt)
		if typeContains(attackType, targetType.Immunities) {
			return 0
		}
		if typeContains(attackType, targetType.Weaknesses) {
			t = t * 2
			continue
		}
		if typeContains(attackType, targetType.Resistances) {
			t = t * 0.5
		}
	}
	return t
}

func typeContains(t pType, tt []pType) bool {
	for _, v := range tt {
		if v == t {
			return true
		}
	}
	return false
}
