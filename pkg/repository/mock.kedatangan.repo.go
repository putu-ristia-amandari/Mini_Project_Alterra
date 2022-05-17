package repository

import (
	"mini_project/pkg/models"
)

type MockKedatanganRepository struct {
	data []models.KedatanganKapal
	// Mock mock.Mock
}

func (repository *MockKedatanganRepository) GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	return repository.data, nil
}

// func (repository *MockKedatanganRepository) GetKedatanganKapalById(id string) (models.KedatanganKapal, error) {
// 	arguments := repository.Mock.Called(id)
// 	if arguments.Get(0) == nil {
// 		return nil
// 	} else {
// 		kedatangan := arguments.Get(0).(data)
// 		return &kedatangan
// 	}
// }

// func (repository *MockKedatanganRepository) CreateNewKedatanganKapal(kedatangan models.KedatanganKapal) error {
// 	repository.data = append(repository.data, kedatangan)
// 	return kedatangan, nil
// }

// func (repository *MockKedatanganRepository) UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error {
// 	args := repository.Called(id)
// 	return agrs.Int(0), args.Error(1)
// }

// func (repository *MockKedatanganRepository) GroupingDaerah() ([]models.KedatanganKapal, error) {
// 	return repository.data, nil
// }

func NewMockKedatanganRepository() *MockKedatanganRepository {
	return &MockKedatanganRepository{
		data: []models.KedatanganKapal{},
	}
}
