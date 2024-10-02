package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// 将username改为自己的MySQL用户名,passwd改为自己的密码,yourdb改为你使用的库
	dsn := "username:passwd@/yourdb?charset=utf8&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = connection.Debug()
	connection.AutoMigrate(&models.User{})
}
