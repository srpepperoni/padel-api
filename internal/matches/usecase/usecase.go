package usecase

import (
	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
)

const (
	basePrefix = "api-matches:"
)

type matchesUC struct {
	matchesRepo matches.Repository
}

func NewMatchsUseCase(matchesRepo matches.Repository) matches.UseCase {
	return &matchesUC{matchesRepo: matchesRepo}
}

// Create news
func (u *matchesUC) Create(match *models.Match) (*models.Match, error) {
	n, err := u.matchesRepo.Create(match)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (u *matchesUC) GetMatches() (*[]models.Match, error) {
	return u.matchesRepo.GetMatches()
}

func (u *matchesUC) GetMatch(playerID int) (*models.Match, error) {
	return u.matchesRepo.GetMatch(playerID)
}

func (u *matchesUC) Update(match *models.Match, matchID int) (*models.Match, error) {
	return u.matchesRepo.Update(match, matchID)
}
