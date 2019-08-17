package model

type Pokemon struct {
	Level     int
	Species   pokemon
	Name      string
	XP        XP
	currentHP int

	Stats         Stats
	iv            Stats
	statStages    Stats
	accuracyStage int
	evasionStage  int

	Moves                []*Move
	NonVolatileCondition NonVolatileCondition
	VolatileConditions   map[VolatileConditionLabel]VolatileCondition
}

type XP struct {
	PreviousXPLevelReq int
	CurrentXP          int
	NextXPLevelReq     int
}

func GetPokemonByName(level int, name string) *Pokemon {
	p := pokemon(GetSpecies(name).ID - 1)
	return GetPokemon(level, p)
}

func GetPokemon(level int, pType pokemon) *Pokemon {
	ivs := generateIVs()
	species := GetSpeciesByID(pType)
	stats := calculateStats(ivs, species.Stats, level)

	return &Pokemon{
		Level:              level,
		Species:            pType,
		Name:               species.Name,
		Stats:              stats,
		iv:                 ivs,
		currentHP:          stats.hp,
		statStages:         GetStats([6]int{}),
		Moves:              make([]*Move, 4),
		VolatileConditions: map[VolatileConditionLabel]VolatileCondition{},
	}
}

func (p *Pokemon) getSpecies() Species {
	return GetSpeciesByID(p.Species)
}

func (p *Pokemon) GetTotalHP() int {
	return p.Stats.hp
}

func (p *Pokemon) TakeDamage(damage int) int {
	temp := p.currentHP
	p.currentHP = max(0, p.currentHP-damage)
	return temp - p.currentHP
}
func (p *Pokemon) Heal(n int) int {
	temp := p.currentHP
	p.currentHP = min(p.Stats.hp, p.currentHP+n)
	return p.currentHP - temp
}

func (p *Pokemon) ChangeStatStages(changes Stats) (Stats, [6]bool) {
	newStages, effectiveChanges, maxed := p.statStages.updateStages(changes)
	p.statStages = newStages
	return effectiveChanges, maxed
}
func (p *Pokemon) ChangeAccuracy(change int) (effective int, maxed bool) {
	p.accuracyStage, effective, maxed = validStatStage(p.accuracyStage, change)
	return effective, maxed
}
func (p *Pokemon) ChangeEvasion(change int) (effective int, maxed bool) {
	p.evasionStage, effective, maxed = validStatStage(p.evasionStage, change)
	return effective, maxed
}

// Returns modified stat rounded down to an int
func (p *Pokemon) Attack() int {
	attack := int(modifyStat(p.Stats.attack, p.statStages.attack))
	if p.NonVolatileCondition == nil || p.NonVolatileCondition.name() == "burned" {
		// TODO: (since Gen III) Instead of modifying the Attack stat,
		// a burn now technically halves the damage a burned Pokémon does with physical moves;
		// it still does not reduce damage done by moves that deal direct damage (such as Seismic Toss)
		return attack / 2
	}
	return attack
}
func (p *Pokemon) Defense() int {
	return int(modifyStat(p.Stats.defense, p.statStages.defense))
}
func (p *Pokemon) SpAttack() int {
	return int(modifyStat(p.Stats.spattack, p.statStages.spattack))
}
func (p *Pokemon) SpDefense() int {
	return int(modifyStat(p.Stats.spdefense, p.statStages.spdefense))
}
func (p *Pokemon) Speed() int {
	speed := int(modifyStat(p.Stats.speed, p.statStages.speed))
	if p.NonVolatileCondition == nil || p.NonVolatileCondition.name() == "paralyzed" {
		return speed / 4 // speed / 2 since Gen VII
	}
	return speed
}

// These stats are handled a little differently
func (p *Pokemon) Accuracy() float64 {
	return accuracyOrEvasionToMod(p.accuracyStage)
}
func (p *Pokemon) Evasion() float64 {
	return accuracyOrEvasionToMod(p.evasionStage)
}

func (p *Pokemon) setNonVolatile(c NonVolatileCondition) (succes bool) {
	if p.NonVolatileCondition != nil {
		return false
	}
	p.NonVolatileCondition = c
	return true
}

func (p *Pokemon) clearNonVolatile() {
	p.NonVolatileCondition = nil
}

func (p *Pokemon) addVolatile(c VolatileCondition) (succes bool) {
	_, ok := p.VolatileConditions[c.getLabel()]
	if ok {
		return false
	}
	p.VolatileConditions[c.getLabel()] = c
	return true
}

func (p *Pokemon) clearVolatile(label VolatileConditionLabel) {
	delete(p.VolatileConditions, label)
}

func (p *Pokemon) clearAllVolatile() {
	p.VolatileConditions = map[VolatileConditionLabel]VolatileCondition{}
}
