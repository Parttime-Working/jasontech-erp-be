package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 創建 Gin 引擎
	r := gin.Default()

	// 設置簡單的 GET 路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 模擬一個 API 路由
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 啟動伺服器，監聽 8000 端口
	r.Run(":8000")
}
