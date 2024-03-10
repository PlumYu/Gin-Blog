package controller

import (
	"Gin/Blog/model"
	"Gin/Blog/repository"
	"Gin/Blog/response"
	"Gin/Blog/vo"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(context *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := context.ShouldBind(&requestCategory); err != nil {
		response.Fail(context, nil, "数据验证错误, 分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(context, nil, "创建失败")
		return
	}

	response.Success(context, gin.H{"category": category}, "")
}

func (c CategoryController) Update(context *gin.Context) {
	// 绑定 body 中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := context.ShouldBind(&requestCategory); err != nil {
		response.Fail(context, nil, "数据验证错误, 分类名称必填")
		return
	}
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(context.Param("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(context, nil, "创建失败")
		return
	}
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(context, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(context *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(context.Param("id"))

	category, err := c.Repository.SelectById(categoryId)
	log.Println(category)
	if err != nil {
		response.Fail(context, nil, "创建失败")
		return
	}
	response.Success(context, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Delete(context *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(context.Param("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(context, nil, "删除失败，请重试")
		return
	}
	response.Success(context, nil, "")
}
