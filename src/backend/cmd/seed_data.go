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
	err = database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.UserRole{}, &models.RolePermission{})
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
		Level:    "super_admin", // 明確設置為超級管理員等級
	}

	// 使用 GORM 的 FirstOrCreate 來避免重複建立
	result := database.Where(models.User{Username: user.Username}).Assign(models.User{
		Email:    user.Email,
		Password: user.Password,
		Level:    user.Level, // 確保等級始終被更新
	}).FirstOrCreate(&user)
	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "建立測試使用者失敗: %v\n", err)
		os.Exit(1)
	}

	if result.RowsAffected > 0 {
		fmt.Println("✅ 管理員使用者建立成功")
		fmt.Println("   📧 信箱: admin@jasontech.com")
		fmt.Println("   👤 使用者名: admin")
		fmt.Println("   🔑 密碼: password123")
		fmt.Println("   👑 等級: super_admin")
		fmt.Printf("   🆔 使用者 ID: %d\n", user.ID)
		fmt.Println("")
		fmt.Println("🎯 系統初始化完成！")
		fmt.Println("   • 超級管理員帳號已建立")
		fmt.Println("   • 新使用者默認等級為 'user'")
		fmt.Println("   • 只有管理員或超級管理員可以創建其他使用者")
	} else {
		fmt.Println("✅ 管理員使用者已存在")
		fmt.Printf("   🆔 使用者 ID: %d\n", user.ID)
	}

	// 建立測試用的 sample user
	samplePassword := "sample123"
	sampleHashedPassword, err := bcrypt.GenerateFromPassword([]byte(samplePassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sample 使用者密碼加密失敗: %v\n", err)
		os.Exit(1)
	}

	sampleUser := models.User{
		Username: "sampleuser",
		Email:    "sample@jasontech.com",
		Password: string(sampleHashedPassword),
		Level:    "user", // 一般使用者等級
	}

	// 使用 GORM 的 FirstOrCreate 來避免重複建立
	sampleResult := database.Where(models.User{Username: sampleUser.Username}).FirstOrCreate(&sampleUser)
	if sampleResult.Error != nil {
		fmt.Fprintf(os.Stderr, "建立 sample 使用者失敗: %v\n", sampleResult.Error)
		os.Exit(1)
	}

	if sampleResult.RowsAffected > 0 {
		fmt.Println("✅ Sample 使用者建立成功")
		fmt.Println("   📧 信箱: sample@jasontech.com")
		fmt.Println("   👤 使用者名: sampleuser")
		fmt.Println("   🔑 密碼: sample123")
		fmt.Println("   👤 等級: user")
		fmt.Printf("   🆔 使用者 ID: %d\n", sampleUser.ID)
	} else {
		fmt.Println("✅ Sample 使用者已存在")
		fmt.Printf("   🆔 使用者 ID: %d\n", sampleUser.ID)
	}

	// 建立範例角色
	roles := []models.Role{
		{Name: "hr_manager", DisplayName: "人力資源經理", Description: "人力資源經理，負責人力資源管理"},
		{Name: "hr_specialist", DisplayName: "人力資源專員", Description: "人力資源專員，負責人事行政工作"},
		{Name: "employee", DisplayName: "一般員工", Description: "一般員工"},
		{Name: "finance_manager", DisplayName: "財務經理", Description: "財務經理，負責財務管理"},
		{Name: "finance_specialist", DisplayName: "財務專員", Description: "財務專員，負責財務行政工作"},
	}

	for _, role := range roles {
		result := database.Where(models.Role{Name: role.Name}).FirstOrCreate(&role)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "建立角色失敗: %v\n", result.Error)
			os.Exit(1)
		}
		if result.RowsAffected > 0 {
			fmt.Printf("✅ 角色 '%s' 建立成功\n", role.Name)
		}
	}

	// 建立範例權限
	permissions := []models.Permission{
		// 人力資源模組
		{ModuleName: "hr", Resource: "employees", Action: "view", Code: "hr.employees.view", DisplayName: "查看員工資訊", Description: "查看員工基本資訊和聯絡方式"},
		{ModuleName: "hr", Resource: "employees", Action: "create", Code: "hr.employees.create", DisplayName: "創建員工記錄", Description: "創建新的員工記錄"},
		{ModuleName: "hr", Resource: "employees", Action: "edit", Code: "hr.employees.edit", DisplayName: "編輯員工資訊", Description: "編輯員工的基本資訊"},
		{ModuleName: "hr", Resource: "employees", Action: "delete", Code: "hr.employees.delete", DisplayName: "刪除員工記錄", Description: "刪除員工記錄"},
		{ModuleName: "hr", Resource: "attendance", Action: "manage", Code: "hr.attendance.manage", DisplayName: "管理出勤記錄", Description: "管理員工的出勤記錄"},
		{ModuleName: "hr", Resource: "reports", Action: "view", Code: "hr.reports.view", DisplayName: "查看人事報表", Description: "查看各種人事相關的報表"},
		{ModuleName: "hr", Resource: "documents", Action: "manage", Code: "hr.documents.manage", DisplayName: "管理員工文件", Description: "管理員工的相關文件"},

		// 財務模組
		{ModuleName: "finance", Resource: "financials", Action: "view", Code: "finance.financials.view", DisplayName: "查看財務資訊", Description: "查看財務相關資訊"},
		{ModuleName: "finance", Resource: "payroll", Action: "manage", Code: "finance.payroll.manage", DisplayName: "管理薪資", Description: "管理員工薪資"},
		{ModuleName: "finance", Resource: "timesheets", Action: "manage", Code: "finance.timesheets.manage", DisplayName: "管理工時記錄", Description: "管理員工工時記錄"},
		{ModuleName: "finance", Resource: "purchase_requests", Action: "create", Code: "finance.purchase_requests.create", DisplayName: "創建請購單", Description: "創建請購單"},
		{ModuleName: "finance", Resource: "purchase_requests", Action: "approve", Code: "finance.purchase_requests.approve", DisplayName: "審核請購單", Description: "審核請購單"},
		{ModuleName: "finance", Resource: "payment_requests", Action: "create", Code: "finance.payment_requests.create", DisplayName: "創建請款單", Description: "創建請款單"},
		{ModuleName: "finance", Resource: "payment_requests", Action: "approve", Code: "finance.payment_requests.approve", DisplayName: "審核請款單", Description: "審核請款單"},
		{ModuleName: "finance", Resource: "budget", Action: "view", Code: "finance.budget.view", DisplayName: "查看預算", Description: "查看預算資訊"},
		{ModuleName: "finance", Resource: "budget", Action: "manage", Code: "finance.budget.manage", DisplayName: "管理預算", Description: "管理預算"},
		{ModuleName: "finance", Resource: "expense_reports", Action: "view", Code: "finance.expense_reports.view", DisplayName: "查看費用報表", Description: "查看費用報表"},
		{ModuleName: "finance", Resource: "vendors", Action: "manage", Code: "finance.vendors.manage", DisplayName: "管理供應商", Description: "管理供應商資訊"},

		// 專案管理模組 (服務型公司核心)
		{ModuleName: "project", Resource: "projects", Action: "view", Code: "project.projects.view", DisplayName: "查看專案資訊", Description: "查看專案基本資訊"},
		{ModuleName: "project", Resource: "projects", Action: "create", Code: "project.projects.create", DisplayName: "創建專案", Description: "創建新專案"},
		{ModuleName: "project", Resource: "projects", Action: "edit", Code: "project.projects.edit", DisplayName: "編輯專案資訊", Description: "編輯專案基本資訊"},
		{ModuleName: "project", Resource: "projects", Action: "manage_team", Code: "project.projects.manage_team", DisplayName: "管理專案團隊", Description: "管理專案團隊成員"},
		{ModuleName: "project", Resource: "reports", Action: "view", Code: "project.reports.view", DisplayName: "查看專案報表", Description: "查看專案相關報表"},
		{ModuleName: "project", Resource: "budget", Action: "manage", Code: "project.budget.manage", DisplayName: "管理專案預算", Description: "管理專案預算"},

		// 系統管理模組
		{ModuleName: "system", Resource: "users", Action: "manage", Code: "system.users.manage", DisplayName: "管理系統使用者", Description: "管理系統使用者帳號"},
		{ModuleName: "system", Resource: "roles", Action: "manage", Code: "system.roles.manage", DisplayName: "管理角色和權限", Description: "管理角色和權限設定"},
		{ModuleName: "system", Resource: "logs", Action: "view", Code: "system.logs.view", DisplayName: "查看系統日誌", Description: "查看系統操作日誌"},
		{ModuleName: "system", Resource: "settings", Action: "manage", Code: "system.settings.manage", DisplayName: "管理系統設定", Description: "管理系統配置設定"},
	}

	for _, perm := range permissions {
		result := database.Where(models.Permission{Code: perm.Code}).FirstOrCreate(&perm)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "建立權限失敗: %v\n", result.Error)
			os.Exit(1)
		}
		if result.RowsAffected > 0 {
			fmt.Printf("✅ 權限 '%s' 建立成功\n", perm.DisplayName)
		}
	}

	fmt.Println("")
	fmt.Println("🎯 系統初始化完成！")
	fmt.Println("   • 超級管理員帳號已建立")
	fmt.Println("   • Sample 使用者已建立")
	fmt.Println("   • 新使用者默認等級為 'user'")
	fmt.Println("   • 只有管理員或超級管理員可以創建其他使用者")
}
