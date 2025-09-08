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

// RoleRepository 角色資料存取介面
type RoleRepository interface {
	Create(role *models.Role) error
	GetByID(id uint) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetAll() ([]models.Role, error)
	Update(role *models.Role) error
	Delete(id uint) error
}

// PermissionRepository 權限資料存取介面
type PermissionRepository interface {
	Create(permission *models.Permission) error
	GetByID(id uint) (*models.Permission, error)
	GetByCode(code string) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
	GetByModule(module string) ([]models.Permission, error)
	Update(permission *models.Permission) error
	Delete(id uint) error
}

// RolePermissionRepository 角色權限關聯資料存取介面
type RolePermissionRepository interface {
	Create(rolePermission *models.RolePermission) error
	Delete(roleID, permissionID uint) error
	GetPermissionsByRoleID(roleID uint) ([]models.Permission, error)
	GetRolesByPermissionID(permissionID uint) ([]models.Role, error)
	Exists(roleID, permissionID uint) (bool, error)
}

// UserRoleRepository 使用者角色關聯資料存取介面
type UserRoleRepository interface {
	Create(userRole *models.UserRole) error
	Delete(userID, roleID uint) error
	GetRolesByUserID(userID uint) ([]models.Role, error)
	GetUsersByRoleID(roleID uint) ([]models.User, error)
	Exists(userID, roleID uint) (bool, error)
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

// === Role Repository 實作 ===

// roleRepository 角色資料存取實作
type roleRepository struct {
	db *DB
}

// NewRoleRepository 建立角色 repository
func NewRoleRepository(db *DB) RoleRepository {
	return &roleRepository{db: db}
}

// Create 建立角色
func (r *roleRepository) Create(role *models.Role) error {
	return r.db.DB.Create(role).Error
}

// GetByID 根據 ID 獲取角色
func (r *roleRepository) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.DB.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetByName 根據名稱獲取角色
func (r *roleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.DB.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAll 獲取所有角色
func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.DB.Find(&roles).Error
	return roles, err
}

// Update 更新角色
func (r *roleRepository) Update(role *models.Role) error {
	return r.db.DB.Save(role).Error
}

// Delete 刪除角色
func (r *roleRepository) Delete(id uint) error {
	return r.db.DB.Delete(&models.Role{}, id).Error
}

// === Permission Repository 實作 ===

// permissionRepository 權限資料存取實作
type permissionRepository struct {
	db *DB
}

// NewPermissionRepository 建立權限 repository
func NewPermissionRepository(db *DB) PermissionRepository {
	return &permissionRepository{db: db}
}

// Create 建立權限
func (r *permissionRepository) Create(permission *models.Permission) error {
	return r.db.DB.Create(permission).Error
}

// GetByID 根據 ID 獲取權限
func (r *permissionRepository) GetByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.DB.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetByCode 根據代碼獲取權限
func (r *permissionRepository) GetByCode(code string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.DB.Where("code = ?", code).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetAll 獲取所有權限
func (r *permissionRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.DB.Find(&permissions).Error
	return permissions, err
}

// GetByModule 根據模組獲取權限
func (r *permissionRepository) GetByModule(module string) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.DB.Where("module_name = ?", module).Find(&permissions).Error
	return permissions, err
}

// Update 更新權限
func (r *permissionRepository) Update(permission *models.Permission) error {
	return r.db.DB.Save(permission).Error
}

// Delete 刪除權限
func (r *permissionRepository) Delete(id uint) error {
	return r.db.DB.Delete(&models.Permission{}, id).Error
}

// === RolePermission Repository 實作 ===

// rolePermissionRepository 角色權限關聯資料存取實作
type rolePermissionRepository struct {
	db *DB
}

// NewRolePermissionRepository 建立角色權限關聯 repository
func NewRolePermissionRepository(db *DB) RolePermissionRepository {
	return &rolePermissionRepository{db: db}
}

// Create 建立角色權限關聯
func (r *rolePermissionRepository) Create(rolePermission *models.RolePermission) error {
	return r.db.DB.Create(rolePermission).Error
}

// Delete 刪除角色權限關聯
func (r *rolePermissionRepository) Delete(roleID, permissionID uint) error {
	return r.db.DB.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&models.RolePermission{}).Error
}

// GetPermissionsByRoleID 根據角色 ID 獲取權限列表
func (r *rolePermissionRepository) GetPermissionsByRoleID(roleID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.DB.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).Find(&permissions).Error
	return permissions, err
}

// GetRolesByPermissionID 根據權限 ID 獲取角色列表
func (r *rolePermissionRepository) GetRolesByPermissionID(permissionID uint) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.DB.Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Where("role_permissions.permission_id = ?", permissionID).Find(&roles).Error
	return roles, err
}

// Exists 檢查角色權限關聯是否存在
func (r *rolePermissionRepository) Exists(roleID, permissionID uint) (bool, error) {
	var count int64
	err := r.db.DB.Model(&models.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).Count(&count).Error
	return count > 0, err
}

// === UserRole Repository 實作 ===

// userRoleRepository 使用者角色關聯資料存取實作
type userRoleRepository struct {
	db *DB
}

// NewUserRoleRepository 建立使用者角色關聯 repository
func NewUserRoleRepository(db *DB) UserRoleRepository {
	return &userRoleRepository{db: db}
}

// Create 建立使用者角色關聯
func (r *userRoleRepository) Create(userRole *models.UserRole) error {
	return r.db.DB.Create(userRole).Error
}

// Delete 刪除使用者角色關聯
func (r *userRoleRepository) Delete(userID, roleID uint) error {
	return r.db.DB.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&models.UserRole{}).Error
}

// GetRolesByUserID 根據使用者 ID 獲取角色列表
func (r *userRoleRepository) GetRolesByUserID(userID uint) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.DB.Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).Find(&roles).Error
	return roles, err
}

// GetUsersByRoleID 根據角色 ID 獲取使用者列表
func (r *userRoleRepository) GetUsersByRoleID(roleID uint) ([]models.User, error) {
	var users []models.User
	err := r.db.DB.Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Where("user_roles.role_id = ?", roleID).Find(&users).Error
	return users, err
}

// Exists 檢查使用者角色關聯是否存在
func (r *userRoleRepository) Exists(userID, roleID uint) (bool, error) {
	var count int64
	err := r.db.DB.Model(&models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).Count(&count).Error
	return count > 0, err
}
