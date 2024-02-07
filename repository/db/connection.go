package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDataBase(args ...string) (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONNECTION_STRING")
	if len(args) > 0 {
		dbConfig = args[0]
	}
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
