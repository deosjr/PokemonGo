package model

import (
	"math"
)

type NonVolatileCondition interface {
	name() string
	initMessage() string
	applicable(*Pokemon) bool
	preMoveEffect(*Logger, *Pokemon, int) (cantAttack bool)
	postMoveEffect(*Logger, int, *Pokemon)
}

type nonVolatileCondition struct{}

func (c nonVolatileCondition) name() string                              { return "" }
func (c nonVolatileCondition) initMessage() string                       { return "" }
func (c nonVolatileCondition) applicable(*Pokemon) bool                  { return true }
func (c nonVolatileCondition) preMoveEffect(*Logger, *Pokemon, int) bool { return false }
func (c nonVolatileCondition) postMoveEffect(*Logger, int, *Pokemon)     {}

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

func (c Burn) postMoveEffect(log *Logger, i int, p *Pokemon) {
	// TODO: hp / 16 since Gen VII
	damage := int(math.Ceil(float64(p.Stats.hp) / 8))
	damageTaken := p.TakeDamage(damage)
	log.damage(i, damageTaken)
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
		log.add(NonVolatileConditionLog{Index: index, Condition: nil})
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

func (c Poison) postMoveEffect(log *Logger, i int, p *Pokemon) {
	damage := max(1, p.Stats.hp/8)
	damageTaken := p.TakeDamage(damage)
	log.damage(i, damageTaken)
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

func (c *BadPoison) postMoveEffect(log *Logger, i int, p *Pokemon) {
	damage := max(1, p.Stats.hp/16)
	damageTaken := p.TakeDamage(damage * c.counter)
	log.damage(i, damageTaken)
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
		log.add(NonVolatileConditionLog{Index: index, Condition: nil})
		log.f("%s woke up!", p.Name)
		return false
	}
	log.f("%s is fast asleep!", p.Name)
	return true
}

type VolatileConditionLabel int

const (
	ConfusionLabel VolatileConditionLabel = iota
	FlinchLabel
    ProtectCounterLabel
    ProtectLabel
)

type VolatileCondition interface {
	getLabel() VolatileConditionLabel
	initMessage() string
	failMessage() string
	preMoveEffect(*Logger, *Pokemon, int) (cantAttack bool)
	postMoveEffect(*Pokemon)
}

type volatileCondition struct {
	label VolatileConditionLabel
}

func (v volatileCondition) getLabel() VolatileConditionLabel {
	return v.label
}

func (v volatileCondition) initMessage() string                       { return "" }
func (v volatileCondition) failMessage() string                       { return "" }
func (v volatileCondition) preMoveEffect(*Logger, *Pokemon, int) bool { return false }
func (v volatileCondition) postMoveEffect(*Pokemon)                   {}

type Confusion struct {
	volatileCondition
	counter int
}

func NewConfusion() *Confusion {
	return &Confusion{
		volatileCondition: volatileCondition{
			label: ConfusionLabel,
		},
		// random integer [1,2,3,4]
		counter: random.Intn(4) + 1,
	}
}

func (c *Confusion) initMessage() string {
	return "%s became confused!"
}

func (c *Confusion) failMessage() string {
	return "%s is already confused!"
}

func (c *Confusion) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	if c.counter == 0 {
		p.clearVolatile(c.label)
		log.add(VolatileConditionLog{Index: index, Label: c.label, Condition: nil})
		log.f("%s snapped out of its confusion!", p.Name)
		return false
	}
	log.f("%s is confused!", p.Name)
	cantAttack := false
	if random.Float64() < 0.5 {
		cantAttack = true
		log.f("It hurt itself in its confusion!")
		// TODO: attack itself w 40power notype physical attack, never crits
		typelessMove := MoveData{Power: 40, Type: NOTYPE, Category: physical}
		dmg, _, _ := dealDamage(p, p, typelessMove)
		log.damage(index, dmg)
		p.TakeDamage(dmg)
	}
	c.counter -= 1
	return cantAttack
}

type Flinch struct {
	volatileCondition
}

func NewFlinch() Flinch {
	return Flinch{volatileCondition{FlinchLabel}}
}

// TODO: p only flinches if hit with flinch this turn before p could attack
func (f Flinch) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	p.clearVolatile(f.label)
	log.f("%s flinched!", p.Name)
	return true
}

func (f Flinch) postMoveEffect(p *Pokemon) {
	p.clearVolatile(f.label)
}

// shared by Protect, Detect etc
type ProtectCounter struct {
    volatileCondition
    counter int
}

func NewProtectCounter() *ProtectCounter {
    return &ProtectCounter{
        volatileCondition:volatileCondition{ProtectCounterLabel},
        counter: 0,
    }
}

func (c *ProtectCounter) Success() bool {
    return random.Float64() < math.Pow(1.0/3.0, float64(c.counter))
}

type Protect struct {
	volatileCondition
}

func NewProtect() Protect {
	return Protect{volatileCondition{ProtectLabel}}
}

func (p Protect) initMessage() string {
    return "%s protected itself!"
}

func (f Protect) preMoveEffect(log *Logger, p *Pokemon, index int) bool {
	p.clearVolatile(f.label)
    return false
}
