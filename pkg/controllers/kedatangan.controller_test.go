package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"mini_project/pkg/models"
// 	"mini_project/pkg/repository"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestGetAllKedatanganKapalController(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", nil)
// 	req.Header.Set("Content-Type", "application/jason")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/")

// 	kedatangan := repository.NewMockKedatanganRepository()
// 	kedatangan.Add(models.KedatanganKapal{Id: 1, Id_Kapal: 1, Id_Pelabuhan: 1, Id_Jenis_Muatan: [3, 8], Daerah_Penangkapan: "Laut Lepas Samudra Hindia",
// 		Total_Tangkapan: "1000 kg", Tgl_Keberangkatan: "2021-12-04 14:20:00", Tgl_Kedatangan: "2022-12-04 14:20:00",
// 		Created_At: "2022-12-04 14:20:00", Updated_At: "2022-12-04 14:20:00"})

// 	kedatangan.Add(models.KedatanganKapal{Id: 2, Id_Kapal: 2, Id_Pelabuhan: 2, Id_Jenis_Muatan: [2, 5], Daerah_Penangkapan: "ZEE 718 Laut Aru",
// 		Total_Tangkapan: "500 kg", Tgl_Keberangkatan: "2021-12-01 14:25:00", Tgl_Kedatangan: "2022-12-01 09:20:00",
// 		Created_At: "2022-12-01 09:20:00", Updated_At: "2022-12-01 09:20:00"})

// 	kapal := NewGetAllKedatanganKapalController(kedatangan)
// 	kapal.Get(c)

// 	var listkapal []models.KedatanganKapal
// 	if err := json.Unmarshal(rec.Body.Bytes(), &listkapal); err != nil {
// 		t.Errorf("unmarshalling returned list kapal failed")
// 	}
// 	if len(listkapal) != 2 {
// 		t.Errorf("Expecting len(listkapal) to be 2, get %d", len(listkapal))
// 	}
// 	if listkapal[0].Id != 1 {
// 		t.Errorf("Expecting listkapal[0].Id to be 1, get %s", listkapal[0].Id)
// 	}
// 	if listkapal[1].Id != 2 {
// 		t.Errorf("Expecting listkapal[0].Id to be 1, get %s", listkapal[0].Id)
// 	}
// }
