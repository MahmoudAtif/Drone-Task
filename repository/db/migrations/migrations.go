package migrations

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

func MigrateDrone(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Drone{})
	return err
}

func MigrateMedication(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Medication{})
	return err
}

func MigrateDroneLoad(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.DroneLoad{})
	return err
}
