package main 

import (
    "fmt"

    "model"
)

func main() {
    model.MustLoadConfig()
    fmt.Printf("%+v", model.TestRemoveGetPokemonByID(1))
}