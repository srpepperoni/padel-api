package usecase

import (
	"encoding/json"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
)

type matchesUC struct {
	matchesRepo matches.Repository
}

func NewMatchUseCase(matchesRepo matches.Repository) matches.UseCase {
	return &matchesUC{matchesRepo: matchesRepo}
}

func (u *matchesUC) Create(body []byte) (*models.Match, error) {
	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	match := models.NewMatch(matchJSON.CoupleOne[0], matchJSON.CoupleOne[1], matchJSON.CoupleTwo[0], matchJSON.CoupleTwo[1], matchJSON.Status, matchJSON.TournamentID, result)

	return u.matchesRepo.Create(match)
}

func (u *matchesUC) Update(body []byte, matchID int) (*models.Match, error) {
	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	updatedMatch := models.NewMatch(matchJSON.CoupleOne[0], matchJSON.CoupleOne[1], matchJSON.CoupleTwo[0], matchJSON.CoupleTwo[1], matchJSON.Status, matchJSON.TournamentID, result)

	return u.matchesRepo.Update(updatedMatch, matchID)
}

func (u *matchesUC) Delete(matchId int) error {
	return u.matchesRepo.Delete(matchId)
}

func (u *matchesUC) GetMatches() ([]models.Match, error) {
	return u.matchesRepo.GetMatches()
}

func (u *matchesUC) GetMatch(playerID int) (*models.Match, error) {
	return u.matchesRepo.GetMatch(playerID)
}

func (u *matchesUC) GetMatchesByTournamentId(tournamentId int) (*[]models.Match, error) {
	return u.matchesRepo.GetMatchesByTournamentId(tournamentId)
}
