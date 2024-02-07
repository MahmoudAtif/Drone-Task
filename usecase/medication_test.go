package usecase

import (
	"context"
	"drone-task/repository"
	"drone-task/repository/mock"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMedicationUseCaseCreate(t *testing.T) {
	type mocks struct {
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
			name: "Test create multiple medications without validation",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"name": "kolosolo","code": "123","weight": 100,"image": "gogo.png"}, {"name": "kolosolo","code": "123","weight": 100,"image": "gogo.png"}]`),
			},
			want: `{"created_medications":[{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","name":"kolosolo","weight":100,"code":"123","image":"gogo.png"},{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","name":"kolosolo","weight":100,"code":"123","image":"gogo.png"}]}`,
			mocks: mocks{
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot create medication with wrong name format",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"name": "kolo solo","code": "123","weight": 100,"image": "gogo.png"}]`),
			},
			want: `{"created_medications":[],"errors":["medication kolo solo: Invalid name. allowed only: letters, numbers, '-', '_'."]}`,
			mocks: mocks{
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot create medication with wrong code format",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"name": "kolosolo","code": "ssss123","weight": 100,"image": "gogo.png"}]`),
			},
			want: `{"created_medications":[],"errors":["medication kolosolo: Invalid code. allowed only: upper case letters, numbers, '_'."]}`,
			mocks: mocks{
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test cannot create medication with wrong image extention",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"name": "kolosolo","code": "XXXX123","weight": 100,"image": "gogo.txt"}]`),
			},
			want: `{"created_medications":[],"errors":["medication kolosolo: Invalid image. allowed extentions: (.jpg, .jpeg, .png)"]}`,
			mocks: mocks{
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
		{
			name: "Test create multiple medications one with validation and one without validation",
			args: args{
				ctx:     context.Background(),
				request: []byte(`[{"name": "kolosolo","code": "123","weight": 100,"image": "gogo.png"}, {"name": "kolo solo","code": "1234","weight": 100,"image": "gogo.png"}]`),
			},
			want: `{"created_medications":[{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","name":"kolosolo","weight":100,"code":"123","image":"gogo.png"}],"errors":["medication kolo solo: Invalid name. allowed only: letters, numbers, '-', '_'."]}`,
			mocks: mocks{
				MockedMedicationRepository: mock.NewMockedMedicationRepository(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUseCase := NewMedicationUseCase(tt.mocks.MockedMedicationRepository)
			got, err := medicationUseCase.Create(tt.args.ctx, tt.args.request)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}
