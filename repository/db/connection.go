package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func GetNewInstance(db *gorm.DB) *gorm.DB {
	tx := &gorm.DB{Config: db.Config, Error: db.Error}
	tx.Statement = &gorm.Statement{
		DB:       tx,
		ConnPool: db.Statement.ConnPool,
		Context:  db.Statement.Context,
		Clauses:  map[string]clause.Clause{},
		Vars:     make([]interface{}, 0, 8),
	}
	return tx
}
