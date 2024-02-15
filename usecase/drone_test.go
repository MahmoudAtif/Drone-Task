package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/mock"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDroneUseCaseGet(t *testing.T) {
	type mocks struct {
		mockedDroneRepository repository.IDroneRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		want    string
		args    args
		wantErr string
		mocks   mocks
	}{
		{
			name: "Test Get usecase",
			want: `[{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","serial_number":"1","model":"Lightweight","weight":55,"battery_capacity":20,"state":"IDLE"}]`,
			args: args{ctx: context.Background()},
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUseCase := NewDroneUseCase(tt.mocks.mockedDroneRepository)
			got, err := droneUseCase.Get(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestDroneUseCaseCreate(t *testing.T) {
	type mocks struct {
		mockedDroneRepository repository.IDroneRepository
	}
	type args struct {
		ctx     context.Context
		request []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name: "Test Can create drones without validation",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"serial_number": "11","model": "Lightweight","weight": 55.5,"battery_capacity": 100,"state": "IDLE"}]`),
			},
			want: `{"created_drones":[{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","serial_number":"11","model":"Lightweight","weight":55.5,"battery_capacity":100,"state":"IDLE"}]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test Cannot Create drones with wrong serial number",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"serial_number": "sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss","model": "Lightweight","weight_limit": 55.5,"battery_capacity": 55,"state": "IDLE"}]`),
			},
			want: `{"created_drones":[],"errors":["drone sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss: serial number must be 100 characters max"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test Cannot Create drones with wrong model",
			args: args{
				ctx: context.Background(),
				request: []byte(`[{"serial_number": "11","model": "dfdf","weight": 55.5,"battery_capacity": 55,"state": "IDLE"
				}]`),
			},
			want: `{"created_drones":[],"errors":["drone 11: invalid drone models"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test Cannot Create drones with wrong state",
			args: args{
				ctx: context.Background(),
				request: []byte(`[{"serial_number": "11","model": "Lightweight","weight": 55.5,"battery_capacity": 55,"state": "ddd"
				}]`),
			},
			want: `{"created_drones":[],"errors":["drone 11: invalid drone state"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test Cannot Create drones with wrong batery capacity",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"serial_number": "11","model": "Lightweight","weight": 55.5,"battery_capacity": 200,"state": "IDLE"}]`),
			},
			want: `{"created_drones":[],"errors":["drone 11: invalid batery_cabacity precentage"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test Cannot Create drones with wrong weight",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"serial_number": "11","model": "Lightweight","weight": 555.5,"battery_capacity": 100,"state": "IDLE"}]`),
			},
			want: `{"created_drones":[],"errors":["drone 11: weight must be less than or equal 500gr"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
		{
			name: "Test create multiple drones one with validation and one without validation",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"serial_number": "11","model": "Light","weight": 555.5,"battery_capacity": 101,"state": "44"}, {"serial_number": "12","model": "Lightweight","weight": 55.5,"battery_capacity": 100,"state": "IDLE"}]`),
			},
			want: `{"created_drones":[{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","serial_number":"12","model":"Lightweight","weight":55.5,"battery_capacity":100,"state":"IDLE"}],"errors":["drone 11: invalid drone models","drone 11: weight must be less than or equal 500gr","drone 11: invalid batery_cabacity precentage","drone 11: invalid drone state"]}`,
			mocks: mocks{
				mockedDroneRepository: mock.NewMockedDroneRepository(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneUseCase(tt.mocks.mockedDroneRepository)
			got, err := droneUsecase.Create(tt.args.ctx, tt.args.request)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}
