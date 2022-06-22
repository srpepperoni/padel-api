package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"github.com/gorilla/mux"
)

type matchesHandlers struct {
	matchesUC matches.UseCase
}

func NewMatchesHandlers(matchesUC matches.UseCase) matches.Handlers {
	return &matchesHandlers{matchesUC: matchesUC}
}

func (h matchesHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	match := models.NewMatch(matchJSON.CoupleOne[0], matchJSON.CoupleOne[1], matchJSON.CoupleTwo[0], matchJSON.CoupleTwo[1], matchJSON.Status, matchJSON.TournamentID, result)

	if _, err := h.matchesUC.Create(&match); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h matchesHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	updatedMatch := models.NewMatch(matchJSON.CoupleOne[0], matchJSON.CoupleOne[1], matchJSON.CoupleTwo[0], matchJSON.CoupleTwo[1], matchJSON.Status, matchJSON.TournamentID, result)

	_, err = h.matchesUC.Update(&updatedMatch, id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h matchesHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	h.matchesUC.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h matchesHandlers) GetMatches(w http.ResponseWriter, r *http.Request) {
	var matches *[]models.Match
	var err error

	if matches, err = h.matchesUC.GetMatches(); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&matches)
}

func (h matchesHandlers) GetMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var player *models.Match
	var err error

	if player, err = h.matchesUC.GetMatch(id); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (h matchesHandlers) GetMatchesByTournamentId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var matches *[]models.Match
	var err error

	if matches, err = h.matchesUC.GetMatchesByTournamentId(id); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
