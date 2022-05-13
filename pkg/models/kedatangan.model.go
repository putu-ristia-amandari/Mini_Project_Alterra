package models

import (
	"time"
)

type KedatanganKapal struct {
	Id       int `json:"id" gorm:"primaryKey" form:"id"`
	Id_Kapal int `json:"id_kapal" form:"id_kapal"`
	// Kapal              Kapal     `json:"kapal" gorm:foreignKey:IdKedatanganKapal;references:Id`
	Id_Pelabuhan int `json:"id_pelabuhan" form:"id_pelabuhan"`
	// Pelabuhan          Pelabuhan `json:"pelabuhan" gorm:foreignKey:IdKedatanganKapal;references:Id`
	Id_Jenis_Muatan int `json:"id_jenis_muatan" form:"id_jenis_muatan"`
	// Jenis_Muatan       Muatan    `json:"jenis_muatan" gorm:foreignKey:IdKedatanganKapal;references:Id`
	Daerah_Penangkapan string    `json:"daerah_penangkapan" form:"daerah_penangkapan"`
	Total_Tangkapan    string    `json:"total_tangkapan" form:"total_tangkapan"`
	Tgl_Keberangkatan  time.Time `json:"tgl_keberangkatan" form:"tgl_keberangkatan"`
	Tgl_Kedatangan     time.Time `json:"tgl_kedatangan" form:"tgl_kedatangan"`
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}

func (KedatanganKapal) TableName() string {
	return "kedatangan_kapal"
}
