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
