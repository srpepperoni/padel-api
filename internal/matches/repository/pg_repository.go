package repository

import (
	"fmt"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"gorm.io/gorm"
)

type matchesRepo struct {
	db *gorm.DB
}

func NewMatchesRepository(db *gorm.DB) matches.Repository {
	return &matchesRepo{db: db}
}

func (r *matchesRepo) Create(match *models.Match) (*models.Match, error) {
	if result := r.db.Create(&match); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return match, nil
}

func (r *matchesRepo) GetMatches() (*[]models.Match, error) {
	var matches []models.Match

	if result := r.db.Find(&matches); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &matches, nil
}

func (r *matchesRepo) GetMatch(matchID int) (*models.Match, error) {
	var match models.Match

	if result := r.db.First(&match, matchID); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &match, nil
}

func (r *matchesRepo) Update(updatedMatch *models.Match, matchID int) (*models.Match, error) {
	var match models.Match

	if result := r.db.First(&match, matchID); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	match.PlayerOne = updatedMatch.PlayerOne
	match.PlayerTwo = updatedMatch.PlayerTwo
	match.PlayerThree = updatedMatch.PlayerThree
	match.PlayerFour = updatedMatch.PlayerFour
	match.Status = updatedMatch.Status
	match.Result = updatedMatch.Status

	r.db.Save(&match)

	return &match, nil
}
