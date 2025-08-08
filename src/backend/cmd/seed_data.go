package main

import (
	"fmt"
	"os"

	"erp/db"
	"erp/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("開始使用 GORM 建立測試資料...")

	// 初始化資料庫連接
	database, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "無法連接到資料庫: %v\n", err)
		os.Exit(1)
	}
	defer database.Close()

	// 自動建立資料表 (如果不存在)
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "資料表建立失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ 資料表建立/更新成功")

	// 建立測試使用者
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "密碼加密失敗: %v\n", err)
		os.Exit(1)
	}

	user := models.User{
		Username: "admin",
		Email:    "admin@jasontech.com",
		Password: string(hashedPassword),
	}

	// 使用 GORM 的 FirstOrCreate 來避免重複建立
	result := database.Where(models.User{Username: user.Username}).FirstOrCreate(&user)
	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "建立測試使用者失敗: %v\n", err)
		os.Exit(1)
	}

	if result.RowsAffected > 0 {
		fmt.Println("✅ 測試使用者建立成功")
	} else {
		fmt.Println("✅ 測試使用者已存在")
	}

	fmt.Println("   使用者名稱: admin")
	fmt.Println("   密碼: password123")
	fmt.Println("   信箱: admin@jasontech.com")
	fmt.Printf("   使用者 ID: %d\n", user.ID)
}
