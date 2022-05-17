package repository

import (
	"mini_project/pkg/models"
)

type MockKapalRepository struct {
	data []models.Kapal
	// Mock mock.Mock
}

func (repository *MockKapalRepository) GetAllKapal() ([]models.Kapal, error) {
	return repository.data, nil
}

// func (repository *MockKapalRepository) GetKapalById(id string) (models.Kapal, error) {
// 	arguments := repository.Mock.Called(id)
// 	if arguments.Get(0) == nil {
// 		return nil
// 	} else {
// 		kapal := arguments.Get(0).(data)
// 		return &kapal
// 	}
// }

// func (repository *MockKapalRepository) CreateNewKedatanganKapal(kedatangan models.Kapal) error {
// 	repository.data = append(repository.data, kapal)
// 	return kapal, nil
// }

// func (repository *MockKapalRepository) UpdateKapalById(id string, kedatangan models.Kapal) error {
// 	args := repository.Called(id)
// 	return agrs.Int(0), args.Error(1)
// }

// func (repository *MockKapalRepository) GroupingKapal() ([]models.Kapal, error) {
// 	return repository.data, nil
// }

func NewMockKapalRepository() *MockKapalRepository {
	return &MockKapalRepository{
		data: []models.Kapal{},
	}
}
