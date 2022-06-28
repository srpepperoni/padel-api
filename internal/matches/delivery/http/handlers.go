package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"k8s.io/klog/v2"

	"fake.com/padel-api/internal/matches"
	"github.com/gorilla/mux"
)

type matchesHandlers struct {
	matchesUC matches.UseCase
}

func NewMatchesHandlers(matchesUC matches.UseCase) matches.Handlers {
	return &matchesHandlers{matchesUC: matchesUC}
}

// Create
// @Summary Create new match
// @Description create new match
// @Tags Matches
// @Accept  json
// @Param match body models.MatchJSON true "Match object for API"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /match [post]
func (h matchesHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		klog.Fatalln(err)
	}

	if _, err = h.matchesUC.Create(body); err != nil {
		klog.Errorf("Error creating match: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

// Update
// @Summary Update match
// @Description update match
// @Tags Matches
// @Accept  json
// @Param player body models.MatchJSON true "Match object for API"
// @Param id path int true "Match ID"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /match/{id} [put]
func (h matchesHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		klog.Fatalln(err)
	}

	if _, err = h.matchesUC.Update(body, id); err != nil {
		klog.Errorf("Error updating match: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

// Delete
// @Summary Delete match by id
// @Description delete match
// @Tags Matches
// @Accept  json
// @Param id path int true "Match ID"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /match/{id} [delete]
func (h matchesHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	h.matchesUC.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

// Get All
// @Summary Get all matches
// @Description get all matches
// @Tags Matches
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Match
// @Router /matches [get]
func (h matchesHandlers) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.matchesUC.GetMatches()

	if err != nil {
		klog.Error(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&matches)
}

// Get Match
// @Summary get match by id
// @Description get match by id
// @Tags Matches
// @Accept  json
// @Param id path int true "Match ID"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /match/{id} [get]
func (h matchesHandlers) GetMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	player, err := h.matchesUC.GetMatch(id)

	if err != nil {
		klog.Error(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

// Get Matches by TournamentId
// @Summary Get matches by tournamentId
// @Description get matches by tournamentId
// @Tags Matches
// @Accept  json
// @Param id path int true "Match ID"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /tournament/match/{id} [get]
func (h matchesHandlers) GetMatchesByTournamentId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	matches, err := h.matchesUC.GetMatchesByTournamentId(id)

	if err != nil {
		klog.Errorf("Error getting matches by tournamentId: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

// Set Match Result
// @Summary Set Match Result
// @Description set match result and set status
// @Tags Matches
// @Accept  json
// @Param tournament body models.Result true "Result object for API"
// @Param id path int true "Match ID"
// @Produce  json
// @Success 201 {object} models.Match
// @Router /match/{id}/result [post]
func (h *matchesHandlers) SetResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		klog.Fatalln(err)
	}

	if err = h.matchesUC.SetResult(id, body); err != nil {
		klog.Errorf("Error setting results match: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Result Setted")
}
