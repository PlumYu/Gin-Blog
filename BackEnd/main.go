package main

import (
	"Gin/Blog/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	router := gin.Default()
	router = CollectRoute(router)
	panic(router.Run())
}
