package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-challenge/src/model"
)

func ConnectToPostgres() (*gorm.DB, error) {
	uri := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	db.AutoMigrate(model.Product{})

	if db != nil {
		log.Printf("Connected to the database!")
	}

	return db, nil
}
