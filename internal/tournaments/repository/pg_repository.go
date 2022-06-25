package repository

import (
	"k8s.io/klog/v2"

	"fake.com/padel-api/internal/models"
	"fake.com/padel-api/internal/tournaments"
	"gorm.io/gorm"
)

type tournamentsRepo struct {
	db *gorm.DB
}

func NewTournamentsRepository(db *gorm.DB) tournaments.Repository {
	return &tournamentsRepo{db: db}
}

func (r *tournamentsRepo) Create(tournament *models.Tournament) (*models.Tournament, error) {
	if result := r.db.Create(&tournament); result.Error != nil {
		klog.Errorf("Error creating tournament: %v", result.Error)
		return nil, result.Error
	}

	return tournament, nil
}

func (r *tournamentsRepo) Update(updatedTournament *models.Tournament, tournamentID int) (*models.Tournament, error) {
	var tournament models.Tournament

	if result := r.db.Find(&tournament, tournamentID); result.Error != nil {
		klog.Errorf("Error finding tournament: %v", result.Error)
		return nil, result.Error
	}

	tournament.Attrs = updatedTournament.Attrs

	r.db.Save(&tournament)

	return &tournament, nil
}

func (r *tournamentsRepo) Delete(tournamentId int) error {
	var tournament models.Tournament

	if result := r.db.Find(&tournament, tournamentId); result.Error != nil {
		klog.Errorf("Error finding tournament: %v", result.Error)
	}

	r.db.Delete(&tournament)

	return nil
}

func (r *tournamentsRepo) GetTournaments() ([]models.Tournament, error) {
	var tournaments []models.Tournament

	if result := r.db.Find(&tournaments); result.Error != nil {
		klog.Errorf("Error finding tournaments: %v", result.Error)
		return nil, result.Error
	}

	return tournaments, nil
}

func (r *tournamentsRepo) GetTournament(tournamentID int) (*models.Tournament, error) {
	var tournament models.Tournament

	if result := r.db.Find(&tournament, tournamentID); result.Error != nil {
		klog.Errorf("Error finding tournament: %v", result.Error)
		return nil, result.Error
	}

	return &tournament, nil
}
