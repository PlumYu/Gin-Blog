package common

import (
	"Gin/Blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasources.driverName")
	host := viper.GetString("datasources.host")
	port := viper.GetString("datasources.port")
	database := viper.GetString("datasources.database")
	username := viper.GetString("datasources.username")
	password := viper.GetString("datasources.password")
	charset := viper.GetString("datasources.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
