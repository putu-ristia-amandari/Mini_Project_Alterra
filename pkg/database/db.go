package database

import (
	// "fmt"
	// "mini_project/pkg/config"
	"mini_project/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDBConnect() {

	dsn := "root:MieKuah.1123#@tcp(127.0.0.1:3306)/informasi_kedatangan_kapal_perikanan?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func DBConnect() *gorm.DB {
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&models.KedatanganKapal{})
}
