package usecase

import (
	"encoding/json"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/tournaments"
	"fake.com/padel-api/pkg/utils"
	"k8s.io/klog/v2"
)

type matchesUC struct {
	matchesRepo     matches.Repository
	tournamentsRepo tournaments.Repository
}

func NewMatchUseCase(matchesRepo matches.Repository, tournamentsRepo tournaments.Repository) matches.UseCase {
	return &matchesUC{matchesRepo: matchesRepo, tournamentsRepo: tournamentsRepo}
}

func (u *matchesUC) Create(body []byte) (*models.Match, error) {
	klog.Info("ENTERING # UseCaseMatch - Create")
	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	match := models.NewMatch(matchJSON.CoupleOne[0], matchJSON.CoupleOne[1], matchJSON.CoupleTwo[0], matchJSON.CoupleTwo[1], matchJSON.Status, matchJSON.TournamentID, result)

	klog.Info("ENDING # UseCaseMatch - Create")
	return u.matchesRepo.Create(match)
}

func (u *matchesUC) Update(body []byte, matchID int) (*models.Match, error) {
	klog.Info("ENTERING # UseCaseMatch - Update")
	var matchJSON models.MatchJSON
	json.Unmarshal(body, &matchJSON)

	var result models.Result

	if len(matchJSON.Result[0]) > len(matchJSON.Result[1]) {
		result = models.Result{len(matchJSON.Result[0]), matchJSON.Result[0], matchJSON.Result[1]}
	} else {
		result = models.Result{len(matchJSON.Result[1]), matchJSON.Result[0], matchJSON.Result[1]}
	}

	oldMatch, _ := u.matchesRepo.GetMatch(matchID)
	updatedMatch := models.NewMatch(oldMatch.GetAttrs().CoupleOne[0], oldMatch.GetAttrs().CoupleOne[1], oldMatch.GetAttrs().CoupleTwo[0], oldMatch.GetAttrs().CoupleTwo[1], matchJSON.Status, oldMatch.GetAttrs().TournamentID, result)

	tournament, _ := u.tournamentsRepo.GetTournament(oldMatch.GetAttrs().TournamentID)
	tournAttrs := tournament.GetAttrs()

	if getCoupleWinner(&updatedMatch.GetAttrs().Result) {
		for i, p := range tournAttrs.Players {
			if utils.Contains(updatedMatch.GetAttrs().CoupleOne, p.PlayerID) {
				tournAttrs.Players[i].PlayerScore = p.PlayerScore + 1
				tournAttrs.Players[i].RoundsPlayed = p.RoundsPlayed + 1
			} else if utils.Contains(updatedMatch.GetAttrs().CoupleTwo, p.PlayerID) {
				tournAttrs.Players[i].RoundsPlayed = p.RoundsPlayed + 1
			}
		}
	} else {
		for i, p := range tournAttrs.Players {
			if utils.Contains(updatedMatch.GetAttrs().CoupleTwo, p.PlayerID) {
				tournAttrs.Players[i].PlayerScore = p.PlayerScore + 1
				tournAttrs.Players[i].RoundsPlayed = p.RoundsPlayed + 1
			} else if utils.Contains(updatedMatch.GetAttrs().CoupleOne, p.PlayerID) {
				tournAttrs.Players[i].RoundsPlayed = p.RoundsPlayed + 1
			}
		}
	}

	tournament.SetAttrs(tournAttrs)
	u.tournamentsRepo.Update(tournament, oldMatch.GetAttrs().TournamentID)

	klog.Info("ENDING # UseCaseMatch - Update")
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

func (u *matchesUC) GetMatchesByTournamentId(tournamentId int) ([]models.Match, error) {
	return u.matchesRepo.GetMatchesByTournamentId(tournamentId)
}

func (u *matchesUC) SetResult(matchId int, body []byte) error {
	klog.Info("ENTERING # UseCaseMatch - SetResult")
	var result models.Result

	json.Unmarshal(body, &result)

	m, err := u.matchesRepo.GetMatch(matchId)

	if err != nil {
		klog.Error(err)
	}

	matchAttrs := m.GetAttrs()
	matchAttrs.Result.CoupleOneSets = result.CoupleOneSets
	matchAttrs.Result.CoupleTwoSets = result.CoupleTwoSets
	matchAttrs.Result.SetsCounter = result.SetsCounter
	matchAttrs.Status = "Played"
	m.SetAttrs(matchAttrs)

	m, err = u.matchesRepo.Update(m, m.MatchId)

	if err != nil {
		klog.Error(err)
	}

	t, err := u.tournamentsRepo.GetTournament(matchAttrs.TournamentID)

	if err != nil {
		klog.Error(err)
	}

	tournamentAttrs := t.GetAttrs()

	for i, p := range tournamentAttrs.Players {
		if p.PlayerID == matchAttrs.CoupleOne[0] || p.PlayerID == matchAttrs.CoupleOne[1] {
			tournamentAttrs.Players[i].PlayerScore++
			tournamentAttrs.Players[i].RoundsPlayed++
		} else if p.PlayerID == matchAttrs.CoupleTwo[0] || p.PlayerID == matchAttrs.CoupleTwo[1] {
			tournamentAttrs.Players[i].RoundsPlayed++
		}
	}

	t.SetAttrs(tournamentAttrs)

	t, err = u.tournamentsRepo.Update(t, t.TournamentID)

	if err != nil {
		klog.Error(err)
	}

	klog.Info("ENDING # UseCaseMatch - SetResult")

	return nil
}

// Resturn true when CoupleOne wins or false when CoupleTwo wins
func getCoupleWinner(result *models.Result) bool {
	klog.Info("ENTERING # UseCaseMatch - getCoupleWinner")
	winner := [2]int{0, 0}
	for i, s := range result.CoupleOneSets {
		if result.CoupleTwoSets[i] < s {
			winner[0]++
		} else {
			winner[1]++
		}
	}

	klog.Info("ENDING # UseCaseMatch - getCoupleWinner")
	return winner[0] > winner[1]
}
