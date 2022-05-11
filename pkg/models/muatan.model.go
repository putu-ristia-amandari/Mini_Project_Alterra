package models

import (
	"time"
)

type Muatan struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	Jenis_Muatan   string    `json:"jenis_muatan"`
	Status_Muatan  bool      `json:"status_muatan"`
	Note_Tangkapan string    `json:"note_tangkapan"`
	Created_At     time.Time `json:"created_at"`
	Updated_At     time.Time `json:"updated_at"`
}

func (Muatan) TableName() string {
	return "jenis_muatan"
}
