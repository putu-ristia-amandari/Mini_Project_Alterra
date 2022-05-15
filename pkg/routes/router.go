package routes

import (
	"mini_project/pkg/controllers"

	g "mini_project/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()
	g.LoggerMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	// Routes Login
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/createuser", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	// Routes Kedatangan
	e.GET("/kedatangan", controllers.GetAllKedatanganKapalController)
	e.GET("/kedatangan/:id", controllers.GetKedatanganKapalController)
	e.DELETE("/kedatangan/:id", controllers.DeleteKedatanganKapalController, g.IsAuthenticated)
	e.POST("/kedatangan", controllers.CreateKedatanganKapalController, g.IsAuthenticated)
	e.PUT("/kedatangan/:id", controllers.UpdateKedatanganKapalController, g.IsAuthenticated)

	// Routes Kapal
	e.GET("/kapal", controllers.GetAllKapalController)
	e.GET("/kapal/:id", controllers.GetKapalController)
	e.DELETE("/kapal/:id", controllers.DeleteKapalController, g.IsAuthenticated)
	e.POST("/kapal", controllers.CreateKapalController, g.IsAuthenticated)
	e.PUT("/kapal/:id", controllers.UpdateKapalController, g.IsAuthenticated)
	e.GET("/grouping", controllers.GroupingKapalController)

	// Routes Pelabuhan
	e.GET("/pelabuhan", controllers.GetAllPelabuhanController)
	e.GET("/pelabuhan/:id", controllers.GetPelabuhanController)
	e.DELETE("/pelabuhan/:id", controllers.DeletePelabuhanController, g.IsAuthenticated)
	e.POST("/pelabuhan", controllers.CreatePelabuhanController, g.IsAuthenticated)
	e.PUT("/pelabuhan/:id", controllers.UpdatePelabuhanController, g.IsAuthenticated)

	// Routes Perusahaan
	e.GET("/perusahaan", controllers.GetAllPerusahaanController)
	e.GET("/perusahaan/:id", controllers.GetPerusahaanController)
	e.DELETE("/perusahaan/:id", controllers.DeletePerusahaanController, g.IsAuthenticated)
	e.POST("/perusahaan", controllers.CreatePerusahaanController, g.IsAuthenticated)
	e.PUT("/perusahaan/:id", controllers.UpdatePerusahaanController, g.IsAuthenticated)

	//Routes Muatan
	e.GET("/muatan", controllers.GetAllMuatanController)
	e.GET("/muatan/:id", controllers.GetMuatanController)
	e.DELETE("/muatan/:id", controllers.DeleteMuatanController, g.IsAuthenticated)
	e.POST("/muatan", controllers.CreateMuatanController, g.IsAuthenticated)
	e.PUT("/muatan/:id", controllers.UpdateMuatanController, g.IsAuthenticated)

	// Routes Alat Tangkap
	e.GET("/alat", controllers.GetAllAlatTangkapController)
	e.GET("/alat/:id", controllers.GetAlatTangkapController)
	e.DELETE("/alat/:id", controllers.DeleteAlatTangkapController, g.IsAuthenticated)
	e.POST("/alat", controllers.CreateAlatTangkapController, g.IsAuthenticated)
	e.PUT("/alat/:id", controllers.UpdateAlatTangkapController, g.IsAuthenticated)

	// Routes Tipe Kapal
	e.GET("/tipe", controllers.GetAllTipeKapalController)
	e.GET("/tipe/:id", controllers.GetTipeKapalController)
	e.DELETE("/tipe/:id", controllers.DeleteTipeKapalController, g.IsAuthenticated)
	e.POST("/tipe", controllers.CreateTipeKapalController, g.IsAuthenticated)
	e.PUT("/tipe/:id", controllers.UpdateTipeKapalController, g.IsAuthenticated)

	return e
}
