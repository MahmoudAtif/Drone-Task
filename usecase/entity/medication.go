package entity

import "drone-task/repository/entity"


type CreatedMedications struct {
	CreatedMedications []entity.Medication `json:"created_medications"`
	Errors        []string            `json:"errors,omitempty"`
}
