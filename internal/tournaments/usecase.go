package tournaments

import (
	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	Create(body []byte) (*models.Tournament, error)
	Update(body []byte, tournamentID int) (*models.Tournament, error)
	Delete(tournamentId int) error
	GetTournaments() (*[]models.Tournament, error)
	GetTournament(tournamentID int) (*models.Tournament, error)
}
