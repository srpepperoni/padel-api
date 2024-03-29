package repository

import (
	"k8s.io/klog/v2"

	"fake.com/padel-api/internal/matches"
	"fake.com/padel-api/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type matchesRepo struct {
	db *gorm.DB
}

func NewMatchesRepository(db *gorm.DB) matches.Repository {
	return &matchesRepo{db: db}
}

func (r *matchesRepo) Create(match *models.Match) (*models.Match, error) {
	klog.Info("CREATE REPO MATCH")
	if result := r.db.Create(&match); result.Error != nil {
		klog.Errorf("Error creating match: %v", result.Error)
		return nil, result.Error
	}

	return match, nil
}

func (r *matchesRepo) Update(updatedMatch *models.Match, matchID int) (*models.Match, error) {
	var match models.Match

	if result := r.db.Find(&match, matchID); result.Error != nil {
		klog.Errorf("Error updating match: %v", result.Error)
		return nil, result.Error
	}

	match.Attrs = updatedMatch.Attrs

	r.db.Save(&match)

	return &match, nil
}

func (r *matchesRepo) Delete(matchId int) error {
	var match models.Match

	if result := r.db.Find(&match, matchId); result.Error != nil {
		klog.Errorf("Error deleting match: %v", result.Error)
	}

	r.db.Delete(&match)

	return nil
}

func (r *matchesRepo) GetMatches() ([]models.Match, error) {
	var matches []models.Match

	if result := r.db.Find(&matches); result.Error != nil {
		klog.Errorf("Error getting matches: %v", result.Error)
		return nil, result.Error
	}

	return matches, nil
}

func (r *matchesRepo) GetMatch(matchID int) (*models.Match, error) {
	var match models.Match

	if result := r.db.Find(&match, matchID); result.Error != nil {
		klog.Errorf("Error getting match: %v", result.Error)
		return nil, result.Error
	}

	return &match, nil
}

func (r *matchesRepo) GetMatchesByTournamentId(tournmentId int) ([]models.Match, error) {
	matches := []models.Match{}
	r.db.Find(&matches, datatypes.JSONQuery("attrs").Equals(tournmentId, "tournamentID"))
	return matches, nil
}

func (r *matchesRepo) GetMatchesByTournamentIdAndStatus(tournmentId int, status string) ([]models.Match, error) {
	matches := []models.Match{}
	r.db.Find(&matches, datatypes.JSONQuery("attrs").Equals(tournmentId, "tournamentID"), datatypes.JSONQuery("attrs").Equals(status, "status"))
	return matches, nil
}
