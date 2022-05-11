package models

import (
	"time"
)

type AlatTangkap struct {
	Id                int       `gorm:"primaryKey" json:"id"`
	Nama_Alat_Tangkap string    `json:"nama_alat_tangkap"`
	Created_At        time.Time `json:"created_at"`
	Updated_At        time.Time `json:"updated_at"`
}

func (AlatTangkap) TableName() string {
	return "alat_tangkap"
}
