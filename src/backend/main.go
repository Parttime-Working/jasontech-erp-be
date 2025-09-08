package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"erp/db"
	"erp/controllers"
	"erp/models"
	"erp/routes"
)

func main() {
	// 創建 Gin 引擎
	r := gin.Default()

	// 啟用自動重定向 - 統一處理斜線問題
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	// 設定 CORS 中間件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 初始化資料庫連接
	database, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "無法連接到資料庫: %v\n", err)
		os.Exit(1)
	}
	defer database.Close()

	// 測試資料庫連線
	err = database.TestConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "資料庫連線測試失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("資料庫連線測試成功")

	// 自動建立/更新資料表 schema
	err = database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.UserRole{}, &models.RolePermission{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "資料庫 migration 失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("資料庫 schema 同步完成")
	fmt.Println("成功連接到資料庫")

	// 初始化 controllers 並注入資料庫依賴
	controllers.SetDB(database)

	// 設置路由
	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterRoleRoutes(r)
	routes.RegisterPermissionRoutes(r)

	// 啟動伺服器，監聽 8000 端口
	r.Run(":8000")
}
