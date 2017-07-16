package model

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func MustLoadConfig() {
	initTypeData()
	initMoveData()
	initSpeciesData()
}
