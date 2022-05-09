package database

import (
	"fmt"
	"mini_project/pkg/config"
	"mini_project/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDBConnect() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_PORT,
		conf.DB_HOST,
		conf.DB_NAME,
	)

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
