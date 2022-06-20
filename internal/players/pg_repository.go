package players

import (
	"fake.com/padel-api/internal/models"
)

// News Repository
type Repository interface {
	Create(player *models.Player) (*models.Player, error)
	GetPlayers() (*[]models.Player, error)
	GetPlayer(playerID int) (*models.Player, error)
}
