package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/entity"
	useCaseEntity "drone-task/usecase/entity"
	"encoding/json"
	"log"
)

type IDroneUseCase interface {
	Get(ctx context.Context) ([]byte, error)
	Create(ctx context.Context, request []byte) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository repository.IDroneRepository
}

func NewDroneUseCase(droneRepository repository.IDroneRepository) IDroneUseCase {
	return DroneUseCase{
		droneRepository: droneRepository,
	}
}

func (dr DroneUseCase) Get(ctx context.Context) ([]byte, error) {
	log.Printf("usecaseeeeeeeeeeeeeeeeeeeeeeeee")
	drones, err := dr.droneRepository.Get()
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		return nil, err
	}
	return json.Marshal(drones)
}

func (dr DroneUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	drones := []entity.Drone{}
	message := ""
	err := json.Unmarshal(request, &drones)
	for _, drone := range drones {
		if drone.WeightLimit > 500 {
			message = "weight must be less than or equal 500"
			break
		}
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if message == "" {
		drones, err = dr.droneRepository.Create(drones)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	response := useCaseEntity.CreatedDrones{
		CreatedDrones: drones,
		Message:       message,
	}
	return json.Marshal(response)
}
