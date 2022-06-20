package http

import (
	"net/http"

	"fake.com/padel-api/internal/players"
	"github.com/gorilla/mux"
)

func MapPlayersRoutes(router *mux.Router, playersHandlers players.Handlers) {
	router.HandleFunc("/players", playersHandlers.GetPlayers).Methods(http.MethodGet)
	router.HandleFunc("/player", playersHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/player/{id}", playersHandlers.GetPlayer).Methods(http.MethodGet)
}
