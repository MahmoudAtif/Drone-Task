package entity

import (
	"drone-task/repository/entity"
)

type Drone struct {
	SerialNumber    string  `json:"serial_number"`
	Model           string  `json:"model"`
	Weight          float64 `json:"weight"`
	BatteryCapacity int     `json:"batery_capacity"`
	State           string  `json:"state"`
}

type CreatedDrones struct {
	CreatedDrones []entity.Drone `json:"created_drones"`
	Errors        []string       `json:"errors,omitempty"`
}

type Errors struct {
	Errors []string `json:"errors,omitempty"`
}
