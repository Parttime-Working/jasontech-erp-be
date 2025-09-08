package routes

import (
	"erp/controllers"
	"erp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		permissions := api.Group("/permissions")
		permissions.Use(middleware.AuthMiddleware())
		{
			// 權限管理需要管理員權限
			permissions.POST("/", middleware.AdminMiddleware(), controllers.CreatePermission)
			permissions.GET("/", controllers.GetPermissions)
			permissions.GET("/:id", controllers.GetPermissionByID)
			permissions.PUT("/:id", middleware.AdminMiddleware(), controllers.UpdatePermission)
			permissions.DELETE("/:id", middleware.AdminMiddleware(), controllers.DeletePermission)

			// 角色權限分配 - 使用不同的路徑結構避免參數衝突
			permissions.POST("/:id/roles/:roleId", middleware.AdminMiddleware(), controllers.AssignPermissionToRole)
			permissions.DELETE("/:id/roles/:roleId", middleware.AdminMiddleware(), controllers.RemovePermissionFromRole)
		}
	}
}
