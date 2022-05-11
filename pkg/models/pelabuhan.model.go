package models

import (
	"time"
)

type Pelabuhan struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	Nama_Pelabuhan string    `json:"nama_pelabuhan"`
	Created_At     time.Time `json:"created_at"`
	Updated_At     time.Time `json:"updated_at"`
}

func (Pelabuhan) TableName() string {
	return "pelabuhan"
}
