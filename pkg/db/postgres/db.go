package postgres

import (
	"log"
	"os"

	"fake.com/padel-api/config"
	"fake.com/padel-api/internal/models"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://" + cfg.Postgres.PostgresqlUser + ":" + cfg.Postgres.PostgresqlPassword + "@" +
			cfg.Postgres.PostgresqlHost + ":" + cfg.Postgres.PostgresqlPort + "/" + cfg.Postgres.PostgresqlDbname
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Player{}, &models.Match{}, &models.Tournament{})

	if err != nil {
		panic(err)
	}

	return db, nil
}
