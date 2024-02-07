package repository

import (
	"drone-task/repository/entity"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDroneLoadRepositoryCreate(t *testing.T) {
	db := dbClient.Begin()
	defer db.Rollback()
	medicationRepository := NewDroneLoadRepository(db)
	tests := []struct {
		name    string
		args    entity.DroneLoad
		want    entity.DroneLoad
		wantErr string
	}{
		{
			name: "Test can create drone load",
			args: entity.DroneLoad{
				DroneSerialNumber: "123XXX",
				MedicationCode:    "123XXX",
			},
			want: entity.DroneLoad{
				DroneSerialNumber: "123XXX",
				MedicationCode:    "123XXX",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := medicationRepository.Create(tt.args)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := []entity.DroneLoad{}
			err = db.Find(&got).Order("id").Error
			if err != nil {
				t.Errorf("[Error] Cannot retrieve drones %v", err)
			}
			checkDroneLoadEquality(t, []entity.DroneLoad{tt.want}, got)
		})
	}
}

func checkDroneLoadEquality(t *testing.T, want []entity.DroneLoad, got []entity.DroneLoad) {
	t.Helper()
	assert.Equal(t, len(want), len(got))
	for i, droneLoad := range got {
		assert.Equal(t, want[i].DroneSerialNumber, droneLoad.DroneSerialNumber)
		assert.Equal(t, want[i].MedicationCode, droneLoad.MedicationCode)
	}
}
