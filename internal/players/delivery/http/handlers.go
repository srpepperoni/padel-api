package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
	"github.com/gorilla/mux"
)

type playersHandlers struct {
	playersUC players.UseCase
}

func NewPlayersHandlers(playersUC players.UseCase) players.Handlers {
	return &playersHandlers{playersUC: playersUC}
}

func (h playersHandlers) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var playerJSON models.PlayerJSON
	json.Unmarshal(body, &playerJSON)

	player := models.NewPlayer(playerJSON.Name, playerJSON.LastName, playerJSON.PlayerName)

	if _, err := h.playersUC.Create(&player); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h playersHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedPlayer models.PlayerJSON
	json.Unmarshal(body, &updatedPlayer)

	player := models.NewPlayer(updatedPlayer.Name, updatedPlayer.LastName, updatedPlayer.PlayerName)

	_, err = h.playersUC.Update(&player, id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h playersHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	h.playersUC.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h playersHandlers) GetPlayers(w http.ResponseWriter, r *http.Request) {
	var players *[]models.Player
	var err error

	if players, err = h.playersUC.GetPlayers(); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&players)
}

func (h playersHandlers) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var player *models.Player
	var err error

	if player, err = h.playersUC.GetPlayer(id); err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}
