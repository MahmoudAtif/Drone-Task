package repository

import (
	"drone-task/repository/entity"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDroneRepositoryCreate(t *testing.T) {
	txn := dbClient.Begin()
	defer txn.Rollback()
	droneRepository := NewDroneRepository(txn)
	type args struct {
		drones []entity.Drone
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Drone
		wantErr string
	}{
		{
			name: "test can create multiple drones",
			args: args{
				drones: []entity.Drone{
					{SerialNumber: "1", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
					{SerialNumber: "2", Model: "Lightweight", Weight: 30, BatteryCapacity: 50, State: "IDLE"},
				},
			},
			want: []entity.Drone{
				{SerialNumber: "1", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
				{SerialNumber: "2", Model: "Lightweight", Weight: 30, BatteryCapacity: 50, State: "IDLE"},
			},
		},
		{
			name: "test cannot create drone with the same serial number",
			args: args{
				drones: []entity.Drone{
					{SerialNumber: "2", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
					{SerialNumber: "2", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
				},
			},
			wantErr: "duplicated key not allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := droneRepository.Create(tt.args.drones)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := []entity.Drone{}
			err = txn.Find(&got).Order("id").Error
			if err != nil {
				t.Errorf("[Error] Cannot retrieve drones %v", err)
			}
			checkDronesEquality(t, tt.want, got)
		})
	}
}

func TestDroneRepositoryGet(t *testing.T) {
	txn := dbClient.Begin()
	defer txn.Rollback()
	droneRepository := NewDroneRepository(txn)
	drones := []entity.Drone{
		{SerialNumber: "1", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
		{SerialNumber: "2", Model: "Lightweight", Weight: 30, BatteryCapacity: 50, State: "IDLE"},
	}
	err := txn.Create(&drones).Error
	if err != nil {
		t.Errorf("[Error] Cannot create Drones: %v", err)
	}
	tests := []struct {
		name    string
		want    []entity.Drone
		wantErr bool
	}{
		{
			name: "Test drones data returned from Get Method",
			want: []entity.Drone{drones[0], drones[1]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := droneRepository.Get()
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			checkDronesEquality(t, tt.want, got)
		})
	}
}

func checkDronesEquality(t *testing.T, want []entity.Drone, got []entity.Drone) {
	t.Helper()
	assert.Equal(t, len(want), len(got))
	for i, drone := range got {
		assert.Equal(t, want[i].SerialNumber, drone.SerialNumber)
		assert.Equal(t, want[i].Model, drone.Model)
		assert.Equal(t, want[i].Weight, drone.Weight)
		assert.Equal(t, want[i].BatteryCapacity, drone.BatteryCapacity)
		assert.Equal(t, want[i].State, drone.State)
	}
}

func TestDroneRepositoryDelete(t *testing.T) {
	txn := dbClient.Begin()
	defer txn.Rollback()
	droneRepository := NewDroneRepository(txn)
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Drone
		wantErr bool
	}{
		{
			name: "Test Can delete drone by id",
			args: args{
				id: 1,
			},
			want: entity.Drone{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := droneRepository.Delete(tt.args.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := entity.Drone{}
			err = dbClient.Find(&got, tt.args.id).Error
			if err != nil {
				t.Errorf("[Error] Cannot retrieve credentials: %v", err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDroneRepositoryUpdate(t *testing.T) {
	txn := dbClient.Begin()
	defer txn.Rollback()
	droneRepository := NewDroneRepository(txn)
	type args struct {
		drone entity.Drone
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Drone
		wantErr bool
	}{
		{
			name: "Test Can update drone date",
			args: args{
				drone: entity.Drone{
					AbstractModel: entity.AbstractModel{
						ID: 1,
					},
					SerialNumber:    "12345",
					Model:           "Lightweight",
					Weight:          20,
					BatteryCapacity: 50,
					State:           "IDLE",
				},
			},
			want: []entity.Drone{
				{
					AbstractModel: entity.AbstractModel{
						ID: 1,
					},
					SerialNumber:    "12345",
					Model:           "Lightweight",
					Weight:          20,
					BatteryCapacity: 50,
					State:           "IDLE",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := droneRepository.Update(tt.args.drone)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			checkDronesEquality(t, tt.want, []entity.Drone{got})
		})
	}
}
