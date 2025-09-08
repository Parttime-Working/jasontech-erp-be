package controllers

import (
	"erp/db"
)

// 全局資料庫實例 (依賴注入)
var database *db.DB

// Repository 實例
var userRepo db.UserRepository
var roleRepo db.RoleRepository
var permissionRepo db.PermissionRepository
var rolePermissionRepo db.RolePermissionRepository
var userRoleRepo db.UserRoleRepository

// SetDB 設定資料庫依賴 (依賴注入)
func SetDB(dbInstance *db.DB) {
	database = dbInstance
	userRepo = db.NewUserRepository(dbInstance)
	roleRepo = db.NewRoleRepository(dbInstance)
	permissionRepo = db.NewPermissionRepository(dbInstance)
	rolePermissionRepo = db.NewRolePermissionRepository(dbInstance)
	userRoleRepo = db.NewUserRoleRepository(dbInstance)
}

// GetUserRepo 獲取使用者 repository
func GetUserRepo() db.UserRepository {
	return userRepo
}

// GetRoleRepo 獲取角色 repository
func GetRoleRepo() db.RoleRepository {
	return roleRepo
}

// GetPermissionRepo 獲取權限 repository
func GetPermissionRepo() db.PermissionRepository {
	return permissionRepo
}

// GetRolePermissionRepo 獲取角色權限關聯 repository
func GetRolePermissionRepo() db.RolePermissionRepository {
	return rolePermissionRepo
}

// GetUserRoleRepo 獲取使用者角色關聯 repository
func GetUserRoleRepo() db.UserRoleRepository {
	return userRoleRepo
}
