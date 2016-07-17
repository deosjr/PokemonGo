package model

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	dataPath    = "../data/"
	pokemonPath = "pokemon.txt"
	movePath    = "moves.txt"
)

func MustLoadConfig() {
	absPath, _ := filepath.Abs(dataPath)
	initTypeMap()
	mustLoadMoves(absPath + movePath)
	mustLoadSpecies(absPath + pokemonPath)

	rand.Seed(time.Now().UnixNano())
}

func mustLoadSpecies(path string) {
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
			pokemonData = append(pokemonData, &Species{ID: i + 1})
			i++
			continue
		}
		assignment := strings.Split(line, "=")
		pokemonData[i-1].update(assignment[0], assignment[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func mustLoadMoves(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		features := strings.Split(line, ",")
		LoadMove(features)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func atoiTrusted(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
