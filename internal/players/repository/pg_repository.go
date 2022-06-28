package repository

import (
	"k8s.io/klog/v2"

	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
	"gorm.io/gorm"
)

type playersRepo struct {
	db *gorm.DB
}

func NewPlayersRepository(db *gorm.DB) players.Repository {
	return &playersRepo{db: db}
}

func (r *playersRepo) Create(player *models.Player) (*models.Player, error) {
	if result := r.db.Create(&player); result.Error != nil {
		klog.Errorf("Error creating player: %v", result.Error)
	}

	return player, nil
}

func (r *playersRepo) Update(updatedPlayer *models.Player, playerId int) (*models.Player, error) {
	var player models.Player

	if result := r.db.Find(&player, playerId); result.Error != nil {
		klog.Errorf("Error updating player: %v", result.Error)
		return nil, result.Error
	}

	player.Attrs = updatedPlayer.Attrs

	r.db.Save(&player)

	return &player, nil
}

func (r *playersRepo) Delete(playerId int) error {
	var player models.Player

	if result := r.db.Find(&player, playerId); result.Error != nil {
		klog.Errorf("Error deleting player: %v", result.Error)
	}

	r.db.Delete(&player)

	return nil
}

func (r *playersRepo) GetPlayers() ([]models.Player, error) {
	var players []models.Player

	if result := r.db.Find(&players); result.Error != nil {
		klog.Errorf("Error getting players: %v", result.Error)
		return nil, result.Error
	}

	return players, nil
}

func (r *playersRepo) GetPlayer(playerID int) (*models.Player, error) {
	var player models.Player

	if result := r.db.First(&player, playerID); result.Error != nil {
		klog.Errorf("Error getting player: %v", result.Error)
		return nil, result.Error
	}

	return &player, nil
}
