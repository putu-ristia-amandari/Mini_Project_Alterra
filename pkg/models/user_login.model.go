package models

type User struct {
	Id       int    `gorm:"primaryKey" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"-" form:"password"`
}

func (User) TableName() string {
	return "users"
}
