package models

import (
	"time"
)

type Muatan struct {
	Id           int       `gorm:"primaryKey" json:"id" form:"id"`
	Jenis_Muatan string    `json:"jenis_muatan" form:"jenis_muatan"`
	Created_At   time.Time `json:"created_at" form:"created_at"`
	Updated_At   time.Time `json:"updated_at" form:"updated_at"`
}

func (Muatan) TableName() string {
	return "jenis_muatan"
}
