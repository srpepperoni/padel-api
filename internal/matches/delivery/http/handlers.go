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

	var player models.Match
	json.Unmarshal(body, &player)

	if _, err := h.matchesUC.Create(&player); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
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
	id, _ := strconv.Atoi(vars["playerID"])

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

func (h matchesHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedMatch models.Match
	json.Unmarshal(body, &updatedMatch)

	_, err = h.matchesUC.Update(&updatedMatch, id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
