package main

import (
	"fmt"

	"model"
)

func main() {
	model.MustLoadConfig()
	species := model.GetSpecies("BULBASAUR")
	fmt.Printf("%+v", model.GetPokemon(10, species))
}
