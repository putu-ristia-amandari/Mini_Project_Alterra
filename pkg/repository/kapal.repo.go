package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllKapal() ([]models.Kapal, error) {
	var ListNamaKapal []models.Kapal

	err := database.DB.Find(&ListNamaKapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListNamaKapal, err
}

func GetKapalById(id string) (models.Kapal, error) {
	var kapal models.Kapal
	err := database.DB.First(&kapal, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return kapal, err
}

func DeleteKapalById(id string) error {
	err := database.DB.Delete(&models.Kapal{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewKapal(NamaKapal models.Kapal) error {
	err := database.DB.Save(&NamaKapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateKapalById(id string, NamaKapal models.Kapal) error {
	err := database.DB.Model(&NamaKapal).Where("id = ?", id).Updates(NamaKapal).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
