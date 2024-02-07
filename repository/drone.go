package repository

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

type IDroneRepository interface {
	Get() ([]entity.Drone, error)
	Create(drones []entity.Drone) ([]entity.Drone, error)
	GetById(id int) (entity.Drone, error)
	Delete(id int) error
	Update(drone entity.Drone) (entity.Drone, error)
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

func (dr DroneRepository) GetById(id int) (entity.Drone, error) {
	drone := entity.Drone{}
	err := dr.DB.First(&drone, id).Error
	return drone, err
}

func (dr DroneRepository) Delete(id int) error {
	err := dr.DB.Delete(&entity.Drone{}, id).Error
	return err
}

func (dr DroneRepository) Update(drone entity.Drone) (entity.Drone, error) {
	err := dr.DB.Model(&drone).Updates(&drone).Error
	return drone, err
}
