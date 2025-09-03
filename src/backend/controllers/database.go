package controllers

import (
	"erp/db"
)

// 全局資料庫實例 (依賴注入)
var database *db.DB

// Repository 實例
var userRepo db.UserRepository

// SetDB 設定資料庫依賴 (依賴注入)
func SetDB(dbInstance *db.DB) {
	database = dbInstance
	userRepo = db.NewUserRepository(dbInstance)
}

// GetUserRepo 獲取使用者 repository
func GetUserRepo() db.UserRepository {
	return userRepo
}
