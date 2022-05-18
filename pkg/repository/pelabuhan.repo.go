package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllPelabuhan() ([]models.Pelabuhan, error) {
	var listpelabuhan []models.Pelabuhan

	err := database.DB.Find(&listpelabuhan).Error
	if err != nil {
		fmt.Println(err)
	}
	return listpelabuhan, err
}

func GetPelabuhanById(id string) (models.Pelabuhan, error) {
	var NamaPelabuhan models.Pelabuhan
	err := database.DB.First(&NamaPelabuhan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return NamaPelabuhan, err
}

func DeletePelabuhanById(id string) error {
	err := database.DB.Delete(&models.Pelabuhan{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewPelabuhan(pelabuhan models.Pelabuhan) error {
	err := database.DB.Create(&pelabuhan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdatePelabuhanById(id string, pelabuhan models.Pelabuhan) error {
	err := database.DB.Model(&pelabuhan).Where("id = ?", id).Updates(pelabuhan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
