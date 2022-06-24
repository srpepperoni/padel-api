package postgres

import (
	"log"

	"fake.com/padel-api/internal/models"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	//TODO: make the user and password configurable
	dbURL := "postgres://postgres:example@localhost:5432/sysdig_padel"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Player{}, &models.Match{}, &models.Tournament{})

	return db, nil
}
