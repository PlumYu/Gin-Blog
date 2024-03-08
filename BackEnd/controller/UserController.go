package controller

import (
	"Gin/Blog/common"
	"Gin/Blog/dto"
	"Gin/Blog/model"
	"Gin/Blog/response"
	"Gin/Blog/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	// 获取参数  获取不到~
	//// 使用 map 获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(context.Request.Body).Decode(&requestMap)
	// 使用结构体来获取参数
	var requestUser = model.User{}
	//json.NewDecoder(context.Request.Body).Decode(&requestUser)
	// gin 框架提供的 bind 参数
	context.Bind(&requestUser)
	// 获取参数
	DB := common.GetDB()
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为 11 位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于 6 位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "该用户已存在")
		return
	}
	// 创建用户 用户密码进行加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 发放 token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	// 返回结果
	response.Success(context, gin.H{"token": token}, "登录成功")
}

func Login(context *gin.Context) {
	// 获取参数  获取不到~
	//// 使用 map 获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(context.Request.Body).Decode(&requestMap)
	// 使用结构体来获取参数
	var requestUser = model.User{}
	//json.NewDecoder(context.Request.Body).Decode(&requestUser)
	// gin 框架提供的 bind 参数
	context.Bind(&requestUser)
	// 获取参数
	DB := common.GetDB()
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	log.Println(name, telephone, password)
	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为 11 位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于 6 位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "User not found")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	// 返回结果
	response.Success(context, gin.H{"token": token}, "登录成功")
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
