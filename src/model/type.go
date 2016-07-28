package model

import (
	"strings"
)

type Type struct {
	Name        string
	Weaknesses  []TYPE
	Resistances []TYPE
	Immunities  []TYPE
}

var nameToType = make(map[string]*Type)

func initTypeData() {
	for _, t := range typeData {
		nameToType[strings.ToUpper(t.Name)] = t
	}
}

func GetType(name string) *Type {
	return nameToType[strings.ToUpper(name)]
}

func GetTypeByID(t TYPE) *Type {
	return typeData[t]
}

func getSTAB(attackType TYPE, sourceTypes []TYPE) float64 {
	if typeContains(attackType, sourceTypes) {
		return 1.5
	}
	return 1.0
}

func typeEffectiveness(attackType TYPE, targetTypes []TYPE) (t float64) {
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

func typeContains(t TYPE, tt []TYPE) bool {
	for _, v := range tt {
		if v == t {
			return true
		}
	}
	return false
}
