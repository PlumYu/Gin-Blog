package main

import (
	"Gin/Blog/controller"
	"Gin/Blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	categoryRoutes := r.Group("/categories")
	categorycontroller := controller.NewCategoryController()
	categoryRoutes.POST("", categorycontroller.Create)
	categoryRoutes.PUT("/:id", categorycontroller.Update)
	categoryRoutes.GET("/:id", categorycontroller.Show)
	categoryRoutes.DELETE("/:id", categorycontroller.Delete)
	return r
}
