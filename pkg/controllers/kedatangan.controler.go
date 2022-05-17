package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KedatanganController struct {
	datang repository.KedatanganRepository
}

func GetAllKedatanganKapalController(c echo.Context) error {
	data, err := repository.GetAllKedatanganKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                           "Kamu Berhasil Melihat Semua Data Kedatangan Kapal",
		"Daftar Kedatangan Kapal Perikanan": data,
	})
}

func GetKedatanganKapalController(c echo.Context) error {
	kapal, err := repository.GetKedatanganKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func CreateKedatanganKapalController(c echo.Context) error {
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := repository.CreateNewKedatanganKapal(kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Kedatangan Kapal",
	})

}

func UpdateKedatanganKapalController(c echo.Context) error {
	kapal := models.KedatanganKapal{}
	c.Bind(&kapal)
	err := repository.UpdateKedatanganKapalById(c.Param("id"), kapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Memperbaharui Data",
	})

}

func GroupingDaerahController(c echo.Context) error {
	data, err := repository.GroupingDaerah()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "Berhasil Grouping Kapal Berdasarkan Daerah Penangkapannya",
		"Daftar Kapal Perikanan": data,
	})
}

func NewKedatanganController(datang repository.KedatanganRepository) KedatanganController {
	return KedatanganController{
		datang: datang,
	}
}

// func SubQueryNamaKapalController(c echo.Context) error {
// 	kapal := models.KedatanganKapal{}
// 	kapal, err := repository.SubQueryNamaKapal()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

// 	}
// 	return c.JSON(http.StatusOK, kapal)

// }
