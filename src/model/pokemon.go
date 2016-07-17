package model

import (
    
)

type Pokemon struct {
    Level int
    Species *Species
    Name string
    XP XP

    Stats Stats
    iv Stats
    statStages Stats

    //Moves Move[]
    //NonVolatileCondition Condition
    //VolatileConditions VolatileCondition[]
}

type XP struct {
    PreviousXPLevelReq int
    CurrentXP int
    NextXPLevelReq int
}

func GetPokemon(level int, species *Species) *Pokemon {
    ivs := generateIVs()

    return &Pokemon{
        Level : level,
        Species : species,
        Name : species.Name,
        Stats : calculateStats(ivs, species.Stats, level),
        iv : ivs,
        statStages : Stats{make([]int, 6)},

    }
}