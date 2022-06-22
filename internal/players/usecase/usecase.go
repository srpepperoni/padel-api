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

func (u *playersUC) Create(player *models.Player) (*models.Player, error) {
	n, err := u.playersRepo.Create(player)
	if err != nil {
		return nil, err
	}

	return n, err
}

func (u *playersUC) Update(player *models.Player, playerID int) (*models.Player, error) {
	return u.playersRepo.Update(player, playerID)
}

func (u *playersUC) Delete(playerID int) error {
	return u.playersRepo.Delete(playerID)
}

func (u *playersUC) GetPlayers() (*[]models.Player, error) {
	return u.playersRepo.GetPlayers()
}

func (u *playersUC) GetPlayer(playerID int) (*models.Player, error) {
	return u.playersRepo.GetPlayer(playerID)
}
