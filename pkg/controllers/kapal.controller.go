package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKapalController(c echo.Context) error {
	data, err := repository.GetAllKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "Kamu Berhasil Melihat Semua Data Kapal Perikanan",
		"Daftar Kapal Perikanan": data,
	})
}

func GetKapalController(c echo.Context) error {
	kapal, err := repository.GetKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, kapal)

}

func DeleteKapalController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete data kapal dengan id '" + id + "'",
	})

}

func CreateKapalController(c echo.Context) error {
	id := c.Param("id")
	NamaKapal := models.Kapal{}
	c.Bind(&NamaKapal)
	err := repository.CreateNewKapal(NamaKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes insert data kapal dengan id '" + id + "'",
	})

}

func UpdateKapalController(c echo.Context) error {
	// user, err := utils.ParsingJWT(c)
	// if err != nil {
	// 	return c.JSON(500, echo.Map{
	// 		"error": err.Error(),
	// 	})
	// } else if user.Role != "admin" {
	// 	return c.JSON(200, echo.Map{
	// 		"error": "restricted (*only for admin)",
	// 	})
	// }

	NamaKapal := models.Kapal{}
	c.Bind(&NamaKapal)
	err := repository.UpdateKapalById(c.Param("id"), NamaKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaKapal)

}

func GroupingKapalController(c echo.Context) error {
	data, err := repository.GroupingKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "Berhasil Grouping Kapal Berdasarkan Id Perusahaan",
		"Daftar Kapal Perikanan": data,
	})
}
