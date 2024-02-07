package entity

type Drone struct {
	AbstractModel
	SerialNumber    string  `gorm:"type:varchar(100);unique;not null" json:"serial_number"`
	Model           string  `json:"model"`
	Weight          float64 `json:"weight"`
	BatteryCapacity float64 `json:"batery_capacity"`
	State           string  `json:"state"`
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
