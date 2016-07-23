package main

import (
	"fmt"
	"model"
)

func main() {
	model.MustLoadConfig()
	fmt.Printf("%+v", model.GetPokemon(10, model.BULBASAUR))
}
