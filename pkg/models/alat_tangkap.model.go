package models

import (
	"time"
)

type AlatTangkap struct {
	Id                int       `gorm:"primaryKey" json:"id" form:"id"`
	Id_Tipe_Kapal     int       `json:"id_tipe_kapal" form:"id_tipe_kapal"`
	Nama_Alat_Tangkap string    `json:"nama_alat_tangkap" form:"nama_alat_tangkap"`
	Created_At        time.Time `json:"created_at" form:"created_at"`
	Updated_At        time.Time `json:"updated_at" form:"updated_at"`
}

func (AlatTangkap) TableName() string {
	return "alat_tangkap"
}
