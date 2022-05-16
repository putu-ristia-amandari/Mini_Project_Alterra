package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

type iKedatanganRepo interface {
	GetAllKedatanganKapal() ([]models.KedatanganKapal, error)
	GetKedatanganKapalById(id string) (models.KedatanganKapal, error)
	DeleteKedatanganKapalById(id string) error
	CreateNewKedatanganKapal(kedatangan models.KedatanganKapal)
	UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error
}

func GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	var listkapal []models.KedatanganKapal

	err := database.DB.Debug().Find(&listkapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return listkapal, err
}

func GetKedatanganKapalById(id string) (models.KedatanganKapal, error) {
	var kedatangan models.KedatanganKapal
	err := database.DB.First(&kedatangan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return kedatangan, err
}

func DeleteKedatanganKapalById(id string) error {
	err := database.DB.Delete(&models.KedatanganKapal{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewKedatanganKapal(kedatangan models.KedatanganKapal) error {
	err := database.DB.Save(&kedatangan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error {
	err := database.DB.Debug().Model(&kedatangan).Where("id = ?", id).Updates(&kedatangan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
