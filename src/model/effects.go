package model

import (
	"math"
)

// change one or more stat stages [attack, defense, spattack, spdefense, speed]
func changeStatStages(log *Logger, p *Pokemon, index int, changes Stats) {
	effectiveChanges, maxed := p.ChangeStatStages(changes)
	log.statStages(p.Name, index, effectiveChanges, maxed)
}

// change accuracy stat stage
func changeAccuracy(log *Logger, p *Pokemon, index int, change int) {
	effectiveChanges, maxed := p.ChangeAccuracy(change)
	log.add(StatStageLog{Index: index, Accuracy: effectiveChanges})
	log.statStage(p.Name, index, "accuracy", effectiveChanges, maxed)
}

// change evasion stat stage
func changeEvasion(log *Logger, p *Pokemon, index int, change int) {
	effectiveChanges, maxed := p.ChangeEvasion(change)
	log.add(StatStageLog{Index: index, Evasion: effectiveChanges})
	log.statStage(p.Name, index, "evasion", effectiveChanges, maxed)
}

// attempt to inflict a non-volatile condition
// fails if the target already has a non-volatile condition
func inflictNonVolatileCondition(log *Logger, p *Pokemon, index int, condition NonVolatileCondition) {
	if !condition.applicable(p) {
		return
	}
	success := p.setNonVolatile(condition)
	log.nonVolatileCondition(p.Name, index, success, condition)
}

func inflictVolatileCondition(log *Logger, p *Pokemon, index int, condition VolatileCondition) {
	success := p.addVolatile(condition)
	log.volatileCondition(p.Name, index, success, condition)
}

// TODO: Fails if its HP is full. Cannot be used if the user is affected by Heal Block.
func healWithFail(log *Logger, p *Pokemon, index, amount int) {
	actualHealed := p.Heal(amount)
	log.f("%s regained health!", p.Name)
	log.heal(index, actualHealed)
}

// drain heals the user for half the HP the target lost (min 1)
// TODO: Gen 6: Cannot be used if the user is affected by Heal Block.
func drain(log *Logger, p *Pokemon, index, dmg int, name string) {
	actualHealed := p.Heal(max(1, dmg/2))
	log.f("%s had its energy drained!", name)
	log.heal(index, actualHealed)
}

func recoil(log *Logger, p *Pokemon, index, dmg, factor int) {
	recoilDamage := max(1, int(math.Ceil(float64(dmg)/float64(factor))))
	log.damage(index, recoilDamage)
	p.TakeDamage(recoilDamage)
}

// deal fixed amount of damage unless target is immune to move type
func fixedDamage(damage int, typ pType, target *Pokemon) (int, float64, float64) {
	t := typeEffectiveness(typ, target.getSpecies().Types)
	if t == 0 {
		return 0, 0, 1
	}
	return damage, 1, 1
}
