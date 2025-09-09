package routes

import (
	"erp/controllers"
	"erp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.RouterGroup) {
	roles := r.Group("/roles")
	roles.Use(middleware.AuthMiddleware())
	{
		// 角色管理需要管理員權限
		roles.POST("/", middleware.AdminMiddleware(), controllers.CreateRole)
		roles.GET("/", controllers.GetRoles)
		roles.GET("/:id", controllers.GetRoleByID)
		roles.PUT("/:id", middleware.AdminMiddleware(), controllers.UpdateRole)
		roles.DELETE("/:id", middleware.AdminMiddleware(), controllers.DeleteRole)

		// 使用者角色分配 - 使用不同的路徑結構避免參數衝突
		roles.POST("/:id/users/:userId", middleware.AdminMiddleware(), controllers.AssignRoleToUser)
		roles.DELETE("/:id/users/:userId", middleware.AdminMiddleware(), controllers.RemoveRoleFromUser)
	}
}
