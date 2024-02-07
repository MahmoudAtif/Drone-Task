package repository

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

type IDroneLoadRepository interface {
	Create(dronLoad entity.DroneLoad) (entity.DroneLoad, error)
}

type DroneLoadRepository struct {
	DB *gorm.DB
}

func NewDroneLoadRepository(DB *gorm.DB) IDroneLoadRepository {
	return DroneLoadRepository{DB: DB}
}

func (dlr DroneLoadRepository) Create(dronLoad entity.DroneLoad) (entity.DroneLoad, error) {
	err := dlr.DB.Create(&dronLoad).Error
	return dronLoad, err
}
