package matches

import (
	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	Create(match *models.Match) (*models.Match, error)
	Update(match *models.Match, matchID int) (*models.Match, error)
	Delete(matchID int) error
	GetMatches() (*[]models.Match, error)
	GetMatch(matchID int) (*models.Match, error)
	GetMatchesByTournamentId(tournamentId int) (*[]models.Match, error)
}
