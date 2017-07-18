package model

import (
	"math"
)

type NonVolatileCondition interface {
	name() string
	initMessage() string
	applicable(*Pokemon) bool
	preMoveEffect(*Logger, *Pokemon, int) (cantAttack bool)
	postMoveEffect(*Logger, *Pokemon)
}

type nonVolatileCondition struct{}

func (c nonVolatileCondition) name() string                              { return "" }
func (c nonVolatileCondition) initMessage() string                       { return "" }
func (c nonVolatileCondition) applicable(*Pokemon) bool                  { return true }
func (c nonVolatileCondition) preMoveEffect(*Logger, *Pokemon, int) bool { return false }
func (c nonVolatileCondition) postMoveEffect(*Logger, *Pokemon)          {}

type Burn struct {
	nonVolatileCondition
}

func (c Burn) name() string {
	return "burned"
}
func (c Burn) initMessage() string {
	return "%s was burned!"
}

func (c Burn) applicable(p *Pokemon) bool {
	if typeContains(FIRE, p.getSpecies().Types) {
		return false
	}
	return true
}

func (c Burn) postMoveEffect(log *Logger, p *Pokemon) {
	// TODO: hp / 16 since Gen VII
	damage := int(math.Ceil(float64(p.Stats.hp) / 8))
	p.TakeDamage(damage)
	log.f("%s is hurt by its burn!", p.Name)
}

type Freeze struct {
	nonVolatileCondition
}

func (c Freeze) name() string {
	return "frozen"
}
func (c Freeze) initMessage() string {
	return "%s is frozen solid!"
}

func (c Freeze) applicable(p *Pokemon) bool {
	if typeContains(ICE, p.getSpecies().Types) {
		return false
	}
	return true
}

func (c Freeze) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	if random.Float64() < 0.2 {
		p.clearNonVolatile()
		log.add(nonVolatileConditionLog{Index: index, Condition: nil})
		log.f("%s thaws!", p.Name)
		return false
	}
	log.f("%s is frozen solid!", p.Name)
	return true
}

type Paralysis struct {
	nonVolatileCondition
}

func (c Paralysis) name() string {
	return "paralyzed"
}
func (c Paralysis) initMessage() string {
	return "%s is paralyzed!"
}

func (c Paralysis) applicable(p *Pokemon) bool {
	if typeContains(ELECTRIC, p.getSpecies().Types) {
		return false
	}
	return true
}

func (c Paralysis) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	if random.Float64() < 0.25 {
		log.f("%s is paralyzed! It can't move!", p.Name)
		return true
	}
	return false
}

type Poison struct {
	nonVolatileCondition
}

func (c Poison) name() string {
	return "poisoned"
}
func (c Poison) initMessage() string {
	return "%s is poisoned!"
}

func (c Poison) applicable(p *Pokemon) bool {
	types := p.getSpecies().Types
	if typeContains(POISON, types) || typeContains(STEEL, types) {
		return false
	}
	return true
}

func (c Poison) postMoveEffect(log *Logger, p *Pokemon) {
	damage := int(math.Ceil(float64(p.Stats.hp) / 8))
	p.TakeDamage(damage)
	log.f("%s is hurt by poison!", p.Name)
}

type BadPoison struct {
	nonVolatileCondition
	counter int
}

func NewBadPoison() *BadPoison {
	return &BadPoison{
		counter: 1,
	}
}

func (c *BadPoison) name() string {
	return "poisoned"
}
func (c *BadPoison) initMessage() string {
	return "%s is badly poisoned!"
}

func (c *BadPoison) applicable(p *Pokemon) bool {
	types := p.getSpecies().Types
	if typeContains(POISON, types) || typeContains(STEEL, types) {
		return false
	}
	return true
}

func (c *BadPoison) postMoveEffect(log *Logger, p *Pokemon) {
	damage := int(math.Ceil(float64(p.Stats.hp) / 16))
	p.TakeDamage(damage * c.counter)
	c.counter += 1
	log.f("%s is hurt by poison!", p.Name)
}

type Sleep struct {
	nonVolatileCondition
	counter int
}

func NewSleep() *Sleep {
	return &Sleep{
		// random integer [1,2,3]
		counter: random.Intn(3) + 1,
	}
}

func (c *Sleep) name() string {
	return "asleep"
}
func (c *Sleep) initMessage() string {
	return "%s fell asleep!"
}

func (c *Sleep) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	c.counter -= 1
	if c.counter == 0 {
		p.clearNonVolatile()
		log.add(nonVolatileConditionLog{Index: index, Condition: nil})
		log.f("%s woke up!", p.Name)
		return false
	}
	log.f("%s is fast asleep!", p.Name)
	return true
}

// TODO
type VolatileCondition struct {}