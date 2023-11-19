package repository

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

type IDroneRepository interface {
	Get() ([]entity.Drone, error)
	Create(Drones []entity.Drone) ([]entity.Drone, error)
}

type DroneRepository struct {
	DB *gorm.DB
}

func NewDroneRepository(DB *gorm.DB) IDroneRepository {
	return &DroneRepository{DB: DB}
}

func (dr DroneRepository) Get() ([]entity.Drone, error) {
	drones := []entity.Drone{}
	err := dr.DB.Find(&drones).Error
	return drones, err
}

func (dr DroneRepository) Create(drones []entity.Drone) ([]entity.Drone, error) {
	err := dr.DB.Create(&drones).Error
	return drones, err
}
