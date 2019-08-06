package model

import (
	"fmt"
)

type Battle interface {
	Log() *Logger
	pokemonAtIndex(int) (*Pokemon, error)
	getTargets(int, int, target) []int
	isOver() bool
}

type singleBattle struct {
	pokemon [2]*Pokemon
	logger  *Logger
}

type Command struct {
	SourceIndex int `json:"source"`
	TargetIndex int `json:"target"`
	MoveIndex   int `json:"move"`
}

type attemptedMove struct {
	Source      *Pokemon
	SourceIndex int
	TargetIndex int
	Move        *Move
}

func HandleTurn(b Battle, commands []Command) error {
	m, err := lookupAttemptedMoves(b, commands)
	if err != nil {
		return err
	}
	err = HandleMoves(b, m)
	b.Log().nextTurn()
	return err
}

func HandleMoves(b Battle, attemptedMoves []attemptedMove) error {
	sortedMoves := sortMoves(attemptedMoves)
	for _, m := range sortedMoves {
		if err := HandleMove(b, m); err != nil {
			return err
		}
	}
	for _, m := range sortedMoves {
		HandlePostMoveEffect(b, m.Source)
	}

	return nil
}

func HandleMove(b Battle, m attemptedMove) error {
	md := m.Move.Data
	b.Log().f("%s used %s!", m.Source.Name, md.Name)
	if md.functionCode != "" {
		b.Log().f("Debug: this move is not (fully) implemented")
	}
	if cantAttack := HandlePreMoveEffect(b, m.Source, m.SourceIndex); cantAttack {
		return nil
	}
	for _, targetIndex := range b.getTargets(m.SourceIndex, m.TargetIndex, md.Target) {
		target, err := b.pokemonAtIndex(targetIndex)
		if err != nil {
			return err
		}
		miss := determineHit(md, m.Source, target)
		if miss {
			b.Log().f("%s's attack missed!", m.Source.Name)
			continue
		}

		var damageTaken int
		if md.Category != statusEffect {
			dmg, t, crit := dealDamage(m.Source, target, md)
			b.Log().damageWithMessages(target.Name, targetIndex, dmg, t, crit)
			damageTaken = target.TakeDamage(dmg)
		}

		if md.effect == nil || !moveHasEffect(md.AddEffectChance) {
			continue
		}
		md.effect(b.Log(), m.Source, target, m.SourceIndex, targetIndex, damageTaken)
	}
	return nil
}

func HandlePreMoveEffect(b Battle, p *Pokemon, index int) (cantAttack bool) {
	if p.NonVolatileCondition == nil {
		return false
	}
	return p.NonVolatileCondition.preMoveEffect(b.Log(), p, index)
}

func HandlePostMoveEffect(b Battle, p *Pokemon) {
	if p.NonVolatileCondition == nil {
		return
	}
	p.NonVolatileCondition.postMoveEffect(b.Log(), p)
}

//Semantics of Battle.total right now (prone to change):
//If len(side1)==side1Count==n and len(side2)==m then total==n+m
//Command.SourceIndex and Command.TargetIndex as follows:
//
//	side1 [0 ... n-1] and side2 [n ... m-1]

func lookupAttemptedMoves(b Battle, commands []Command) ([]attemptedMove, error) {
	attemptedMoves := []attemptedMove{}
	for _, c := range commands {
		m, err := lookupAttemptedMove(b, c)
		if err != nil {
			return nil, err
		}
		attemptedMoves = append(attemptedMoves, m)
	}
	return attemptedMoves, nil
}

func lookupAttemptedMove(b Battle, c Command) (attemptedMove, error) {
	source, err := b.pokemonAtIndex(c.SourceIndex)
	if err != nil {
		return attemptedMove{}, fmt.Errorf("error parsing command: %v", err)
	}

	//TODO: validation according to specific MoveData (target choice!)
	return attemptedMove{
		Source:      source,
		SourceIndex: c.SourceIndex,
		TargetIndex: c.TargetIndex,
		Move:        source.Moves[c.MoveIndex],
	}, nil
}

func NewSingleBattle(p1, p2 *Pokemon) Battle {
	return &singleBattle{
		pokemon: [2]*Pokemon{p1, p2},
		logger:  NewLogger(),
	}
}

func (b *singleBattle) Log() *Logger {
	return b.logger
}

func (b *singleBattle) pokemonAtIndex(i int) (p *Pokemon, err error) {
	if i < 0 || i > 1 {
		return nil, fmt.Errorf("invalid index %d", i)
	}
	return b.pokemon[i], nil
}

func (b *singleBattle) getTargets(sourceIndex, targetIndex int, target target) []int {
	switch target {
	case user, usersSide, singleUsersSide:
		return []int{sourceIndex}
	default:
		return []int{targetIndex}
	}
}

func (b *singleBattle) isOver() bool {
	return b.pokemon[0].currentHP == 0 || b.pokemon[1].currentHP == 0
}
