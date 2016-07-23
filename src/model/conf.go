package model

import (
	"math/rand"
	"time"
)

func MustLoadConfig() {
	initTypeData()
	initMoveData()
	initSpeciesData()

	rand.Seed(time.Now().UnixNano())
}
