package singleplayer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/deosjr/PokemonGo/src/model"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// TODO: take input, now fully automatic
func Repl() {
	p1 := getRentalPokemon()
	p2 := getRentalPokemon()
	battle := model.NewSingleBattle(p1, p2)
	fmt.Printf("%s vs %s!\n", p1.Name, p2.Name)

	moveNamesP1 := []string{}
	for _, m := range p1.Moves {
		moveNamesP1 = append(moveNamesP1, m.Data.Name)
	}
	moveNamesP2 := []string{}
	for _, m := range p2.Moves {
		moveNamesP2 = append(moveNamesP2, m.Data.Name)
	}

	fmt.Printf("%s knows %v\n", p1.Name, moveNamesP1)
	fmt.Printf("%s knows %v\n", p2.Name, moveNamesP2)
	totalHP1 := p1.GetTotalHP()
	currentHP1 := totalHP1
	totalHP2 := p2.GetTotalHP()
	currentHP2 := totalHP2

	i := 0
	for !battle.IsOver() {
		i++
		sourceCommand := model.Command{
			SourceIndex: 0,
			TargetIndex: 1,
			MoveIndex:   random.Intn(4),
		}
		targetCommand := model.Command{
			SourceIndex: 1,
			TargetIndex: 0,
			MoveIndex:   random.Intn(4),
		}
		commands := []model.Command{sourceCommand, targetCommand}
		err := model.HandleTurn(battle, commands)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %d/%d | %s: %d/%d\n", p1.Name, currentHP1, totalHP1, p2.Name, currentHP2, totalHP2)
		for _, l := range battle.Log().Logs()[i] {
			switch log := l.(type) {
			case model.DamageLog:
				if log.Index == 0 {
					currentHP1 -= log.Damage
				} else {
					currentHP2 -= log.Damage
				}
			default:
			}
			fmt.Println(l)
		}
	}
}
