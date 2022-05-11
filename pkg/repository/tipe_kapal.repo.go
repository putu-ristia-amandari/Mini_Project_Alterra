package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllTipeKapal() ([]models.TipeKapal, error) {
	var ListTipeKapal []models.TipeKapal

	err := database.DB.Find(&ListTipeKapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListTipeKapal, err
}

func GetTipeKapalById(id string) (models.TipeKapal, error) {
	var TipeKapal models.TipeKapal
	err := database.DB.First(&TipeKapal, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return TipeKapal, err
}

func DeleteTipeKapalById(id string) error {
	err := database.DB.Delete(&models.TipeKapal{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewTipeKapal(TipeKapal models.TipeKapal) error {
	err := database.DB.Save(&TipeKapal).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateTipeKapalById(id string, TipeKapal models.TipeKapal) error {
	err := database.DB.Model(&TipeKapal).Where("id = ?", id).Updates(TipeKapal).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
