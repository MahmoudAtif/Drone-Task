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
	GetBySerialNumber(serialNumber string) (entity.Drone, error)
	UpdateByID(id int, fields map[string]interface{}) (entity.Drone, error)
	Filter(filters entity.DroneFilters) ([]entity.Drone, error)
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

func (dr DroneRepository) GetBySerialNumber(serialNumber string) (entity.Drone, error) {
	drone := entity.Drone{}
	err := dr.DB.Where("serial_number = ?", serialNumber).Last(&drone).Error
	return drone, err
}

func (dr DroneRepository) UpdateByID(id int, fields map[string]interface{}) (entity.Drone, error) {
	err := dr.DB.Model(&entity.Drone{}).Where("id = ?", id).Updates(fields).Error
	if err != nil {
		return entity.Drone{}, err
	}
	return dr.GetById(id)
}

func (dr DroneRepository) Filter(filters entity.DroneFilters) ([]entity.Drone, error) {
	drones := []entity.Drone{}
	filterConditions := dr.DB
	if len(filters.States) > 0 {
		filterConditions = filterConditions.Where("state IN ?", filters.States)
	}
	if len(filters.SerialNumbers) > 0 {
		filterConditions = filterConditions.Where("serial_number IN ?", filters.SerialNumbers)
	}
	err := filterConditions.Find(&drones).Error
	if err != nil {
		return nil, err
	}
	return drones, nil
}
