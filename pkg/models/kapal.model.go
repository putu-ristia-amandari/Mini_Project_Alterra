package models

import (
	"time"
)

type Kapal struct {
	Id              int       `gorm:"primaryKey" json:"id" form:"id"`
	Id_Perusahaan   int       `json:"id_perusahaan" form:"id_perusahaan"`
	Id_Pelabuhan    int       `json:"id_pelabuhan" form:"id_pelabuhan"`
	Id_Tipe_Kapal   int       `json:"id_tipe_kapal" form:"id_tipe_kapal"`
	Nama_Kapal      string    `json:"nama_kapal" form:"nama_kapal"`
	Ukuran_Kapal    string    `json:"ukuran_kapal" form:"ukuran_kapal"`
	Kekuatan_Mesin  string    `json:"kekuatan_mesin" form:"kekuatan_mesin"`
	Call_Sign_Kapal string    `json:"call_sign_kapal" form:"call_sign_kapal"`
	Nama_Nahkoda    string    `json:"nama_nahkoda" form:"nama_nahkoda"`
	Created_At      time.Time `json:"created_at" form:"created_at"`
	Updated_At      time.Time `json:"updated_at" form:"updated_at"`
}

func (Kapal) TableName() string {
	return "kapal"
}
