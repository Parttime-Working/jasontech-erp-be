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
			users.POST("/", controllers.CreateUser)
			users.GET("/", controllers.GetUsers)
			users.GET("/:id", controllers.GetUserByID)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
}
