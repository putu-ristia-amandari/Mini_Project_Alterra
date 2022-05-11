package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

func GetAllPerusahaanKapal() ([]models.PerusahaanKapal, error) {
	var listperusahaan []models.PerusahaanKapal

	err := database.DB.Find(&listperusahaan).Error
	if err != nil {
		fmt.Println(err)
	}
	return listperusahaan, err
}

func GetPerusahaanById(id string) (models.PerusahaanKapal, error) {
	var NamaPerusahaan models.PerusahaanKapal
	err := database.DB.First(&NamaPerusahaan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return NamaPerusahaan, err
}

func DeletePerusahaanById(id string) error {
	err := database.DB.Delete(&models.PerusahaanKapal{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewPerusahaan(perusahaan models.PerusahaanKapal) error {
	err := database.DB.Save(&perusahaan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdatePerusahaanById(id string, perusahaan models.PerusahaanKapal) error {
	err := database.DB.Model(&perusahaan).Where("id = ?", id).Updates(perusahaan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
