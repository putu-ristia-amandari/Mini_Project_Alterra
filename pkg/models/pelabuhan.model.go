package models

import (
	"time"
)

type Pelabuhan struct {
	Id             int       `gorm:"primaryKey" json:"id" form:"id"`
	Nama_Pelabuhan string    `json:"nama_pelabuhan" form:"nama_pelabuhan"`
	Created_At     time.Time `json:"created_at" form:"created_at"`
	Updated_At     time.Time `json:"updated_at" form:"updated_at"`
}

func (Pelabuhan) TableName() string {
	return "pelabuhan"
}
