package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllAlatTangkap() ([]models.AlatTangkap, error) {
	var ListAlatTangkap []models.AlatTangkap

	err := database.DB.Find(&ListAlatTangkap).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListAlatTangkap, err
}

func GetAlatTangkapById(id string) (models.AlatTangkap, error) {
	var AlatTangkap models.AlatTangkap
	err := database.DB.First(&AlatTangkap, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return AlatTangkap, err
}

func DeleteAlatTangkapById(id string) error {
	err := database.DB.Delete(&models.AlatTangkap{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewAlatTangkap(AlatTangkap models.AlatTangkap) error {
	err := database.DB.Save(&AlatTangkap).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateAlatTangkapById(id string, AlatTangkap models.AlatTangkap) error {
	err := database.DB.Model(&AlatTangkap).Where("id = ?", id).Updates(AlatTangkap).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
