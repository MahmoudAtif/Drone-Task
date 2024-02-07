package entity

import "drone-task/repository/entity"

type CreatedDroneLoad struct {
	CreatedDroneLoad entity.DroneLoad `json:"created_drone_load,omitempty"`
	Errors           []string         `json:"errors,omitempty"`
}
