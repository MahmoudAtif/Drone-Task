package entity

type Medication struct {
	AbstractModel
	Name   string  `gorm:"not null" json:"name"`
	Weight float64 `gorm:"not null" json:"weight"`
	Code   string  `gorm:"not null; unique" json:"code"`
	Image  string  `gorm:"" json:"image"`
}
