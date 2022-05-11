package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	var listkapal []models.KedatanganKapal

	err := database.DB.Find(&listkapal).Error
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
	err := database.DB.Model(&kedatangan).Where("id = ?", id).Updates(kedatangan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
