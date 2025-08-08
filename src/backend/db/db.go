package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 包裝 GORM 資料庫連接
type DB struct {
	*gorm.DB
}

// New 建立新的 GORM 資料庫連接
func New() (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 顯示 SQL 查詢日誌
	})
	if err != nil {
		return nil, fmt.Errorf("無法連接到資料庫: %v", err)
	}

	return &DB{db}, nil
}

// Close 關閉資料庫連接
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// TestConnection 測試資料庫連接
func (db *DB) TestConnection() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("無法取得 SQL DB 實例: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("資料庫連接測試失敗: %v", err)
	}

	return nil
}

// AutoMigrate 自動建立或更新資料表 schema
func (db *DB) AutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models...)
}
