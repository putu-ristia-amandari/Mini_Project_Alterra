package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

type iKedatanganRepo interface {
	GetAllKedatanganKapal() ([]models.KedatanganKapal, error)
	GetKedatanganKapalById(id string) (models.KedatanganKapal, error)
	DeleteKedatanganKapalById(id string) error
	CreateNewKedatanganKapal(kedatangan models.KedatanganKapal)
	UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error
}

func GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	var listkapal []models.KedatanganKapal

	err := database.DB.Find(&listkapal).Debug().Error
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
	err := database.DB.Model(&kedatangan).Where("id = ?", id).Updates(&kedatangan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GroupingDaerah() ([]models.KedatanganKapal, error) {
	var Data []models.KedatanganKapal
	err := database.DB.Select("daerah_penangkapan", "id_kapal").Order("daerah_penangkapan, id_kapal asc").
		Find(&Data).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Data, err
}

// func SubQueryNamaKapal([]models.KedatanganKapal, []models.Kapal) error {
// 	var (
// 		kedatangan []models.KedatanganKapal
// 		kapal      []models.Kapal
// 	)
// 	subQuery1 := database.DB.Select("id_kapal").Find(&kedatangan)
// 	subQuery2 := database.DB.Select("nama_kapal").Find(&kapal)

// 	err := database.DB.Raw("(?) as u, (?) as p", subQuery1, subQuery2).Find(&rows).Debug().Error

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return err
// }
