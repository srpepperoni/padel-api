package http

import (
	"net/http"

	"fake.com/padel-api/internal/tournaments"
	"github.com/gorilla/mux"
)

func MapTournamentsRoutes(router *mux.Router, tournamentsHandlers tournaments.Handlers) {
	router.HandleFunc("/tournament", tournamentsHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/tournament/{id}", tournamentsHandlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/tournament/{id}", tournamentsHandlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/tournaments", tournamentsHandlers.GetTournaments).Methods(http.MethodGet)
	router.HandleFunc("/tournament/{id}", tournamentsHandlers.GetTournament).Methods(http.MethodGet)
	router.HandleFunc("/tournament/{id}/next-round", tournamentsHandlers.NextRound).Methods(http.MethodPost)
}
