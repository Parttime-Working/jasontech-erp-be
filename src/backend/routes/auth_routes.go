package routes

import (
	"erp/controllers"
	"erp/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)

		// 添加驗證端點，使用標準的身份驗證中間件
		auth.GET("/verify", middleware.AuthMiddleware(), controllers.VerifyToken)

		// 新增測試路由
		auth.POST("/test", func(c *gin.Context) {
			var body map[string]interface{}
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"received": body,
			})
		})
	}
}
