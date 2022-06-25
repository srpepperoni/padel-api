package usecase

import (
	"encoding/json"

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
func (u *tournamentsUC) Create(body []byte) (*models.Tournament, error) {
	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds)

	return u.tournamentsRepo.Create(tournament)
}

func (u *tournamentsUC) Update(body []byte, tournamentID int) (*models.Tournament, error) {
	var tournamentJSON models.TournamentJSON
	json.Unmarshal(body, &tournamentJSON)

	tournament := models.NewTournament(tournamentJSON.Icon, tournamentJSON.Name, tournamentJSON.Description, tournamentJSON.Rounds, tournamentJSON.ActualRounds)

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
