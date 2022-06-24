package players

import (
	"fake.com/padel-api/internal/models"
)

// Repository News Repository
type Repository interface {
	Create(player *models.Player) (*models.Player, error)
	Update(player *models.Player, playerId int) (*models.Player, error)
	Delete(playerId int) error
	GetPlayers() (*[]models.Player, error)
	GetPlayer(playerID int) (*models.Player, error)
}
