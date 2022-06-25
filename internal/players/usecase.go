package players

import (
	"fake.com/padel-api/internal/models"
)

type UseCase interface {
	Create(body []byte) (*models.Player, error)
	Update(body []byte, playerId int) (*models.Player, error)
	Delete(playerId int) error
	GetPlayers() (*[]models.Player, error)
	GetPlayer(playerID int) (*models.Player, error)
}
