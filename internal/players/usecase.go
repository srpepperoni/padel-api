package players

import (
	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	Create(player *models.Player) (*models.Player, error)
	GetPlayers() (*[]models.Player, error)
	GetPlayer(playerID int) (*models.Player, error)
}
