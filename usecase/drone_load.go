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
	err := json.Unmarshal(request, &droneLoad)
	errors := []string{}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = dlu.droneRepository.GetBySerialNumber(droneLoad.DroneSerialNumber)
	if err != nil {
		errors = append(errors, "Invalid drone serial number")
	}
	_, err = dlu.medicationRepository.GetByCode(droneLoad.MedicationCode)
	if err != nil {
		
		errors = append(errors, "Invalid mecication code")
	}
	if len(errors) > 0 {
		return json.Marshal(errors)
	}
	droneLoad, err = dlu.droneLoadRepository.Create(droneLoad)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return json.Marshal(droneLoad)
}
