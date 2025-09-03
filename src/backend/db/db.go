package db

import (
	"fmt"
	"os"

	"erp/models"
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

// === Repository 介面定義 ===

// UserRepository 使用者資料存取介面
type UserRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

// === Repository 實作 ===

// userRepository 使用者資料存取實作
type userRepository struct {
	db *DB
}

// NewUserRepository 建立使用者 repository
func NewUserRepository(db *DB) UserRepository {
	return &userRepository{db: db}
}

// Create 建立使用者
func (r *userRepository) Create(user *models.User) error {
	return r.db.DB.Create(user).Error
}

// GetByUsername 根據用戶名獲取使用者
func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID 根據 ID 獲取使用者
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll 獲取所有使用者
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.DB.Find(&users).Error
	return users, err
}

// Update 更新使用者
func (r *userRepository) Update(user *models.User) error {
	return r.db.DB.Save(user).Error
}

// Delete 刪除使用者
func (r *userRepository) Delete(id uint) error {
	return r.db.DB.Delete(&models.User{}, id).Error
}
