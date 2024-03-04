package main

import (
	"Gin/Blog/common"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	router := gin.Default()
	router = CollectRoute(router)
	port := viper.GetString("server.port")
	if port != "" {
		panic(router.Run(":" + port))
	}
	panic(router.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
