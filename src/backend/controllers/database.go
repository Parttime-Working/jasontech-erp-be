package controllers

import "erp/db"

// 全局資料庫實例 (依賴注入)
var database *db.DB

// SetDB 設定資料庫依賴 (依賴注入)
func SetDB(db *db.DB) {
	database = db
}
