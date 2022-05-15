package models

type User struct {
	Id       int    `gorm:"primaryKey" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"-" form:"password"`
	// Role       string    `json:"role"`
	// Created_At time.Time `json:"created_at"`
	// Updated_At time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
