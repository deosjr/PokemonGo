package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"model"
)

var (
	battle         model.Battle
	requestChannel chan request
)

type request struct {
	Command         model.Command
	ResponseChannel chan int
}

type input struct {
	Player  string        `json:"player"`
	Command model.Command `json:"command"`
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	var i input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Player %s issued command %+v \n", i.Player, i.Command)
	// TODO: command validation

	resp := make(chan int)
	req := request{i.Command, resp}
	requestChannel <- req

	// CORS Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	resolvedTurn := <-resp
	turnResolution := battle.Log().Logs()[resolvedTurn]
	json.NewEncoder(w).Encode(turnResolution)
}

func Handle(b model.Battle) {
	var commands []model.Command
	var channels []chan int
	battle = b
	requestChannel = make(chan request)

	for {
		select {
		case req := <-requestChannel:
			commands = append(commands, req.Command)
			channels = append(channels, req.ResponseChannel)
			if len(commands) != 2 {
				break
			}
			err := model.HandleTurn(battle, commands)
			if err != nil {
				panic(err)
			}
			for _, r := range channels {
				r <- battle.Log().Turn() - 1
			}
			commands = []model.Command{}
			channels = []chan int{}
		}
	}
}
