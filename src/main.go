package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/deosjr/PokemonGo/src/handler"
	"github.com/deosjr/PokemonGo/src/model"
	"github.com/deosjr/PokemonGo/src/singleplayer"
)

var serve bool

func init() {
	flag.BoolVar(&serve, "serve", false, "run as server")
	flag.Parse()
}

func main() {
	model.MustLoadConfig()

	if serve {
		fmt.Println("Serving..")

		h := handler.NewHandler()
		http.HandleFunc("/game", h.HandleNewGame)
		http.HandleFunc("/join", h.HandleJoinGame)
		http.HandleFunc("/move", h.HandleMove)
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Starting single player battle")
		singleplayer.Repl()
	}
}
