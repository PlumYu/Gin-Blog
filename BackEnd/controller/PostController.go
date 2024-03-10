package controller

import (
	"Gin/Blog/common"
	"Gin/Blog/model"
	"Gin/Blog/response"
	"Gin/Blog/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

type IPostController interface {
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(context *gin.Context) {
	var requestPost vo.CreatePostRequst
	// 数据验证
	if err := context.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(context, nil, "数据验证错误")
		return
	}

	// 获取登录用户 user
	user, _ := context.Get("user")
	// 创建 post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}

	response.Success(context, nil, "创建成功")
}

func (p PostController) Update(context *gin.Context) {
	var requestPost vo.CreatePostRequst
	// 数据验证
	if err := context.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(context, nil, "数据验证错误")
		return
	}

	// 获取 path 中的 id
	postId := context.Param("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context, nil, "数据验证错误")
		return
	}

	// 当前文章是否为文章的作者

	// 获取登录用户 user
	user, _ := context.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(context, nil, "文章不属于您， 请勿进行非法操作")
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(context, nil, "文章不属于您， 请勿进行非法操作")
		return
	}
	response.Success(context, gin.H{"psot": post}, "更新成功")
}

func (p PostController) Show(context *gin.Context) {
	// 获取 path 中的 id
	postId := context.Param("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context, nil, "文章不存在")
		return
	}

	response.Success(context, gin.H{"psot": post}, "更新成功")
}

func (p PostController) Delete(context *gin.Context) {
	// 获取 path 中的 id
	postId := context.Param("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context, nil, "文章不存在")
		return
	}
	// 当前文章是否为文章的作者

	// 获取登录用户 user
	user, _ := context.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(context, nil, "文章不属于您， 请勿进行非法操作")
		return
	}

	p.DB.Delete(&post)
	response.Success(context, gin.H{"psot": post}, "删除成功")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}
