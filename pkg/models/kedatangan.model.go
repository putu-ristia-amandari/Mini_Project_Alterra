package models

import (
	"time"
)

type KedatanganKapal struct {
	Id                 int       `gorm:"primaryKey" json:"id"`
	Id_Kapal           int       `gorm:"foreignKey: kapalID" json:"id_kapal"`
	Id_Pelabuhan       int       `gorm:"foreignKey: pelabuhanID" json:"id_pelabuhan"`
	Id_Jenis_Muatan    int       `gorm:"foreignKey: jenis_muatanID" json:"id_jenis_muatan"`
	Daerah_Penangkapan string    `json:"daerah_penangkapan"`
	Total_Tangkapan    string    `json:"total_tangkapan"`
	Tgl_Keberangkatan  time.Time `json:"tgl_keberangkatan"`
	Tgl_Kedatangan     time.Time `json:"tgl_kedatangan"`
	Created_At         time.Time `json:"created_at"`
	Updated_At         time.Time `json:"updated_at"`
}

func (KedatanganKapal) TableName() string {
	return "kedatangan_kapal"
}
