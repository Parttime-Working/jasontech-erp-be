package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission 權限模型
type Permission struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ModuleName     string         `gorm:"column:module_name;not null;size:100" json:"module_name"`
	Resource       string         `gorm:"not null;size:100" json:"resource"`
	Action         string         `gorm:"not null;size:100" json:"action"`
	Code           string         `gorm:"uniqueIndex;not null;size:500" json:"code"`
	DisplayName    string         `gorm:"column:display_name;not null;size:200" json:"display_name"`
	Description    string         `gorm:"size:255" json:"description"`
	Status         string         `gorm:"size:20;default:active" json:"status"`
	AutoRegistered bool           `gorm:"column:auto_registered;default:true" json:"auto_registered"`
	RegisteredAt   *time.Time     `json:"registered_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Roles          []Role         `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

// Role 角色模型
type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"uniqueIndex;not null;size:50" json:"name"`
	DisplayName string         `gorm:"not null;size:100" json:"display_name"`
	Description string         `gorm:"size:255" json:"description"`
	IsSystem    bool           `gorm:"default:false" json:"is_system"`
	Status      string         `gorm:"size:20;default:active" json:"status"`
	CreatedBy   *uint          `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Users       []User         `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// TableName 指定資料表名稱
func (Role) TableName() string {
	return "roles"
}

// TableName 指定資料表名稱
func (Permission) TableName() string {
	return "permissions"
}

// UserRole 使用者角色關聯
type UserRole struct {
	UserID uint `gorm:"primaryKey" json:"user_id"`
	RoleID uint `gorm:"primaryKey" json:"role_id"`
}

// TableName 指定資料表名稱
func (UserRole) TableName() string {
	return "user_roles"
}

// RolePermission 角色權限關聯
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey" json:"role_id"`
	PermissionID uint `gorm:"primaryKey" json:"permission_id"`
}

// TableName 指定資料表名稱
func (RolePermission) TableName() string {
	return "role_permissions"
}
