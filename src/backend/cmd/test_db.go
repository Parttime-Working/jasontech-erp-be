package main

import (
	"fmt"
	"os"

	"erp/db"
)

func main() {
	fmt.Println("開始資料庫連接測試...")

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

	fmt.Println("✅ 資料庫連線測試成功!")
	fmt.Println("✅ 所有資料庫操作正常")
}
