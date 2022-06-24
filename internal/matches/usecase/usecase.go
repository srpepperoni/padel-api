package usecase

import (
	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
)

type matchesUC struct {
	matchesRepo matches.Repository
}

func NewMatchUseCase(matchesRepo matches.Repository) matches.UseCase {
	return &matchesUC{matchesRepo: matchesRepo}
}

func (u *matchesUC) Update(match *models.Match, matchID int) (*models.Match, error) {
	return u.matchesRepo.Update(match, matchID)
}

func (u *matchesUC) Create(match *models.Match) (*models.Match, error) {
	n, err := u.matchesRepo.Create(match)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (u *matchesUC) Delete(matchId int) error {
	return u.matchesRepo.Delete(matchId)
}

func (u *matchesUC) GetMatches() (*[]models.Match, error) {
	return u.matchesRepo.GetMatches()
}

func (u *matchesUC) GetMatch(playerID int) (*models.Match, error) {
	return u.matchesRepo.GetMatch(playerID)
}

func (u *matchesUC) GetMatchesByTournamentId(tournamentId int) (*[]models.Match, error) {
	return u.matchesRepo.GetMatchesByTournamentId(tournamentId)
}
