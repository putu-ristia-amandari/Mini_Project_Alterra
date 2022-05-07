package database

import (
	"fmt"
	"mini_project/pkg/config"
	"mini_project/pkg/models"
)

func GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	var listkapal []models.KedatanganKapal

	err := config.DB.Find(&listkapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return listkapal, err

	// 	res.Status = http.StatusOK
	// 	res.Message = "Success"
	// 	res.Data = listkapal

	// 	return res, nil

}

func GetKedatanganKapalById(id string) (models.KedatanganKapal, error) {
	var kedatangan models.KedatanganKapal
	err := config.DB.First(&kedatangan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return kedatangan, err
}

func DeleteKedatanganKapalById(id string) error {
	err := config.DB.Delete(&models.KedatanganKapal{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewKedatanganKapal(kedatangan models.KedatanganKapal) error {
	err := config.DB.Save(&kedatangan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error {
	err := config.DB.Model(&kedatangan).Where("id = ?", id).Updates(kedatangan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
