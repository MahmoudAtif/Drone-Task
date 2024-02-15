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

func (m MockedMedicationRepository) GetByCode(code string) (entity.Medication, error) {
	medication := entity.Medication{}
	return medication, nil
}

func (m MockedMedicationRepository) Filter(filters entity.MedicationFilters) ([]entity.Medication, error) {
	medications := []entity.Medication{
		{
			Name:   "kataflam",
			Weight: 10,
			Code:   "10",
			Image:  "test.png",
		},
		{
			Name:   "kanakomp",
			Weight: 100,
			Code:   "20",
			Image:  "test.png",
		},
		{
			Name:   "zyrtec",
			Weight: 50,
			Code:   "30",
			Image:  "test.png",
		},
	}
	filterdMedications := []entity.Medication{}

	for _, medication := range medications {
		if len(filters.Codes) > 0 {
			found := false
			for _, code := range filters.Codes {
				if medication.Code == code {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		filterdMedications = append(filterdMedications, medication)
	}
	return filterdMedications, nil
}
