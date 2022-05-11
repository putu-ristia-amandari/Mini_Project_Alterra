package models

import (
	"time"
)

type PerusahaanKapal struct {
	Id                   int       `gorm:"primaryKey" json:"id"`
	Nama_Perusahaan      string    `json:"nama_perusahaan"`
	Nama_PenanggungJawab string    `json:"nama_penanggungjawab"`
	Created_At           time.Time `json:"created_at"`
	Updated_At           time.Time `json:"updated_at"`
}

func (PerusahaanKapal) TableName() string {
	return "perusahaan"
}
