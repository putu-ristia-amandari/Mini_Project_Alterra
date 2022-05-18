package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllMuatanKapal() ([]models.Muatan, error) {
	var listmuatan []models.Muatan

	err := database.DB.Find(&listmuatan).Error
	if err != nil {
		fmt.Println(err)
	}
	return listmuatan, err
}

func GetMuatanKapalById(id string) (models.Muatan, error) {
	var JenisMuatan models.Muatan
	err := database.DB.First(&JenisMuatan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return JenisMuatan, err
}

func DeleteMuatanKapalById(id string) error {
	err := database.DB.Delete(&models.Muatan{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewMuatanKapal(MuatanKapal models.Muatan) error {
	err := database.DB.Create(&MuatanKapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateMuatanKapalById(id string, MuatanKapal models.Muatan) error {
	err := database.DB.Model(&MuatanKapal).Where("id = ?", id).Updates(MuatanKapal).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
