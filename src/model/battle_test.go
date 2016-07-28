package model

import (
	"testing"
)

func TestHandleMoveSingleBattle(t *testing.T) {
	for _, tt := range []struct {
		p1     *Pokemon
		p2     *Pokemon
		p1move *Move
		p2move *Move
	}{
		{
			p1:     GetPokemon(10, BULBASAUR),
			p2:     GetPokemon(10, CHARMANDER),
			p1move: &Move{Data: GetMoveDataByID(TACKLE)},
			p2move: &Move{Data: GetMoveDataByID(SCRATCH)},
		},
		{
			p1:     GetPokemon(100, PONYTA),
			p2:     GetPokemon(10, CHARMANDER),
			p1move: &Move{Data: GetMoveDataByID(FLAMETHROWER)},
			p2move: &Move{Data: GetMoveDataByID(SCRATCH)},
		},
	} {
		tt.p1.Moves[0] = tt.p1move
		tt.p2.Moves[0] = tt.p2move
		battle := NewSingleBattle(tt.p1, tt.p2)
		commands := []Command{Command{0, 1, 0}, Command{1, 0, 0}}
		if err := battle.HandleTurn(commands); err != nil {
			t.Errorf("%s", err)
		}
		// test with go test -v
		t.Log(battle.String())
	}
}
