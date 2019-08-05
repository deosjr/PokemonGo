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
	p1       string
	p2       string
	sources  map[int]struct{}
	commands []model.Command
	channels []chan []byte
}

func (h *Handler) addBattle(gi gameinput, b model.Battle) {
	bi := battleInfo{
		battle:  b,
		p1:      gi.P1,
		p2:      gi.P2,
		sources: map[int]struct{}{},
	}
	h.battles[gi.Name] = &bi
}

type request struct {
	command         model.Command
	responseChannel chan []byte
}

type input struct {
	Game    string        `json:"game"`
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

	fmt.Printf("Game %s: P%d issued command %+v \n", i.Game, i.Command.SourceIndex+1, i.Command)

	turnResolution := <-resp
	cors(w)
	w.Write(turnResolution)
}

func (h *Handler) handle(game string, req request) error {
	battleInfo, ok := h.battles[game]
	if !ok {
		return fmt.Errorf("game %s does not exist", game)
	}
	if _, ok := battleInfo.sources[req.command.SourceIndex]; ok {
		return fmt.Errorf("source %d already issues command this turn", req.command.SourceIndex)
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

type gameinput struct {
	Name string `json:"name"`
	P1   string `json:"p1"`
	P2   string `json:"p2"`
}

func (h *Handler) HandleNewGame(w http.ResponseWriter, r *http.Request) {
	var gi gameinput
	err := json.NewDecoder(r.Body).Decode(&gi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := h.battles[gi.Name]; ok {
		http.Error(w, "game already exists", http.StatusBadRequest)
		return
	}

	p1 := model.GetPokemonByName(10, "bulbasaur")
	p2 := model.GetPokemonByName(10, "charmander")
	p1.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.TACKLE)}
	p2.Moves[0] = &model.Move{Data: model.GetMoveDataByID(model.EMBER)}
	battle := model.NewSingleBattle(p1, p2)
	h.addBattle(gi, battle)

	cors(w)
	w.WriteHeader(200)
}

type joininput struct {
	Name   string `json:"name"`
	Player string `json:"player"`
}

func (h *Handler) HandleJoinGame(w http.ResponseWriter, r *http.Request) {
	var ji joininput
	err := json.NewDecoder(r.Body).Decode(&ji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := h.battles[ji.Name]; !ok {
		http.Error(w, "game not found", http.StatusBadRequest)
		return
	}

	//TODO

	cors(w)
	w.WriteHeader(200)
}

//TODO: check how to do this properly
func cors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}
