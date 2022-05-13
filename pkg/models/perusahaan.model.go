package models

import (
	"time"
)

type PerusahaanKapal struct {
	Id                   int       `gorm:"primaryKey" json:"id" form:"id"`
	Nama_Perusahaan      string    `json:"nama_perusahaan" form:"nama_perusahaan"`
	Nama_PenanggungJawab string    `json:"nama_penanggungjawab" form:"nama_penanggungjawab"`
	Created_At           time.Time `json:"created_at" form:"created_at"`
	Updated_At           time.Time `json:"updated_at" form:"created_at"`
}

func (PerusahaanKapal) TableName() string {
	return "perusahaan"
}
