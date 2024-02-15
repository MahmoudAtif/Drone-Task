package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/entity"
	"encoding/json"
	"log"
)

type IDroneLoadUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
}

type DroneLoadUseCase struct {
	droneLoadRepository  repository.IDroneLoadRepository
	droneRepository      repository.IDroneRepository
	medicationRepository repository.IMedicationRepository
}

func NewDroneLoadUseCase(
	droneLoadRepository repository.IDroneLoadRepository,
	droneRepository repository.IDroneRepository,
	medicationRepository repository.IMedicationRepository,
) DroneLoadUseCase {
	return DroneLoadUseCase{
		droneLoadRepository:  droneLoadRepository,
		droneRepository:      droneRepository,
		medicationRepository: medicationRepository,
	}
}

func (dlu DroneLoadUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	droneLoad := entity.DroneLoad{}
	errors := []string{}
	err := json.Unmarshal(request, &droneLoad)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	drones, err := dlu.droneRepository.Filter(entity.DroneFilters{SerialNumbers: []string{droneLoad.DroneSerialNumber}})
	if len(drones) == 0 || err != nil {
		errors = append(errors, "Invalid drone serial number")
	}

	medications, err := dlu.medicationRepository.Filter(entity.MedicationFilters{Codes: []string{droneLoad.MedicationCode}})
	if len(medications) == 0 || err != nil {
		errors = append(errors, "Invalid mecication code")
	}
	if len(errors) > 0 {
		return json.Marshal(errors)
	}
	if drones[0].State != "IDLE" || drones[0].BatteryCapacity < 10 || medications[0].Weight > drones[0].Weight {
		errors = append(errors, "This drone does not meet specifications")
		return json.Marshal(errors)
	}
	droneLoad, err = dlu.droneLoadRepository.Create(droneLoad)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = dlu.droneRepository.UpdateByID(int(drones[0].ID), map[string]interface{}{"state": "LOADED"})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return json.Marshal(droneLoad)
}
