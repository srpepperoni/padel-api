package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"k8s.io/klog/v2"

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

// Create
// @Summary Create new player
// @Description create new player
// @Tags Players
// @Accept  json
// @Param player body models.PlayerJSON true "Player object for API"
// @Produce  json
// @Success 201 {object} models.Player
// @Router /player [post]
func (h playersHandlers) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		klog.Fatalln(err)
	}

	if _, err = h.playersUC.Create(body); err != nil {
		klog.Errorf("Error creating player: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

// Update
// @Summary Update player
// @Description update player
// @Tags Players
// @Accept  json
// @Param player body models.PlayerJSON true "Player object for API"
// @Param id path int true "Player ID"
// @Produce  json
// @Success 201 {object} models.Player
// @Router /player/{id} [put]
func (h playersHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		klog.Fatalln(err)
	}

	_, err = h.playersUC.Update(body, id)
	if err != nil {
		klog.Errorf("Error updating player: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

// Delete
// @Summary Delete player
// @Description delete player
// @Tags Players
// @Accept  json
// @Param id path int true "Player ID"
// @Produce  json
// @Success 201 {object} models.Player
// @Router /player/{id} [delete]
func (h playersHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := h.playersUC.Delete(id)
	if err != nil {
		klog.Errorf("Error deleting player: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

// Get All Players
// @Summary Get All Players
// @Description get all players
// @Tags Players
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Player
// @Router /players [get]
func (h playersHandlers) GetPlayers(w http.ResponseWriter, r *http.Request) {
	var players *[]models.Player
	var err error

	if players, err = h.playersUC.GetPlayers(); err != nil {
		klog.Errorf("Error getting players: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&players)
}

// Get By Id
// @Summary Get Player by Id
// @Description get one player by Id
// @Tags Players
// @Accept  json
// @Param id path int true "Player ID"
// @Produce  json
// @Success 201 {object} models.Player
// @Router /player/{id} [get]
func (h playersHandlers) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	player, err := h.playersUC.GetPlayer(id)

	if err != nil {
		klog.Errorf("Error getting player: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}
