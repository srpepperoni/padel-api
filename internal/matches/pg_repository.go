package matches

import (
	"fake.com/padel-api/internal/models"
)

type Repository interface {
	Create(match *models.Match) (*models.Match, error)
	Update(match *models.Match, matchID int) (*models.Match, error)
	Delete(matchID int) error
	GetMatches() ([]models.Match, error)
	GetMatch(matchID int) (*models.Match, error)
	GetMatchesByTournamentId(tournamentId int) ([]models.Match, error)
	GetMatchesByTournamentIdAndStatus(tournmentId int, status string) ([]models.Match, error)
}
