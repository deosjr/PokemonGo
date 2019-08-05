package main

import (
	"fmt"
	"net/http"

	"github.com/deosjr/PokemonGo/src/handler"
	"github.com/deosjr/PokemonGo/src/model"
)

func main() {
	model.MustLoadConfig()
	fmt.Println("Serving..")

	h := handler.NewHandler()
	http.HandleFunc("/game", h.HandleNewGame)
	http.HandleFunc("/join", h.HandleJoinGame)
	http.HandleFunc("/move", h.HandleMove)
	http.ListenAndServe(":8080", nil)

}
