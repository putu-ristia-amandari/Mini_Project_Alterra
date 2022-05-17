package routes

import (
	c "mini_project/pkg/controllers"

	m "mini_project/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()
	m.LoggerMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Home Page Informasi Kedatangan Kapal")
	})

	// Routes Login
	e.POST("/createuser", c.CreateUserController)
	e.POST("/login", c.LoginController)

	// Routes Kedatangan
	e.GET("/kedatangan", c.GetAllKedatanganKapalController, m.IsAuthenticated)
	e.GET("/kedatangan/:id", c.GetKedatanganKapalController, m.IsAuthenticated)
	e.POST("/kedatangan", c.CreateKedatanganKapalController, m.IsAuthenticated)
	e.PUT("/kedatangan/:id", c.UpdateKedatanganKapalController, m.IsAuthenticated)
	e.GET("/grouparea", c.GroupingDaerahController, m.IsAuthenticated)

	// Routes Kapal
	e.GET("/kapal", c.GetAllKapalController, m.IsAuthenticated)
	e.GET("/kapal/:id", c.GetKapalController, m.IsAuthenticated)
	e.POST("/kapal", c.CreateKapalController, m.IsAuthenticated)
	e.PUT("/kapal/:id", c.UpdateKapalController, m.IsAuthenticated)
	e.GET("/grouping", c.GroupingKapalController, m.IsAuthenticated)

	// Routes Pelabuhan
	e.GET("/pelabuhan", c.GetAllPelabuhanController, m.IsAuthenticated)
	e.GET("/pelabuhan/:id", c.GetPelabuhanController, m.IsAuthenticated)
	e.DELETE("/pelabuhan/:id", c.DeletePelabuhanController, m.IsAuthenticated)
	e.POST("/pelabuhan", c.CreatePelabuhanController, m.IsAuthenticated)
	e.PUT("/pelabuhan/:id", c.UpdatePelabuhanController, m.IsAuthenticated)

	// Routes Perusahaan
	e.GET("/perusahaan", c.GetAllPerusahaanController, m.IsAuthenticated)
	e.GET("/perusahaan/:id", c.GetPerusahaanController, m.IsAuthenticated)
	e.DELETE("/perusahaan/:id", c.DeletePerusahaanController, m.IsAuthenticated)
	e.POST("/perusahaan", c.CreatePerusahaanController, m.IsAuthenticated)
	e.PUT("/perusahaan/:id", c.UpdatePerusahaanController, m.IsAuthenticated)

	//Routes Muatan
	e.GET("/muatan", c.GetAllMuatanController, m.IsAuthenticated)
	e.GET("/muatan/:id", c.GetMuatanController, m.IsAuthenticated)
	e.DELETE("/muatan/:id", c.DeleteMuatanController, m.IsAuthenticated)
	e.POST("/muatan", c.CreateMuatanController, m.IsAuthenticated)
	e.PUT("/muatan/:id", c.UpdateMuatanController, m.IsAuthenticated)

	// Routes Alat Tangkap
	e.GET("/alat", c.GetAllAlatTangkapController, m.IsAuthenticated)
	e.GET("/alat/:id", c.GetAlatTangkapController, m.IsAuthenticated)
	e.DELETE("/alat/:id", c.DeleteAlatTangkapController, m.IsAuthenticated)
	e.POST("/alat", c.CreateAlatTangkapController, m.IsAuthenticated)
	e.PUT("/alat/:id", c.UpdateAlatTangkapController, m.IsAuthenticated)

	// Routes Tipe Kapal
	e.GET("/tipe", c.GetAllTipeKapalController, m.IsAuthenticated)
	e.GET("/tipe/:id", c.GetTipeKapalController, m.IsAuthenticated)

	return e
}
