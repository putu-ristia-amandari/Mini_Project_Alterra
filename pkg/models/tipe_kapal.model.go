package models

import (
	"time"
)

type TipeKapal struct {
	Id         int       `gorm:"primaryKey" json:"id" form:"id"`
	Tipe_Kapal string    `json:"tipe_kapal" form:"tipe_kapal"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
}

func (TipeKapal) TableName() string {
	return "tipe_kapal"
}
