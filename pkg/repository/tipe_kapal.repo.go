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
