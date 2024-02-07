package entity

type DroneLoad struct {
	AbstractModel
	DroneSerialNumber string `gorm:"not null" json:"drone_serial_number"`
	MedicationCode   string `gorm:"not null" json:"medication_code"`
}
