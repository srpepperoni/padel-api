package matches

import (
	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	Create(match *models.Match) (*models.Match, error)
	GetMatches() (*[]models.Match, error)
	GetMatch(matchID int) (*models.Match, error)
	Update(match *models.Match, matchID int) (*models.Match, error)
}
