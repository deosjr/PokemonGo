package main

import (
	"model"
)

func main() {
	model.MustLoadConfig()
	// p1 := model.GetPokemon(10, model.BULBASAUR)
	// p2 := model.GetPokemon(10, model.CHARMANDER)
	// p1.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.TACKLE)}
	// p2.Moves[3] = &model.Move{Data: model.GetMoveDataByID(model.SCRATCH)}
	// battle := model.NewSingleBattle(p1, p2)
	// commands := []model.Command{model.Command{0, 1, 0}, model.Command{1, 0, 3}}
	// if err := battle.HandleTurn(commands); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(battle.String())
}
