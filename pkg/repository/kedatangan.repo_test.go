package repository

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUnit struct {
	Mock            sqlmock.Sqlmock
	iKedatanganRepo iKedatanganRepo
}

func InisialisasiUnitTest() TestUnit {
	tu := TestUnit{}
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql-mock",
		ServerVersion:             "1.0.0",
		DSN:                       "mysql-mock",
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tu.Mock = mock
	iKedatanganRepo := GetAllKedatanganKapal()(dbGorm)
	tu.iKedatanganRepo = iKedatanganRepo
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `kedatangan_kapal`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "id_kapal", "id_pelabuhan", "id_jenis_muatan", "daerah_penangkapan", "total_tangkapan",
				"total_tangkapan", "tgl_keberangkatan", "tgl_kedatangan", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "[1, 2]", "Laut Lepas Samudra Hindia", "90.000 kg", "2021-12-04T14:20:00Z", "2022-01-01T09:43:00Z",
					"2022-01-01T11:00:00Z", "2022-01-01T11:00:00Z"))
	return tu

}

// func TestGetAllKedatanganKapal(t *testing.T) {
// 	mockKedatanganRepo := new(mocks.KedatanganRepositoryMock)
// 	mockKedatanganData := models.KedatanganKapal{
// 		Id:                 "id",
// 		Id_Kapal:           "id_kapal",
// 		Id_Pelabuhan:       "id_pelabuhan",
// 		Id_Jenis_Muatan:    "id_jenis_muatan",
// 		Daerah_Penangkapan: "daerah_penangkapan",
// 		Total_Tangkapan:    "total_tangkapan",
// 		Tgl_Keberangkatan:  "tgl_keberangkatan",
// 		Tgl_Kedatangan:     "tgl_kedatangan",
// 		Created_At:         "created_at",
// 		Updated_At:         "updated_at",
// 	}
// 	t.Run("success", func(t *testing.T) {
// 		mockKedatanganRepo.On("Read", mock.Anything).Return(nil).Once()

// 		GetAll := models.KedatanganKapal.GetAllKedatanganKapal(mockKedatanganRepo)
// 		err := GetAll.GetAllKedatanganKapal(mockKedatanganData)

// 		assert.NoError(t, err)

// 	})

// 	t.Run("failed", func(t *testing.T) {
// 		mockKedatanganRepo.On("GetAllKedatanganKapalController", mock.Anything).Return(errors.New("Unit Test Error")).Once()
// 		GetAll := models.KedatanganKapal.GetAllKedatanganKapal(mockKedatanganRepo)
// 		err := GetAll.GetAllKedatanganKapal(mockKedatanganData)

// 		assert.NoError(t, err)
// 	})
// }

// var testCases = []struct {
// 	name                 string
// 	path                 string
// 	expectStatus         int
// 	expectBodyStartsWith string
// }{
// 	{
// 		name:                 "berhasil",
// 		path:                 "/kedatangan",
// 		expectBodyStartsWith: "{\"status\":\"success\":[",
// 		expectStatus:         http.StatusOK,
// 	},
// }
// e := echo.New()
// req := httptest.NewRequest(http.MethodGet, "/", nil)
// rec := httptest.NewRecorder()
// c := e.NewContext(req, rec)

// for _, testCase := range testCases {
// 	c.SetPath(testCase.path)

// 	if assert.NoError(t, GetAllKedatanganKapalController(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		body := rec.Body.String()

// 		assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
// 	}
// }
// }
