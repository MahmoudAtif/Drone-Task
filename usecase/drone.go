package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/entity"
	useCaseEntity "drone-task/usecase/entity"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/exp/slices"
)

type IDroneUseCase interface {
	Get(ctx context.Context) ([]byte, error)
	Create(ctx context.Context, request []byte) ([]byte, error)
	GetById(ctx context.Context, id int) ([]byte, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, request []byte) ([]byte, error)
	UpdateDroneBateryTask()
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
	drones, err := dr.droneRepository.Get()
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		return nil, err
	}
	return json.Marshal(drones)
}

func (dr DroneUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	drones := []entity.Drone{}
	err := json.Unmarshal(request, &drones)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	createdDrones, errors := dr.ValidateDrones(drones)
	createdDrones, err = dr.droneRepository.Create(createdDrones)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response := useCaseEntity.CreatedDrones{
		CreatedDrones: createdDrones,
		Errors:        errors,
	}
	return json.Marshal(response)
}

func (dr DroneUseCase) GetById(ctx context.Context, id int) ([]byte, error) {
	drone, err := dr.droneRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(drone)
}

func (dr DroneUseCase) Delete(ctx context.Context, id int) error {
	err := dr.droneRepository.Delete(id)
	return err
}

func (dr DroneUseCase) Update(ctx context.Context, request []byte) ([]byte, error) {
	drone := entity.Drone{}
	err := json.Unmarshal(request, &drone)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		return []byte{}, nil
	}
	_, errors := dr.ValidateDrones([]entity.Drone{drone})
	if len(errors) > 0 {
		response := useCaseEntity.Errors{
			Errors: errors,
		}
		return json.Marshal(response)
	}
	updatedDrone, err := dr.droneRepository.Update(drone)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(updatedDrone)
}

func (dr DroneUseCase) ValidateDrones(drones []entity.Drone) ([]entity.Drone, []string) {
	errors := []string{}
	validatedDrones := []entity.Drone{}
	models := []string{
		"Lightweight",
		"Middleweight",
		"Cruiserweight",
		"Heavyweight",
	}
	states := []string{
		"IDLE",
		"LOADING",
		"LOADED",
		"DELIVERING",
		"DELIVERED",
		"RETURNING",
	}

	for _, drone := range drones {
		isValid := true
		if len(drone.SerialNumber) > 100 {
			error := fmt.Sprintf(`drone %v: serial number must be 100 characters max`, drone.SerialNumber)
			errors = append(errors, error)
			isValid = false
		}
		if !slices.Contains(models, drone.Model) {
			error := fmt.Sprintf(`drone %v: invalid drone models`, drone.SerialNumber)
			errors = append(errors, error)
			isValid = false
		}
		if drone.Weight > 500 {
			error := fmt.Sprintf(`drone %v: weight must be less than or equal 500gr`, drone.SerialNumber)
			errors = append(errors, error)
			isValid = false
		}
		if drone.BatteryCapacity < 0 || drone.BatteryCapacity > 100 {
			error := fmt.Sprintf(`drone %v: invalid batery_cabacity precentage`, drone.SerialNumber)
			errors = append(errors, error)
			isValid = false
		}
		if !slices.Contains(states, drone.State) {
			error := fmt.Sprintf(`drone %v: invalid drone state`, drone.SerialNumber)
			errors = append(errors, error)
			isValid = false
		}
		if isValid {
			validatedDrones = append(validatedDrones, drone)
		}
	}
	return validatedDrones, errors
}

func (dr DroneUseCase) UpdateDroneBateryTask() {
	for {
		drones, err := dr.droneRepository.Filter(entity.DroneFilters{States: []string{"LOADED"}})
		if err != nil {
			log.Printf("[Error]: %v", err.Error())
		}
		for _, drone := range drones {
			if drone.BatteryCapacity > 0 {
				_, err = dr.droneRepository.UpdateByID(
					int(drone.ID),
					map[string]interface{}{"battery_capacity": drone.BatteryCapacity - 1},
				)
				if err != nil {
					log.Printf("[Error]: %v", err.Error())
				}
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
