package mock

import (
	"drone-task/repository"
	"drone-task/repository/entity"
)

type MockedMedicationRepository struct{}

func NewMockedMedicationRepository() repository.IMedicationRepository {
	return MockedMedicationRepository{}
}

func (m MockedMedicationRepository) Create(medications []entity.Medication) ([]entity.Medication, error) {
	return medications, nil
}
