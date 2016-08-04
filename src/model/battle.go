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

func WaitForMoves(battle *Battle, c chan Command, t chan []string) {
	var commands []Command
	for {
		select {
		case command := <-c:
			commands = append(commands, command)
			if len(commands) == 2 {
				battle.HandleTurn(commands)
				commands = []Command{}
				// TODO: Remove this 'fix'!
				t <- []string{battle.String()}
				t <- []string{battle.String()}
			}
		}
	}
}

func NewBattle(p1, p2 []*Pokemon) *Battle {
	b := &Battle{
		pokemon:    append(p1, p2...),
		side1Count: len(p1),
		maxIndex:   len(p1) + len(p2),
	}
	b.initialLog()
	return b
}

func NewSingleBattle(p1, p2 *Pokemon) *Battle {
	b := &Battle{
		pokemon:    []*Pokemon{p1, p2},
		side1Count: 1,
		maxIndex:   2,
	}
	b.initialLog()
	return b
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
	sortedMoves := sortMoves(attemptedMoves)

	for _, m := range sortedMoves {
		if err := b.HandleMove(m); err != nil {
			return err
		}
	}

	return nil
}

func (b *Battle) HandleMove(m attemptedMove) error {
	md := m.Move.Data
	b.logf("%s used %s!", m.Source.Name, md.Name)
	for _, targetIndex := range b.getTargets(m.SourceIndex, m.TargetIndex, md.Target) {
		target, err := b.pokemonAtIndex(targetIndex)
		if err != nil {
			return err
		}
		miss := attackMisses()
		if miss {
			b.logf("%s's attack missed!", m.Source.Name)
			continue
		}

		if md.Category != STATUS {
			dmg, t, crit := dealDamage(m.Source, target, md)
			b.logDamageWithMessages(target.Name, targetIndex, dmg, t, crit)
			target.TakeDamage(dmg)
		}

		if md.statStageChanges != emptyStages {
			if !moveHasEffect(md.AddEffectChance) {
				continue
			}
			effectiveChanges, maxed := target.ChangeStatStages(md.statStageChanges)
			b.logStatStageChanges(target.Name, targetIndex, effectiveChanges, maxed)
		}
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

func (b *Battle) pokemonAtIndex(i int) (p *Pokemon, err error) {
	if i < 0 || i >= b.maxIndex {
		return nil, fmt.Errorf("invalid index %d", i)
	}
	return b.pokemon[i], nil
}

// assumes valid single target
func (b *Battle) getTargets(sourceIndex, targetIndex int, target TARGET) []int {
	//TODO: more than just single battle targets
	switch target {
	case USER, USERS_SIDE, SINGLE_USERS_SIDE:
		return []int{sourceIndex}
	default:
		return []int{targetIndex}
	}
}
