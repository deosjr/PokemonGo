package main

import (
	"fmt"
	"net/http"

	"github.com/deosjr/PokemonGo/src/handler"
	"github.com/deosjr/PokemonGo/src/model"
)

func main() {
	model.MustLoadConfig()
	fmt.Println("Serving..")

	p1 := model.GetPokemon(10, model.BULBASAUR)
	p2 := model.GetPokemon(10, model.CHARMANDER)
	p1.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.TACKLE)}
	p2.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.EMBER)}
	battle := model.NewSingleBattle(p1, p2)

	h := handler.NewHandler()
	h.AddBattle("test", battle)

	http.HandleFunc("/move", h.HandleMove)
	http.ListenAndServe(":8080", nil)

}
