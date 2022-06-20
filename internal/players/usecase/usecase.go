package usecase

import (
	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
)

const (
	basePrefix = "api-players:"
)

type playersUC struct {
	playersRepo players.Repository
}

func NewPlayersUseCase(playersRepo players.Repository) players.UseCase {
	return &playersUC{playersRepo: playersRepo}
}

// Create news
func (u *playersUC) Create(player *models.Player) (*models.Player, error) {
	n, err := u.playersRepo.Create(player)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (u *playersUC) GetPlayers() (*[]models.Player, error) {
	return u.playersRepo.GetPlayers()
}

func (u *playersUC) GetPlayer(playerID int) (*models.Player, error) {
	return u.playersRepo.GetPlayer(playerID)
}
