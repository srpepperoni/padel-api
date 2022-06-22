package usecase

import (
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/tournaments"
)

type tournamentsUC struct {
	tournamentsRepo tournaments.Repository
}

func NewTournamentsUseCase(tournamentsRepo tournaments.Repository) tournaments.UseCase {
	return &tournamentsUC{tournamentsRepo: tournamentsRepo}
}

// Create news
func (u *tournamentsUC) Create(tournament *models.Tournament) (*models.Tournament, error) {
	n, err := u.tournamentsRepo.Create(tournament)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (u *tournamentsUC) Update(tournament *models.Tournament, tournamentID int) (*models.Tournament, error) {
	return u.tournamentsRepo.Update(tournament, tournamentID)
}

func (u *tournamentsUC) Delete(tournamentID int) error {
	return u.tournamentsRepo.Delete(tournamentID)
}

func (u *tournamentsUC) GetTournaments() (*[]models.Tournament, error) {
	return u.tournamentsRepo.GetTournaments()
}

func (u *tournamentsUC) GetTournament(playerID int) (*models.Tournament, error) {
	return u.tournamentsRepo.GetTournament(playerID)
}
