package repository

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

type IMedicationRepository interface {
	Create(medications []entity.Medication) ([]entity.Medication, error)
}

type MedicationRepository struct {
	DB *gorm.DB
}

func NewMedicationRepository(DB *gorm.DB) IMedicationRepository {
	return &MedicationRepository{DB: DB}
}

func (mr MedicationRepository) Create(medications []entity.Medication) ([]entity.Medication, error) {
	err := mr.DB.Create(&medications).Error
	return medications, err
}
