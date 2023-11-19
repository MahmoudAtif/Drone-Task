package entity

import (
	repoEntity "drone-task/repository/entity"
)

type Drone struct {
	SerialNumber    string  `json:"serial_number"`
	Model           string  `json:"model"`
	WeightLimit     float64 `json:"weight_limit"`
	BatteryCapacity int     `json:"batery_capacity"`
	State           string  `json:"state"`
}

type CreatedDrones struct {
	CreatedDrones []repoEntity.Drone `json:"created_drones"`
	Message       string             `json:"message,omitempty"`
}
