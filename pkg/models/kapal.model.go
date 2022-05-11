package models

import (
	"time"
)

type Kapal struct {
	Id              int       `gorm:"primaryKey" json:"id"`
	Id_Perusahaan   int       `gorm:"foreignKey: perusahaanId" json:"id_perusahaan"`
	Id_Pelabuhan    int       `gorm:"foreignKey: pelabuhanId" json:"id_pelabuhan"`
	Id_tipe_kapal   int       `gorm:"foreignKey: jenis_muatanID" json:"id_tipe_kapal"`
	Nama_Kapal      string    `json:"nama_kapal"`
	Ukuran_Kapal    string    `json:"ukuran_kapal"`
	Kekuatan_Mesin  string    `json:"kekuatan_mesin"`
	Call_Sign_Kapal string    `json:"call_sign_kapal"`
	Nama_Nahkoda    string    `json:"nama_nahkoda"`
	Created_At      time.Time `json:"created_at"`
	Updated_At      time.Time `json:"updated_at"`
}

func (Kapal) TableName() string {
	return "kapal"
}
