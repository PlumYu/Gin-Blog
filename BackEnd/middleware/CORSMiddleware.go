package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") // 允许访问的地址
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")                      // 设置缓存时间
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")                    // 允许可以访问的方法
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")                    // Header
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")             // 允许访问的地址

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}
