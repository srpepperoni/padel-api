package http

import (
	"net/http"

	"fake.com/padel-api/internal/matches"
	"github.com/gorilla/mux"
)

func MapMatchesRoutes(router *mux.Router, matchesHandlers matches.Handlers) {
	router.HandleFunc("/match", matchesHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/match/{id}", matchesHandlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/match/{id}", matchesHandlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/matches", matchesHandlers.GetMatches).Methods(http.MethodGet)
	router.HandleFunc("/match/{id}", matchesHandlers.GetMatch).Methods(http.MethodGet)
	router.HandleFunc("/tournament/match/{id}", matchesHandlers.GetMatchesByTournamentId).Methods(http.MethodGet)
	router.HandleFunc("/match/{id}/result", matchesHandlers.SetResult).Methods(http.MethodPost)
}
