package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

type KapalRepository interface {
	GetAllKapal() ([]models.Kapal, error)
	GetKapalById(id string) (models.Kapal, error)
	CreateNewKapal(NamaKapal models.Kapal) error
	UpdateKapalById(id string, NamaKapal models.Kapal) error
	GroupingKapal() ([]models.Kapal, error)
}

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

func GroupingKapal() ([]models.Kapal, error) {
	var Data []models.Kapal
	err := database.DB.Select("nama_kapal", "id_perusahaan").Order("id_perusahaan, nama_kapal desc").
		Find(&Data).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Data, err
}
