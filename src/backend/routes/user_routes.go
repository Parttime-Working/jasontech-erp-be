package routes

import (
	"erp/controllers"
	"erp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			// 創建用戶需要管理員權限
			users.POST("/", middleware.AdminMiddleware(), controllers.CreateUser)
			users.GET("/", controllers.GetUsers)
			users.GET("/:id", controllers.GetUserByID)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
}
