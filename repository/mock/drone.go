package mock

import (
	"drone-task/repository"
	"drone-task/repository/entity"
)

type MockedDroneRepository struct{}

func NewMockedDroneRepository() repository.IDroneRepository {
	return MockedDroneRepository{}
}

func (m MockedDroneRepository) Get() ([]entity.Drone, error) {
	drones := []entity.Drone{
		{SerialNumber: "1", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
	}
	return drones, nil

}

func (m MockedDroneRepository) Create(drones []entity.Drone) ([]entity.Drone, error) {
	return drones, nil
}

func (m MockedDroneRepository) GetById(id int) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (m MockedDroneRepository) Delete(id int) error {
	return nil
}

func (m MockedDroneRepository) Update(drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}
