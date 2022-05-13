package models

import "time"

type User struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username"`
	password   string    `json:"-"`
	Role       string    `json:"role"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
