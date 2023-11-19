package entity

type Drone struct {
	AbstractModel
	SerialNumber    string `gorm:"type:varchar(100);not null"`
	Model           ModelTypes
	WeightLimit     float64
	BatteryCapacity int
	State           DroneState
}

type DroneState string
type ModelTypes string

const (
	IDLE       DroneState = "IDLE"
	LOADING    DroneState = "LOADING"
	LOADED     DroneState = "LOADED"
	DELIVERING DroneState = "DELIVERING"
	DELIVERED  DroneState = "DELIVERED"
	RETURNING  DroneState = "RETURNING"
)

const (
	Lightweight   ModelTypes = "Lightweight"
	Middleweight  ModelTypes = "Middleweight"
	Cruiserweight ModelTypes = "Cruiserweight"
	Heavyweight   ModelTypes = "Heavyweight"
)
