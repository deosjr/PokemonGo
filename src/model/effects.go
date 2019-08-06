package model

// change one or more stat stages [attack, defense, spattack, spdefense, speed]
func changeStatStages(log *Logger, p *Pokemon, index int, changes Stats) {
	effectiveChanges, maxed := p.ChangeStatStages(changes)
	log.statStages(p.Name, index, effectiveChanges, maxed)
}

// change accuracy stat stage
func changeAccuracy(log *Logger, p *Pokemon, index int, change int) {
	effectiveChanges, maxed := p.ChangeAccuracy(change)
	log.add(statStageLog{Index: index, Accuracy: effectiveChanges})
	log.statStage(p.Name, index, "accuracy", effectiveChanges, maxed)
}

// change evasion stat stage
func changeEvasion(log *Logger, p *Pokemon, index int, change int) {
	effectiveChanges, maxed := p.ChangeEvasion(change)
	log.add(statStageLog{Index: index, Evasion: effectiveChanges})
	log.statStage(p.Name, index, "evasion", effectiveChanges, maxed)
}

// attempt to inflict a non-volatile condition
// fails if the target already has a non-volatile condition
func inflictNonVolatileCondition(log *Logger, p *Pokemon, index int, condition NonVolatileCondition) {
	success := p.setNonVolatile(condition)
	log.nonVolatileCondition(p.Name, index, success, condition)
}

// TODO: Fails if its HP is full. Cannot be used if the user is affected by Heal Block.
func healWithFail(log *Logger, p *Pokemon, amount int) {
	p.Heal(amount)
	log.f("%s regained health!", p.Name)
}

// drain heals the user for half the HP the target lost (min 1)
// TODO: Gen 6: Cannot be used if the user is affected by Heal Block.
func drain(log *Logger, p *Pokemon, dmg int, name string) {
	p.Heal(max(1, dmg/2))
	log.f("%s had its energy drained!", name)
}

// deal fixed amount of damage unless target is immune to move type
func fixedDamage(damage int, typ pType, target *Pokemon) (int, float64, float64) {
	t := typeEffectiveness(typ, target.getSpecies().Types)
	if t == 0 {
		return 0, 0, 1
	}
	return damage, 1, 1
}
