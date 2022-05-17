package database

import (
	"mini_project/pkg/config"
	"mini_project/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDBConnect() {
	conf := config.GetConfig()

	dsn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"

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
	DB.AutoMigrate(&models.Kapal{})
	DB.AutoMigrate(&models.Pelabuhan{})
	DB.AutoMigrate(&models.PerusahaanKapal{})
	DB.AutoMigrate(&models.Muatan{})
	DB.AutoMigrate(&models.AlatTangkap{})
	DB.AutoMigrate(&models.TipeKapal{})
	DB.AutoMigrate(&models.User{})
}
