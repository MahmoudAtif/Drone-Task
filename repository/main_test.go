package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
)

var dbClient *gorm.DB

func setup(schemaName string) {
	fmt.Println("setup test database ...")
	var err error
	dbClient, err = InitializeTestDatabase(schemaName)
	if err != nil {
		log.Println("Couldn't connect to database", err.Error())
		log.Fatal(err)
	}
}

func shutdown() {
	fmt.Println("destroy test database ...")
	DestroyTestDataBase(dbClient)
}

func TestMain(m *testing.M) {
	setup("drone")
	code := m.Run()
	shutdown()
	os.Exit(code)
}
