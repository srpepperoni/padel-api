package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/tournaments"
	"github.com/gorilla/mux"
)

type tournamentsHandlers struct {
	tournamentsUC tournaments.UseCase
}

func NewTournamentsHandlers(tournamentsUC tournaments.UseCase) tournaments.Handlers {
	return &tournamentsHandlers{tournamentsUC: tournamentsUC}
}

// Create
// @Summary Create new tournament
// @Description create new tournament
// @Tags Tournament
// @Accept  json
// @Param tournament body models.TournamentJSON true "Tournament object for API"
// @Produce  json
// @Success 201 {object} models.Tournament
// @Router /tournament [post]
func (h tournamentsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds)

	if _, err := h.tournamentsUC.Create(&tournament); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

// Update
// @Summary Update tournament
// @Description update tournament
// @Tags Tournament
// @Accept  json
// @Param tournament body models.TournamentJSON true "Tournament object for API"
// @Param id path int true "Tournament ID"
// @Produce  json
// @Success 201 {object} models.Tournament
// @Router /tournament/{id} [put]
func (h tournamentsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	updatedTournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds)

	_, err = h.tournamentsUC.Update(&updatedTournament, id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

// Delete
// @Summary delete tournament
// @Description delete tournament
// @Tags Tournament
// @Accept  json
// @Param id path int true "Tournament ID"
// @Produce  json
// @Success 201 {object} models.Tournament
// @Router /tournament/{id} [delete]
func (h tournamentsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	h.tournamentsUC.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

// Get all
// @Summary Get All tournament
// @Description update tournament
// @Tags Tournament
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Tournament
// @Router /tournaments [get]
func (h tournamentsHandlers) GetTournaments(w http.ResponseWriter, r *http.Request) {
	var tournaments *[]models.Tournament
	var err error

	if tournaments, err = h.tournamentsUC.GetTournaments(); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tournaments)
}

// Get
// @Summary Get tournament by id
// @Description get tournament by id
// @Tags Tournament
// @Accept  json
// @Param id path int true "Tournament ID"
// @Produce  json
// @Success 201 {object} models.Tournament
// @Router /tournament/{id} [get]
func (h tournamentsHandlers) GetTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var player *models.Tournament
	var err error

	if player, err = h.tournamentsUC.GetTournament(id); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}
