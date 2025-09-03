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
		Role:     "admin", // 明確設置為管理員角色
	}

	// 使用 GORM 的 FirstOrCreate 來避免重複建立
	result := database.Where(models.User{Username: user.Username}).FirstOrCreate(&user)
	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "建立測試使用者失敗: %v\n", err)
		os.Exit(1)
	}

	if result.RowsAffected > 0 {
		fmt.Println("✅ 管理員用戶建立成功")
		fmt.Println("   📧 信箱: admin@jasontech.com")
		fmt.Println("   👤 用戶名: admin")
		fmt.Println("   🔑 密碼: password123")
		fmt.Println("   👑 角色: admin")
		fmt.Printf("   🆔 用戶 ID: %d\n", user.ID)
		fmt.Println("")
		fmt.Println("🎯 系統初始化完成！")
		fmt.Println("   • 管理員帳號已建立")
		fmt.Println("   • 新用戶默認角色為 'user'")
		fmt.Println("   • 只有管理員可以創建其他用戶")
	} else {
		fmt.Println("✅ 管理員用戶已存在")
		fmt.Printf("   🆔 用戶 ID: %d\n", user.ID)
	}

	// 建立測試用的 sample user
	samplePassword := "sample123"
	sampleHashedPassword, err := bcrypt.GenerateFromPassword([]byte(samplePassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sample 用戶密碼加密失敗: %v\n", err)
		os.Exit(1)
	}

	sampleUser := models.User{
		Username: "sampleuser",
		Email:    "sample@jasontech.com",
		Password: string(sampleHashedPassword),
		Role:     "user", // 一般用戶角色
	}

	// 使用 GORM 的 FirstOrCreate 來避免重複建立
	sampleResult := database.Where(models.User{Username: sampleUser.Username}).FirstOrCreate(&sampleUser)
	if sampleResult.Error != nil {
		fmt.Fprintf(os.Stderr, "建立 sample 用戶失敗: %v\n", sampleResult.Error)
		os.Exit(1)
	}

	if sampleResult.RowsAffected > 0 {
		fmt.Println("✅ Sample 用戶建立成功")
		fmt.Println("   📧 信箱: sample@jasontech.com")
		fmt.Println("   👤 用戶名: sampleuser")
		fmt.Println("   🔑 密碼: sample123")
		fmt.Println("   👤 角色: user")
		fmt.Printf("   🆔 用戶 ID: %d\n", sampleUser.ID)
	} else {
		fmt.Println("✅ Sample 用戶已存在")
		fmt.Printf("   🆔 用戶 ID: %d\n", sampleUser.ID)
	}

	fmt.Println("")
	fmt.Println("🎯 系統初始化完成！")
	fmt.Println("   • 管理員帳號已建立")
	fmt.Println("   • Sample 用戶已建立")
	fmt.Println("   • 新用戶默認角色為 'user'")
	fmt.Println("   • 只有管理員可以創建其他用戶")
}
