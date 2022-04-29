package database

import (
	"database/sql"
	"fmt"

	"github.com/putu-ristia-amandari/Mini_Project_Alterra/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
	}
}

