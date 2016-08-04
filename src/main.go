package main

import (
	"fmt"
	"net/http"

	"handler"
	"model"
)

func main() {
	model.MustLoadConfig()
	fmt.Println("Serving..")

	p1 := model.GetPokemon(10, model.BULBASAUR)
	p2 := model.GetPokemon(10, model.CHARMANDER)
	p1.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.TACKLE)}
	p2.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.EMBER)}
	battle := model.NewSingleBattle(p1, p2)

	c := make(chan model.Command)
	t := make(chan []string)
	handler.Battle = battle
	handler.CommandChannel = c
	handler.TurnChannel = t
	go model.WaitForMoves(battle, c, t)

	http.HandleFunc("/move", handler.HandleMove)
	http.ListenAndServe(":8080", nil)

}
