package models

import (
	"time"
	"gorm.io/gorm"
)

// User 使用者模型 (GORM)
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email       string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password    string         `gorm:"not null;size:255" json:"-"` // 隱藏密碼欄位
	Level       string         `gorm:"default:user;size:20" json:"level"` // 等級：user, admin, super_admin
	LastLoginAt *time.Time     `json:"last_login_at"` // 最後登入時間
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定資料表名稱
func (User) TableName() string {
	return "users"
}

// Credentials 登入憑證
type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserInput 建立使用者時的輸入
type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Level    string `json:"level,omitempty"` // 可選，默認為 "user"
}

// UpdateUserInput 更新使用者時的輸入
type UpdateUserInput struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Level    *string `json:"level,omitempty"`
}

// UserResponse 回傳給前端的使用者資訊 (不包含密碼)
type UserResponse struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Level       string     `json:"level"`
	LastLoginAt *time.Time `json:"last_login_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
