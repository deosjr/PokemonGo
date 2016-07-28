package model

import (
	"fmt"
	"strings"
)

/*
Semantics of Battle.total right now (prone to change):
If len(side1)==n and len(side2)==m then total==n+m
Command.SourceIndex and Command.TargetIndex as follows:

	side1 [0 ... n-1] and side2 [n ... m-1]
*/

type Battle struct {
	side1 []*Pokemon
	side2 []*Pokemon
	total int

	Logs []string
}

type Command struct {
	SourceIndex int
	TargetIndex int
	MoveIndex   int
}

type attemptedMove struct {
	Source *Pokemon
	Target *Pokemon
	Move   *Move
}

func NewBattle(p1, p2 []*Pokemon) *Battle {
	return &Battle{
		side1: p1,
		side2: p2,
		total: len(p1) + len(p2),
	}
}

func NewSingleBattle(p1, p2 *Pokemon) *Battle {
	return &Battle{
		side1: []*Pokemon{p1},
		side2: []*Pokemon{p2},
		total: 2,
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
		b.logf("%s used %s!", m.Source.Name, m.Move.Data.Name)
		HandleMove(m.Move.Data, m.Source, m.Target)
	}

	return nil
}

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
	var source *Pokemon
	var target *Pokemon
	if c.SourceIndex >= b.total || c.TargetIndex >= b.total || c.MoveIndex > 3 {
		return attemptedMove{}, fmt.Errorf("Invalid command %v", c)
	}

	if c.SourceIndex < len(b.side1) {
		source = b.side1[c.SourceIndex]
	} else {
		source = b.side2[c.SourceIndex-len(b.side1)]
	}
	if c.TargetIndex < len(b.side1) {
		target = b.side1[c.TargetIndex]
	} else {
		target = b.side2[c.TargetIndex-len(b.side1)]
	}

	return attemptedMove{source, target, source.Moves[c.MoveIndex]}, nil
}

func (b *Battle) logf(f string, s ...interface{}) {
	b.Logs = append(b.Logs, fmt.Sprintf(f, s...))
}

func (b *Battle) String() string {
	return fmt.Sprintf(`
P1: %+v 

P2: %+v

%s`,
		b.side1[0], b.side2[0], strings.Join(b.Logs, "\n"))
}
