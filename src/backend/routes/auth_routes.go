package routes

import (
	"erp/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)

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
