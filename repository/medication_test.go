package repository

import (
	"drone-task/repository/entity"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMedicationRepositoryCreate(t *testing.T) {
	db := dbClient.Begin()
	defer db.Rollback()
	medicationRepository := NewMedicationRepository(db)
	type args struct {
		medications []entity.Medication
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Medication
		wantErr string
	}{
		{
			name: "Test can create multiple medication",
			args: args{
				medications: []entity.Medication{
					{Name: "kolosolo", Code: "1234", Weight: 100, Image: "xoxox.png"},
					{Name: "kolosolo2", Code: "12345", Weight: 100, Image: "xoxox2.png"},
				},
			},
			want: []entity.Medication{
				{Name: "kolosolo", Code: "1234", Weight: 100, Image: "xoxox.png"},
				{Name: "kolosolo2", Code: "12345", Weight: 100, Image: "xoxox2.png"},
			},
		},
		{
			name: "Test can create multiple medication",
			args: args{
				medications: []entity.Medication{
					{Name: "kolosolo3", Code: "123", Weight: 100, Image: "xoxox3.png"},
					{Name: "kolosolo4", Code: "123", Weight: 100, Image: "xoxox4.png"},
				},
			},
			wantErr: "duplicated key not allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := medicationRepository.Create(tt.args.medications)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := []entity.Medication{}
			err = db.Find(&got).Order("id").Error
			if err != nil {
				t.Errorf("[Error] Cannot retrieve drones %v", err)
			}
			checkMedicationsEquality(t, tt.want, got)
		})
	}
}

func checkMedicationsEquality(t *testing.T, want []entity.Medication, got []entity.Medication) {
	t.Helper()
	assert.Equal(t, len(want), len(got))
	for i, medication := range got {
		assert.Equal(t, want[i].Name, medication.Name)
		assert.Equal(t, want[i].Code, medication.Code)
		assert.Equal(t, want[i].Weight, medication.Weight)
		assert.Equal(t, want[i].Image, medication.Image)
	}
}
