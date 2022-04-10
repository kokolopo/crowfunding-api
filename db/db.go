package db

import (
	"crowfunding_api/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

func InitDB() {
	conf := config.GetConfig()

	dsn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("berhasil akses db!!!")
	}
}
