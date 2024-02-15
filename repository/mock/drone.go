package mock

import (
	"drone-task/repository"
	"drone-task/repository/entity"
)

type MockedDroneRepository struct{}

func NewMockedDroneRepository() repository.IDroneRepository {
	return MockedDroneRepository{}
}

func (m MockedDroneRepository) Get() ([]entity.Drone, error) {
	drones := []entity.Drone{
		{SerialNumber: "1", Model: "Lightweight", Weight: 55, BatteryCapacity: 20, State: "IDLE"},
	}
	return drones, nil

}

func (m MockedDroneRepository) Create(drones []entity.Drone) ([]entity.Drone, error) {
	return drones, nil
}

func (m MockedDroneRepository) GetById(id int) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (m MockedDroneRepository) Delete(id int) error {
	return nil
}

func (m MockedDroneRepository) Update(drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (m MockedDroneRepository) GetBySerialNumber(serialNumber string) (entity.Drone, error) {
	drone := entity.Drone{}
	return drone, nil
}

func (m MockedDroneRepository) UpdateByID(id int, fields map[string]interface{}) (entity.Drone, error) {
	drone := entity.Drone{}
	return drone, nil
}

func (m MockedDroneRepository) Filter(filters entity.DroneFilters) ([]entity.Drone, error) {
	drones := []entity.Drone{
		{
			SerialNumber:    "123",
			Model:           "LightWeight",
			Weight:          100,
			BatteryCapacity: 50,
			State:           "IDLE",
		},
		{
			SerialNumber:    "1234",
			Model:           "Lightweight",
			Weight:          55,
			BatteryCapacity: 20,
			State:           "IDLE",
		},
		{
			SerialNumber:    "12345",
			Model:           "Lightweight",
			Weight:          55,
			BatteryCapacity: 0,
			State:           "LOADED",
		},
	}
	filteredDrones := []entity.Drone{}

	for _, drone := range drones {
		if len(filters.SerialNumbers) > 0 {
			found := false
			for _, serialNumber := range filters.SerialNumbers {
				if drone.SerialNumber == serialNumber {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		filteredDrones = append(filteredDrones, drone)
	}
	return filteredDrones, nil
}

// func (m MockedDocumentRepository) Get(ctx context.Context, filters entity.DocumentFilters) ([]entity.Document, int64, error) {
// 	documents := []entity.Document{
// 		{Model: entity.Model{ID: 0}, TenantID: ctx.Value(keys.TENANT_ID).(int), FolioInternalID: "100"},
// 		{Model: entity.Model{ID: 1}, TenantID: ctx.Value(keys.TENANT_ID).(int), FolioInternalID: "200", Status: "Pending"},
// 		{Model: entity.Model{ID: 2}, TenantID: ctx.Value(keys.TENANT_ID).(int), FolioInternalID: "300", Status: "Excluded"},
// 		{Model: entity.Model{ID: 3}, TenantID: ctx.Value(keys.TENANT_ID).(int), FolioInternalID: "400", Status: "Valid", Document: []byte(`{}`), DocumentUUID: "uuid1"},
// 		{Model: entity.Model{ID: 4}, TenantID: ctx.Value(keys.TENANT_ID).(int), FolioInternalID: "500", Status: "Valid", Document: []byte(`{}`), DocumentUUID: "uuid2"},
// 	}

// 	filteredDocuments := []entity.Document{}

// 	for _, doc := range documents {
// 		if len(filters.Ids) > 0 {
// 			found := false
// 			for _, id := range filters.Ids {
// 				if doc.ID == id {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				continue
// 			}
// 		}
// 		if filters.Reference != "" {
// 			if doc.Document != nil {
// 				jsonDocument := useCaseEntity.Document{}
// 				err := json.Unmarshal(doc.Document, &jsonDocument)
// 				if err != nil {
// 					log.Printf("[Error]: %v", err.Error())
// 					return nil, 0, err
// 				}
// 				if !slices.Contains(jsonDocument.References, filters.Reference) {
// 					continue
// 				}
// 			} else {
// 				continue
// 			}
// 		}

// 		filteredDocuments = append(filteredDocuments, doc)
// 	}

// 	return filteredDocuments, 0, nil
// }
