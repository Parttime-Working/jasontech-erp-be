package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"erp/routes"
)

func main() {
	// 創建 Gin 引擎
	r := gin.Default()

	// 初始化資料庫連接
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "無法連接到資料庫: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("成功連接到資料庫")

	// 設置路由
	routes.RegisterAuthRoutes(r)

	// 啟動伺服器，監聽 8000 端口
	r.Run(":8000")
}
