package models

import (
	"time"
)

type TipeKapal struct {
	Id              int       `gorm:"primaryKey" json:"id"`
	Id_Alat_Tangkap int       `gorm:"foreignKey: Alat_Tangkap_Id" json:"id_alat_tangkap"`
	Tipe_Kapal      string    `json:"tipe_kapal"`
	Created_At      time.Time `json:"created_at"`
	Updated_At      time.Time `json:"updated_at"`
}

func (TipeKapal) TableName() string {
	return "tipe_kapal"
}
