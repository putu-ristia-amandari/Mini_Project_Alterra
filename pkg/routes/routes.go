package routes

import (
	"mini_project/pkg/controllers"
	// "mini_project/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	// Routes Kedatangan Kapal
	e.GET("/kedatangan", controllers.GetAllKedatanganKapalController)
	e.GET("/kedatangan/:id", controllers.GetKedatanganKapalController)
	e.DELETE("/kedatangan/:id", controllers.DeleteKedatanganKapalController)
	e.POST("/kedatangan", controllers.CreateKedatanganKapalController)
	e.PUT("/kedatangan/:id", controllers.UpdateKedatanganKapalController)

	// Routes Kapal
	// e.GET("/kapal", controllers.GetAllKapalController, middleware.IsAuthenticated)
	e.GET("/kapal/:id", controllers.GetKapalController)
	e.DELETE("/kapal/:id", controllers.DeleteKapalController)
	e.POST("/kapal", controllers.CreateKapalController)
	e.PUT("/kapal/:id", controllers.UpdateKapalController)

	// Routes Perusahaan Kapal
	e.GET("/perusahaan", controllers.GetAllPerusahaanController)
	e.GET("/perusahaan/:id", controllers.GetPerusahaanController)
	e.DELETE("/perusahaan/:id", controllers.DeletePerusahaanController)
	e.POST("/perusahaan", controllers.CreatePerusahaanController)
	e.PUT("/perusahaan/:id", controllers.UpdatePerusahaanController)

	// Routes Pelabuhan
	e.GET("/pelabuhan", controllers.GetAllPelabuhanController)
	e.GET("/pelabuhan/:id", controllers.GetPelabuhanController)
	e.DELETE("/pelabuhan/:id", controllers.DeletePelabuhanController)
	e.POST("/pelabuhan", controllers.CreatePelabuhanController)
	e.PUT("/pelabuhan/:id", controllers.UpdatePelabuhanController)

	// Routes Muatan
	e.GET("/muatan", controllers.GetAllMuatanController)
	e.GET("/muatan/:id", controllers.GetMuatanController)
	e.DELETE("/muatan/:id", controllers.DeleteMuatanController)
	e.POST("/muatan", controllers.CreateMuatanController)
	e.PUT("/muatan/:id", controllers.UpdateMuatanController)

	// Routes Alat Tangkap
	e.GET("/alat", controllers.GetAllAlatTangkapController)
	e.GET("/alat/:id", controllers.GetAlatTangkapController)
	e.DELETE("/alat/:id", controllers.DeleteAlatTangkapController)
	e.POST("/alat", controllers.CreateAlatTangkapController)
	e.PUT("/alat/:id", controllers.UpdateAlatTangkapController)

	// Routes Tipe Kapal
	e.GET("/tipe", controllers.GetAllTipeKapalController)
	e.GET("/tipe/:id", controllers.GetTipeKapalController)
	e.DELETE("/tipe/:id", controllers.DeleteTipeKapalController)
	e.POST("/tipe", controllers.CreateTipeKapalController)
	e.PUT("/tipe/:id", controllers.UpdateTipeKapalController)

	// Routes Login
	// e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("/login", controllers.LoginCheck)

	//Routes Validation
	// e.GET("/validation", controllers.LoginUserValidation)

	return e
}
