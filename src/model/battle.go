package model

import (
	"fmt"
)

type Battle struct {
	pokemon    []*Pokemon
	side1Count int
	maxIndex   int

	Logs []BattleLog
}

type Command struct {
	SourceIndex int
	TargetIndex int
	MoveIndex   int
}

type attemptedMove struct {
	Source      *Pokemon
	SourceIndex int
	Target      *Pokemon
	TargetIndex int
	Move        *Move
}

func NewBattle(p1, p2 []*Pokemon) *Battle {
	return &Battle{
		pokemon:    append(p1, p2...),
		side1Count: len(p1),
		maxIndex:   len(p1) + len(p2),
	}
}

func NewSingleBattle(p1, p2 *Pokemon) *Battle {
	return &Battle{
		pokemon:    []*Pokemon{p1, p2},
		side1Count: 1,
		maxIndex:   2,
	}
}

func (b *Battle) HandleTurn(commands []Command) error {
	m, err := b.lookupAttemptedMoves(commands)
	if err != nil {
		return err
	}
	return b.HandleMoves(m)
}

func (b *Battle) HandleTurnSingleBattle(c1, c2 Command) error {
	return b.HandleTurn([]Command{c1, c2})
}

func (b *Battle) HandleMoves(attemptedMoves []attemptedMove) error {
	// TODO: sort by priority

	for _, m := range attemptedMoves {
		b.HandleMove(m)
	}

	return nil
}

//Semantics of Battle.total right now (prone to change):
//If len(side1)==side1Count==n and len(side2)==m then total==n+m
//Command.SourceIndex and Command.TargetIndex as follows:
//
//	side1 [0 ... n-1] and side2 [n ... m-1]

func (b *Battle) lookupAttemptedMoves(commands []Command) ([]attemptedMove, error) {
	attemptedMoves := []attemptedMove{}
	for _, c := range commands {
		m, err := b.lookupAttemptedMove(c)
		if err != nil {
			return nil, err
		}
		attemptedMoves = append(attemptedMoves, m)
	}
	return attemptedMoves, nil
}

func (b *Battle) lookupAttemptedMove(c Command) (attemptedMove, error) {
	source, err := pokemonAtIndex(c.SourceIndex, b.side1Count, b.maxIndex, b.pokemon)
	if err != nil {
		return attemptedMove{}, fmt.Errorf("error parsing command: %v", err)
	}
	target, err := pokemonAtIndex(c.TargetIndex, b.side1Count, b.maxIndex, b.pokemon)
	if err != nil {
		return attemptedMove{}, fmt.Errorf("error parsing command: %v", err)
	}

	//TODO: validation according to specific MoveData (target choice!)
	return attemptedMove{
		Source:      source,
		SourceIndex: c.SourceIndex,
		Target:      target,
		TargetIndex: c.TargetIndex,
		Move:        source.Moves[c.MoveIndex],
	}, nil
}

func pokemonAtIndex(i, side1Count, total int, pokemon []*Pokemon) (p *Pokemon, err error) {
	if i < 0 || i >= total {
		return nil, fmt.Errorf("invalid index %d", i)
	}
	return pokemon[i], nil
}
