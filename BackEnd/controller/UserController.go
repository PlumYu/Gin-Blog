package controller

import (
	"Gin/Blog/common"
	"Gin/Blog/model"
	"Gin/Blog/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为 11 位"})
		return
	}
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于 6 位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户已存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	context.JSON(200, gin.H{
		"meg": "success",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
