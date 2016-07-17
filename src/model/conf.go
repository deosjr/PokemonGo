package model

import (
    "bufio"
    "math/rand"
    "os"
    "path/filepath"
    "strings"
    "time"
)

var (
    dataPath = "../data/"
    pokemonPath = "pokemon.txt"
)

func MustLoadConfig() {
    absPath, _ := filepath.Abs(dataPath)
    initTypeMap()
    pokemonData = mustLoadSpecies(absPath + pokemonPath)

    rand.Seed(time.Now().UnixNano())
}

func mustLoadSpecies(path string) []*Species {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var i int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "[") {
            pokemonData = append(pokemonData, &Species{ID:i+1})
            i++
            continue
        }
        assignment := strings.Split(line, "=")
        pokemonData[i-1].update(assignment[0], assignment[1])
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    return pokemonData
}