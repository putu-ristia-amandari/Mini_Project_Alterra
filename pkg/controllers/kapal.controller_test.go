package controllers

// import (
// 	"encoding/json"
// 	"mini_project/pkg/models"
// 	"mini_project/pkg/repository"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestGetAllKapalController(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set("Content-Type", "application/jason")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/")

// 	kr := repository.NewMockKapalRepository()
// 	kr.Add(models.Kapal{Id: 2, Id_Perusahaan: 3, Id_Tipe_Kapal: 1, Nama_Kapal: "BANDAR NELAYAN", Ukuran_Kapal: "200 GT",
// 		Kekuatan_Mesin: "370 PK", Call_Sign_Kapal: "YOO5608", Nama_Nahkoda: "AMIR"})

// 	kr.Add(models.Kapal{Id: 5, Id_Perusahaan: 3, Id_Tipe_Kapal: 1, Nama_Kapal: "NELAYAN BIRU", Ukuran_Kapal: "150 GT",
// 		Kekuatan_Mesin: "250 PK", Call_Sign_Kapal: "YAK6808", Nama_Nahkoda: "HASAN"})

// 	kc := NewKapalController(kr)
// 	kc.GetAllKapal(c)

// 	var namakapal []models.Kapal
// 	if err := json.Unmarshal(rec.Body.Bytes(), &namakapal); err != nil {
// 		t.Errorf("unmarshalling returned list kapal failed")
// 	}
// 	if len(namakapal) != 2 {
// 		t.Errorf("Expecting len(namakapal) to be 2, get %d", len(namakapal))
// 	}
// 	if namakapal[0].Id != 3 {
// 		t.Errorf("Expecting nama kapal[0].Id to be 3, get %s", namakapal[0].Id)
// 	}
// 	if namakapal[1].Id != 5 {
// 		t.Errorf("Expecting nama kapal[0].Id to be 5, get %s", namakapal[0].Id)
// 	}
// }
