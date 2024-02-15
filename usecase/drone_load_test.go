package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/mock"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDroneLoadUseCaseCreate(t *testing.T) {
	type mocks struct {
		MockedDroneLoadRepository  repository.IDroneLoadRepository
		mockedDroneRepository      repository.IDroneRepository
		MockedMedicationRepository repository.IMedicationRepository
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
			name: "Test create Drone load without validation",
			args: args{
				ctx:     context.Background(),
				request: []byte(`{"drone_serial_number": "123", "medication_code" : "10"}`),
			},
			want: `{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","drone_serial_number":"123","medication_code":"10"}`,
			mocks: mocks{
				MockedDroneLoadRepository:  mock.NewMockedDroneLoadRepository(),
				mockedDroneRepository:      mock.NewMockedDroneRepository(),
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot load drone with invalid drone serial number",
			args: args{
				ctx:     context.Background(),
				request: []byte(`{"drone_serial_number": "soso", "medication_code" : "10"}`),
			},
			want: `["Invalid drone serial number"]`,
			mocks: mocks{
				MockedDroneLoadRepository:  mock.NewMockedDroneLoadRepository(),
				mockedDroneRepository:      mock.NewMockedDroneRepository(),
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot load drone with invalid medication code",
			args: args{
				ctx:     context.Background(),
				request: []byte(`{"drone_serial_number": "123", "medication_code" : "toto"}`),
			},
			want: `["Invalid mecication code"]`,
			mocks: mocks{
				MockedDroneLoadRepository:  mock.NewMockedDroneLoadRepository(),
				mockedDroneRepository:      mock.NewMockedDroneRepository(),
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot load drone with invalid drone serial number and medication code",
			args: args{
				ctx:     context.Background(),
				request: []byte(`{"drone_serial_number": "soso", "medication_code" : "toto"}`),
			},
			want: `["Invalid drone serial number","Invalid mecication code"]`,
			mocks: mocks{
				MockedDroneLoadRepository:  mock.NewMockedDroneLoadRepository(),
				mockedDroneRepository:      mock.NewMockedDroneRepository(),
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneLoadUseCase := NewDroneLoadUseCase(
				tt.mocks.MockedDroneLoadRepository,
				tt.mocks.mockedDroneRepository,
				tt.mocks.MockedMedicationRepository,
			)
			got, err := droneLoadUseCase.Create(tt.args.ctx, tt.args.request)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}
