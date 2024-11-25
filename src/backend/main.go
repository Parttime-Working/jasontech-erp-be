package main

import (
	"github.com/gin-gonic/gin"
	"erp/routes"
)

func main() {
	// 創建 Gin 引擎
	r := gin.Default()

	// 設置路由
	routes.RegisterAuthRoutes(r)

	// 啟動伺服器，監聽 8000 端口
	r.Run(":8000")
}
