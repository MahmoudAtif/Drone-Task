package repository

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"drone-task/repository/db"
	"drone-task/repository/db/migrations"

	"gorm.io/gorm"
)

func InitializeTestDatabase(schemaName string) (*gorm.DB, error) {
	re1 := regexp.MustCompile(`dbname=(\w+)`)
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	databaseName := re1.FindAllStringSubmatch(connectionString, -1)[0][1]
	defaultConnectionString := strings.ReplaceAll(connectionString, databaseName, "postgres")
	dbClient, err := db.ConnectToDataBase(defaultConnectionString)
	if err != nil {
		fmt.Println(fmt.Sprintf("can not connect postgres error: %s", err.Error()))
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	tx := dbClient.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS test_%s;", databaseName))
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		log.Fatal(err)
	}
	tx = dbClient.Exec(fmt.Sprintf("CREATE DATABASE test_%s;", databaseName))
	if tx.Error != nil {
		fmt.Println(fmt.Sprintf("can not create test database, error: %s", tx.Error.Error()))
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	connectionString = strings.Replace(os.Getenv("DB_CONNECTION_STRING"), "dbname=", "dbname=test_", -1)
	dbClient, err = db.ConnectToDataBase(connectionString)
	if err != nil {
		fmt.Println(fmt.Sprintf("can not connect postgres error: %s", err.Error()))
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	tx = dbClient.Exec(fmt.Sprintf("CREATE SCHEMA %s", schemaName))
	if tx.Error != nil {
		fmt.Println(fmt.Sprintf("can not create schema, error: %s", tx.Error.Error()))
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	err = migrations.MigrateDrone(dbClient)
	if err != nil {
		log.Println(err)
	}
	err = migrations.MigrateMedication(dbClient)
	if err != nil {
		log.Println(err)
	}
	err = migrations.MigrateDroneLoad(dbClient)
	if err != nil {
		log.Println(err)
	}

	return dbClient, nil
}

func DestroyTestDataBase(dbClient *gorm.DB) {
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	re := regexp.MustCompile(`dbname=(\w+)`)
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	databaseName := re.FindAllStringSubmatch(connectionString, -1)[0][1]
	defaultConnectionString := strings.ReplaceAll(connectionString, databaseName, "postgres")
	dbClient, err = db.ConnectToDataBase(defaultConnectionString)
	if err != nil {
		log.Println(err)
	}
	dbClient.Exec("DROP DATABASE IF EXISTS test_drone_task;")
}
