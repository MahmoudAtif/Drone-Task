package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/entity"
	useCaseEntity "drone-task/usecase/entity"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"golang.org/x/exp/slices"
)

type IMedicationUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
}

type MedicationUseCase struct {
	medicationRepository repository.IMedicationRepository
}

func NewMedicationUseCase(medicationRepository repository.IMedicationRepository) IMedicationUseCase {
	return MedicationUseCase{
		medicationRepository: medicationRepository,
	}
}

func (mu MedicationUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	medications := []entity.Medication{}
	createdMedications := []entity.Medication{}
	err := json.Unmarshal(request, &medications)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	validatedMedications, errors := mu.ValidateMedications(medications)
	if len(validatedMedications) > 0 {
		createdMedications, err = mu.medicationRepository.Create(validatedMedications)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	response := useCaseEntity.CreatedMedications{
		CreatedMedications: createdMedications,
		Errors:             errors,
	}
	return json.Marshal(response)
}

func (mu MedicationUseCase) ValidateMedications(medications []entity.Medication) ([]entity.Medication, []string) {
	errors := []string{}
	validatedMedications := []entity.Medication{}
	validExtensions := []string{
		".jpg",
		".jpeg",
		".png",
	}

	for _, medication := range medications {
		isValid := true
		nameRegex := regexp.MustCompile("^[a-zA-Z0-9-_]+$")
		if !nameRegex.MatchString(medication.Name) {
			error := fmt.Sprintf(`medication %v: Invalid name. allowed only: letters, numbers, '-', '_'.`, medication.Name)
			errors = append(errors, error)
			isValid = false
		}
		codeRegex := regexp.MustCompile("^[A-Z0-9_]+$")
		if !codeRegex.MatchString(medication.Code) {
			error := fmt.Sprintf(`medication %v: Invalid code. allowed only: upper case letters, numbers, '_'.`, medication.Name)
			errors = append(errors, error)
			isValid = false
		}
		imageExt := filepath.Ext(medication.Image)
		if medication.Image != "" && !slices.Contains(validExtensions, imageExt) {
			error := fmt.Sprintf(`medication %v: Invalid image. allowed extentions: (.jpg, .jpeg, .png)`, medication.Name)
			errors = append(errors, error)
			isValid = false
		}
		if isValid {
			validatedMedications = append(validatedMedications, medication)
		}
	}
	return validatedMedications, errors

}
