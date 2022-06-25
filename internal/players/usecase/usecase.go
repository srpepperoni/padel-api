package usecase

import (
	"encoding/json"

	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/players"
)

type playersUC struct {
	playersRepo players.Repository
}

func NewPlayersUseCase(playersRepo players.Repository) players.UseCase {
	return &playersUC{playersRepo: playersRepo}
}

func (u *playersUC) Create(body []byte) (*models.Player, error) {
	var playerJSON models.PlayerJSON
	json.Unmarshal(body, &playerJSON)

	player := models.NewPlayer(playerJSON.Name, playerJSON.LastName, playerJSON.PlayerName)

	return u.playersRepo.Create(player)
}

func (u *playersUC) Update(body []byte, playerID int) (*models.Player, error) {
	var updatedPlayer models.PlayerJSON
	json.Unmarshal(body, &updatedPlayer)

	player := models.NewPlayer(updatedPlayer.Name, updatedPlayer.LastName, updatedPlayer.PlayerName)

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
