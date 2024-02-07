package mock

import (
	"drone-task/repository"
	"drone-task/repository/entity"
)

type MockedDroneLoadRepository struct{}

func NewMockedDroneLoadRepository() repository.IDroneLoadRepository {
	return MockedDroneLoadRepository{}
}

func (m MockedDroneLoadRepository) Create(dronLoad entity.DroneLoad) (entity.DroneLoad, error) {
	return dronLoad, nil
}
