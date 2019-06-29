package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deosjr/PokemonGo/src/model"
)

type Handler struct {
	battles map[string]*battleInfo
}

func NewHandler() *Handler {
	return &Handler{
		battles: make(map[string]*battleInfo),
	}
}

type battleInfo struct {
	battle   model.Battle
	sources  map[int]struct{}
	commands []model.Command
	channels []chan []byte
}

func (h *Handler) AddBattle(name string, b model.Battle) {
	bi := battleInfo{
		battle:  b,
		sources: map[int]struct{}{},
	}
	h.battles[name] = &bi
}

type request struct {
	command         model.Command
	responseChannel chan []byte
}

type input struct {
	Game    string        `json:"game"`
	Player  string        `json:"player"`
	Command model.Command `json:"command"`
}

func (h *Handler) HandleMove(w http.ResponseWriter, r *http.Request) {
	var i input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make(chan []byte, 1)
	req := request{i.Command, resp}
	err = h.handle(i.Game, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Game %s: %s issued command %+v \n", i.Game, i.Player, i.Command)

	// CORS Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	turnResolution := <-resp
	w.Write(turnResolution)
}

func (h *Handler) handle(game string, req request) error {
	battleInfo, ok := h.battles[game]
	if !ok {
		return fmt.Errorf("game %s does not exist", game)
	}
	if _, ok := battleInfo.sources[req.command.SourceIndex]; ok {
		return fmt.Errorf("source %s already issues command this turn", req.command.SourceIndex)
	}

	battleInfo.sources[req.command.SourceIndex] = struct{}{}
	battleInfo.commands = append(battleInfo.commands, req.command)
	battleInfo.channels = append(battleInfo.channels, req.responseChannel)

	if len(battleInfo.commands) != 2 {
		return nil
	}

	err := model.HandleTurn(battleInfo.battle, battleInfo.commands)
	if err != nil {
		return err
	}

	resolvedTurn := battleInfo.battle.Log().Turn() - 1
	turnResolution := battleInfo.battle.Log().Logs()[resolvedTurn]
	turnJSON, err := json.Marshal(turnResolution)
	if err != nil {
		return err
	}
	for _, r := range battleInfo.channels {
		r <- turnJSON
	}
	battleInfo.sources = map[int]struct{}{}
	battleInfo.commands = nil
	battleInfo.channels = nil
	return nil
}
