package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "sync"

	"model"
)

var (
	Battle         *model.Battle
	CommandChannel chan model.Command
	TurnChannel    chan []string
)

type input struct {
	Player  string        `json:"player"`
	Command model.Command `json:"command"`
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Requested %s %v\n", r.URL.Path, r.FormValue("input"))

	var i input
	err := json.Unmarshal([]byte(r.FormValue("input")), &i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Issued command %+v \n", i.Command)
	// TODO: command validation
	CommandChannel <- i.Command

	// CORS stuff
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	for {
		select {
		case turnResolution := <-TurnChannel:
			json.NewEncoder(w).Encode(turnResolution)
			return
		}
	}
}
