package repository

import (
	"drone-task/repository/entity"

	"gorm.io/gorm"
)

type IMedicationRepository interface {
	Create(medications []entity.Medication) ([]entity.Medication, error)
	GetByCode(code string) (entity.Medication, error)
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

func (dr MedicationRepository) GetByCode(code string) (entity.Medication, error) {
	medication := entity.Medication{}
	err := dr.DB.Where("code = ?", code).Last(&medication).Error
	return medication, err
}
