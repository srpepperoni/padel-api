package http

import (
	"net/http"

	"fake.com/padel-api/internal/matches"
	"github.com/gorilla/mux"
)

func MapMatchesRoutes(router *mux.Router, matchesHandlers matches.Handlers) {
	router.HandleFunc("/matches", matchesHandlers.GetMatches).Methods(http.MethodGet)
	router.HandleFunc("/match", matchesHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/match/{id}", matchesHandlers.GetMatch).Methods(http.MethodGet)
	router.HandleFunc("/match", matchesHandlers.Update).Methods(http.MethodPut)
}
