package tournaments

import (
	"fake.com/padel-api/internal/models"
)

type Repository interface {
	Create(tournament *models.Tournament) (*models.Tournament, error)
	Update(tournament *models.Tournament, tournamentID int) (*models.Tournament, error)
	Delete(tournamentId int) error
	GetTournaments() (*[]models.Tournament, error)
	GetTournament(tournamentID int) (*models.Tournament, error)
}
