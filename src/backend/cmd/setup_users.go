package main

import (
	"context"
	"fmt"
	"os"

	"erp/db"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("開始建立使用者資料表和測試資料...")

	// 初始化資料庫連接
	database, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "無法連接到資料庫: %v\n", err)
		os.Exit(1)
	}
	defer database.Close()

	conn := database.GetConn()

	// 建立使用者資料表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "建立資料表失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 使用者資料表建立成功")

	// 建立測試使用者
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "密碼加密失敗: %v\n", err)
		os.Exit(1)
	}

	insertUserSQL := `
	INSERT INTO users (username, email, password) 
	VALUES ($1, $2, $3) 
	ON CONFLICT (username) DO NOTHING;`

	_, err = conn.Exec(context.Background(), insertUserSQL, "admin", "admin@jasontech.com", string(hashedPassword))
	if err != nil {
		fmt.Fprintf(os.Stderr, "建立測試使用者失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 測試使用者建立成功")
	fmt.Println("   使用者名稱: admin")
	fmt.Println("   密碼: password123")
	fmt.Println("   信箱: admin@jasontech.com")
}
