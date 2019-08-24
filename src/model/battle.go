package model

import (
	"fmt"
)

type Battle interface {
	Log() *Logger
	IsOver() bool
	Graphics(player int) graphicInfo

	pokemonAtIndex(int) (*Pokemon, error)
	getTargets(int, int, target) []int
}

type singleBattle struct {
	pokemon [2]*Pokemon
	logger  *Logger
}

type MoveCommand struct {
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

func HandleTurn(b Battle, commands []MoveCommand) error {
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
		HandlePostMoveEffect(b, m.SourceIndex, m.Source)
	}
	if b.IsOver() {
		return nil
	}

	return nil
}

func HandleMove(b Battle, m attemptedMove) error {
	md := m.Move.Data
	if cantAttack := HandlePreMoveEffect(b, m.Source, m.SourceIndex); cantAttack {
		return nil
	}
	if b.IsOver() {
		return nil
	}
	b.Log().f("%s used %s!", m.Source.Name, md.Name)
	if md.functionCode != "" {
		b.Log().f("Debug: this move is not (fully) implemented")
	}
	for _, targetIndex := range b.getTargets(m.SourceIndex, m.TargetIndex, md.Target) {
		target, err := b.pokemonAtIndex(targetIndex)
		if err != nil {
			return err
		}
		if md.applicable != nil && !md.applicable(m.Source, target) {
			b.Log().f("But it failed!")
			continue
		}
		miss := determineHit(md, m.Source, target)
		if miss {
			b.Log().f("%s's attack missed!", m.Source.Name)
			continue
		}

		var damageTaken int
		typeEffectiveness := 1.0
		if md.Category != statusEffect {
			dmg, t, crit := dealDamage(m.Source, target, md)
			b.Log().damageWithMessages(target.Name, targetIndex, dmg, t, crit)
			damageTaken = target.TakeDamage(dmg)
			typeEffectiveness = t
		}
		// TODO: this ends the battle too early if a recoil move was used
		if b.IsOver() {
			return nil
		}

		if md.effect == nil || typeEffectiveness == 0 || !moveHasEffect(md.AddEffectChance) {
			continue
		}
		md.effect(b.Log(), m.Source, target, m.SourceIndex, targetIndex, damageTaken)
		if b.IsOver() {
			return nil
		}
	}
	return nil
}

func HandlePreMoveEffect(b Battle, p *Pokemon, index int) (cantAttack bool) {
	if p.NonVolatileCondition != nil && p.NonVolatileCondition.preMoveEffect(b.Log(), p, index) {
		return true
	}
	// TODO: sort so that flinch comes before/after confusion?
	for _, c := range p.VolatileConditions {
		if c.preMoveEffect(b.Log(), p, index) {
			return true
		}
	}
	return false
}

func HandlePostMoveEffect(b Battle, i int, p *Pokemon) {
	if p.NonVolatileCondition != nil {
		p.NonVolatileCondition.postMoveEffect(b.Log(), i, p)
	}
	for _, c := range p.VolatileConditions {
		c.postMoveEffect(p)
	}
}

//Semantics of Battle.total right now (prone to change):
//If len(side1)==side1Count==n and len(side2)==m then total==n+m
//Command.SourceIndex and Command.TargetIndex as follows:
//
//	side1 [0 ... n-1] and side2 [n ... m-1]

func lookupAttemptedMoves(b Battle, commands []MoveCommand) ([]attemptedMove, error) {
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

func lookupAttemptedMove(b Battle, c MoveCommand) (attemptedMove, error) {
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

func (b *singleBattle) IsOver() bool {
	return b.pokemon[0].currentHP == 0 || b.pokemon[1].currentHP == 0
}

type graphicInfo struct {
	// list of species per side
	Sides [][]int
	Moves [4]string
}

func (b *singleBattle) Graphics(player int) graphicInfo {
	if player != 0 && player != 1 {
		panic("unexpected player")
	}
	var self, opp *Pokemon
	switch player {
	case 0:
		self = b.pokemon[0]
		opp = b.pokemon[1]
	case 1:
		self = b.pokemon[1]
		opp = b.pokemon[0]
	}

	moves := [4]string{}
	for i, m := range self.Moves {
		if m == nil {
			moves[i] = "-"
			continue
		}
		moves[i] = m.Data.Name
	}
	return graphicInfo{
		Sides: [][]int{{self.ID()}, {opp.ID()}},
		Moves: moves,
	}
}
