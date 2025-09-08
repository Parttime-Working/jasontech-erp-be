package services

import (
	"erp/models"
	"strings"
)

// PermissionService 權限服務
type PermissionService struct {
}

// NewPermissionService 建立權限服務實例
func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

// HasPermission 檢查用戶是否擁有特定權限
func (s *PermissionService) HasPermission(userID uint, permissionCode string) bool {
	// 1. 獲取用戶等級
	userLevel := s.getUserLevel(userID)

	// 2. 超級管理員檢查：擁有所有權限
	if userLevel == "super_admin" {
		return true
	}

	// 3. 管理員權限檢查：可以訪問分配給他的模組權限
	if userLevel == "admin" {
		userModules := s.getUserAssignedModules(userID)
		for _, module := range userModules {
			if strings.HasPrefix(permissionCode, module+".") {
				return true
			}
		}
	}

	// 4. 一般用戶權限檢查：根據其所屬角色的權限進行判斷
	return s.checkUserRolePermissions(userID, permissionCode)
}

// getUserLevel 獲取用戶等級
func (s *PermissionService) getUserLevel(userID uint) string {
	// TODO: 實作從資料庫獲取用戶等級的邏輯
	// 這裡應該查詢 users 表中的 level 欄位
	return "user" // 臨時返回值
}

// getUserAssignedModules 獲取管理員被分配的模組
func (s *PermissionService) getUserAssignedModules(userID uint) []string {
	// TODO: 實作從資料庫獲取用戶分配模組的邏輯
	// 這裡應該查詢用戶的角色，然後獲取角色對應的權限模組
	return []string{"hr", "finance"} // 臨時返回值
}

// checkUserRolePermissions 檢查用戶角色權限
func (s *PermissionService) checkUserRolePermissions(userID uint, permissionCode string) bool {
	// TODO: 實作檢查用戶角色權限的邏輯
	// 1. 查詢用戶的所有角色
	// 2. 查詢每個角色的所有權限
	// 3. 檢查權限代碼是否匹配
	return false // 臨時返回值
}

// GetUserPermissions 獲取用戶的所有權限
func (s *PermissionService) GetUserPermissions(userID uint) ([]models.Permission, error) {
	// TODO: 實作獲取用戶權限列表的邏輯
	return []models.Permission{}, nil
}

// GetUserRoles 獲取用戶的所有角色
func (s *PermissionService) GetUserRoles(userID uint) ([]models.Role, error) {
	// TODO: 實作獲取用戶角色列表的邏輯
	return []models.Role{}, nil
}

// AssignRoleToUser 為用戶分配角色
func (s *PermissionService) AssignRoleToUser(userID, roleID uint) error {
	// TODO: 實作為用戶分配角色的邏輯
	return nil
}

// RemoveRoleFromUser 從用戶移除角色
func (s *PermissionService) RemoveRoleFromUser(userID, roleID uint) error {
	// TODO: 實作從用戶移除角色的邏輯
	return nil
}

// AssignPermissionToRole 為角色分配權限
func (s *PermissionService) AssignPermissionToRole(roleID, permissionID uint) error {
	// TODO: 實作為角色分配權限的邏輯
	return nil
}

// RemovePermissionFromRole 從角色移除權限
func (s *PermissionService) RemovePermissionFromRole(roleID, permissionID uint) error {
	// TODO: 實作從角色移除權限的邏輯
	return nil
}
