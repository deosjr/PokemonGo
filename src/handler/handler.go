package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deosjr/PokemonGo/src/model"
	"github.com/deosjr/PokemonGo/src/singleplayer"
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
	sources  map[string]struct{}
	commands []model.MoveCommand
	channels []chan []byte
}

func (h *Handler) addBattle(gi gameinput, b model.Battle) {
	bi := battleInfo{
		battle:  b,
		p1:      gi.P1,
		p2:      gi.P2,
		sources: map[string]struct{}{},
	}
	h.battles[gi.Name] = &bi
}

type request struct {
	input           input
	responseChannel chan []byte
}

type input struct {
	GameName   string `json:"game"`
	PlayerName string `json:"player"`
	//TargetIndex int    `json:"target"`
	MoveIndex int `json:"move"`
}

func (h *Handler) HandleMove(w http.ResponseWriter, r *http.Request) {
	var i input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make(chan []byte, 1)
	req := request{i, resp}
	err = h.handle(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Game %s: %s issued move %d \n", i.GameName, i.PlayerName, i.MoveIndex)

	turnResolution := <-resp
	cors(w)
	w.Write(turnResolution)
}

func (h *Handler) handle(req request) error {
	gamename := req.input.GameName
	playername := req.input.PlayerName
	battleInfo, ok := h.battles[gamename]
	if !ok {
		return fmt.Errorf("game %s does not exist", gamename)
	}
	if _, ok := battleInfo.sources[playername]; ok {
		return fmt.Errorf("player %s already issues command this turn", playername)
	}

	battleInfo.sources[playername] = struct{}{}
	sourceIndex := 0
	targetIndex := 1
	if playername == battleInfo.p2 {
		sourceIndex = 1
		targetIndex = 0
	}
	moveCommand := model.MoveCommand{
		SourceIndex: sourceIndex,
		TargetIndex: targetIndex,
		MoveIndex:   req.input.MoveIndex,
	}
	battleInfo.commands = append(battleInfo.commands, moveCommand)
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
	battleInfo.sources = map[string]struct{}{}
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

	p1 := singleplayer.GetRentalPokemon()
	p2 := singleplayer.GetRentalPokemon()
	battle := model.NewSingleBattle(p1, p2)
	h.addBattle(gi, battle)

	cors(w)
	w.WriteHeader(200)
}

type joininput struct {
	Name   string `json:"name"`
	Player string `json:"player"`
}

type joinoutput struct {
	SpeciesSelf int       `json:"self"`
	SpeciesOpp  int       `json:"opp"`
	Moves       [4]string `json:"moves"`
}

func (h *Handler) HandleJoinGame(w http.ResponseWriter, r *http.Request) {
	var ji joininput
	err := json.NewDecoder(r.Body).Decode(&ji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bi, ok := h.battles[ji.Name]
	if !ok {
		http.Error(w, "game not found", http.StatusBadRequest)
		return
	}

	var resp joinoutput
	switch {
	case ji.Player == bi.p1:
		gi := bi.battle.Graphics(0)
		resp.SpeciesSelf = gi.Sides[0][0]
		resp.SpeciesOpp = gi.Sides[1][0]
		resp.Moves = gi.Moves
	case ji.Player == bi.p2:
		gi := bi.battle.Graphics(1)
		resp.SpeciesSelf = gi.Sides[0][0]
		resp.SpeciesOpp = gi.Sides[1][0]
		resp.Moves = gi.Moves
	default:
		w.WriteHeader(403)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cors(w)
	w.Write(respJSON)
}

//TODO: check how to do this properly
func cors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}
