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
				request: []byte(`{"drone_serial_number": "12345", "medication_code" : "1235"}`),
			},
			want: `{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","drone_serial_number":"12345","medication_code":"1235"}`,
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
