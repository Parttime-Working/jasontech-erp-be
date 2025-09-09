package controllers

import (
	"erp/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePermission 建立新權限
func CreatePermission(c *gin.Context) {
	var input struct {
		ModuleName  string `json:"module_name" binding:"required"`
		Resource    string `json:"resource" binding:"required"`
		Action      string `json:"action" binding:"required"`
		Code        string `json:"code" binding:"required"`
		DisplayName string `json:"display_name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission := models.Permission{
		ModuleName:  input.ModuleName,
		Resource:    input.Resource,
		Action:      input.Action,
		Code:        input.Code,
		DisplayName: input.DisplayName,
		Description: input.Description,
	}

	err := GetPermissionRepo().Create(&permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立權限"})
		return
	}

	c.JSON(http.StatusCreated, permission)
}

// GetPermissions 取得所有權限
func GetPermissions(c *gin.Context) {
	permissions, err := GetPermissionRepo().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取權限列表"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

// GetPermissionByID 根據 ID 取得特定權限
func GetPermissionByID(c *gin.Context) {
	id := c.Param("id")
	var permissionID uint
	if _, err := fmt.Sscanf(id, "%d", &permissionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的權限 ID"})
		return
	}

	// 這裡需要實作 PermissionRepo
	// permission, err := GetPermissionRepo().GetByID(permissionID)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "找不到權限"})
	// 	return
	// }

	permission := models.Permission{ID: permissionID} // 暫時假資料
	c.JSON(http.StatusOK, permission)
}

// UpdatePermission 更新權限
func UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	var permissionID uint
	if _, err := fmt.Sscanf(id, "%d", &permissionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的權限 ID"})
		return
	}

	var input struct {
		ModuleName  *string `json:"module_name"`
		Resource    *string `json:"resource"`
		Action      *string `json:"action"`
		Code        *string `json:"code"`
		DisplayName *string `json:"display_name"`
		Description *string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 這裡需要實作 PermissionRepo
	// permission, err := GetPermissionRepo().GetByID(permissionID)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "找不到權限"})
	// 	return
	// }

	// 更新欄位
	// if input.ModuleName != nil { permission.ModuleName = *input.ModuleName }
	// if input.Resource != nil { permission.Resource = *input.Resource }
	// if input.Action != nil { permission.Action = *input.Action }
	// if input.Code != nil { permission.Code = *input.Code }
	// if input.DisplayName != nil { permission.DisplayName = *input.DisplayName }
	// if input.Description != nil { permission.Description = *input.Description }

	// err = GetPermissionRepo().Update(&permission)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "無法更新權限"})
	// 	return
	// }

	permission := models.Permission{ID: permissionID} // 暫時假資料
	c.JSON(http.StatusOK, permission)
}

// DeletePermission 刪除權限
func DeletePermission(c *gin.Context) {
	id := c.Param("id")
	var permissionID uint
	if _, err := fmt.Sscanf(id, "%d", &permissionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的權限 ID"})
		return
	}

	// 這裡需要實作 PermissionRepo
	// err := GetPermissionRepo().Delete(permissionID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "無法刪除權限"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "權限已刪除"})
}

// AssignPermissionToRole 為角色分配權限
func AssignPermissionToRole(c *gin.Context) {
	roleIDStr := c.Param("roleId")
	permissionIDStr := c.Param("permissionId")

	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	permissionID, err := strconv.ParseUint(permissionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的權限 ID"})
		return
	}

	// 檢查使用者等級權限
	userLevel, exists := c.Get("level")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "無法獲取使用者等級"})
		return
	}

	// 根據使用者等級檢查權限
	userLevelStr, ok := userLevel.(string)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "使用者等級格式錯誤"})
		return
	}

	if userLevelStr == "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "一般使用者無法編輯權限"})
		return
	}

	// 如果是 admin，需要檢查使用者是否擁有該權限模組的管理權限
	if userLevelStr == "admin" {
		// 目前暫時允許所有 admin 編輯，但實際上應該檢查具體權限
		// 這裡需要實作：檢查使用者是否擁有該模組的管理權限
		// 例如：檢查使用者的角色是否包含 "manage_" + module 的權限
	}

	// 如果是 super_admin，可以編輯所有權限
	// 不需要額外檢查

	// 實作角色權限分配
	rolePermission := models.RolePermission{RoleID: uint(roleID), PermissionID: uint(permissionID)}
	err = GetRolePermissionRepo().Create(&rolePermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法分配權限"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "權限分配成功"})
}

// RemovePermissionFromRole 從角色移除權限
func RemovePermissionFromRole(c *gin.Context) {
	roleIDStr := c.Param("roleId")
	permissionIDStr := c.Param("permissionId")

	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	permissionID, err := strconv.ParseUint(permissionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的權限 ID"})
		return
	}

	// 檢查使用者等級權限
	userLevel, exists := c.Get("level")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "無法獲取使用者等級"})
		return
	}

	// 根據使用者等級檢查權限
	userLevelStr, ok := userLevel.(string)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "使用者等級格式錯誤"})
		return
	}

	if userLevelStr == "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "一般使用者無法編輯權限"})
		return
	}

	// 如果是 admin，需要檢查使用者是否擁有該權限模組的管理權限
	if userLevelStr == "admin" {
		// 目前暫時允許所有 admin 編輯，但實際上應該檢查具體權限
		// 這裡需要實作：檢查使用者是否擁有該模組的管理權限
	}

	// 如果是 super_admin，可以編輯所有權限

	// 實作角色權限移除
	err = GetRolePermissionRepo().Delete(uint(roleID), uint(permissionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法移除權限"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "權限移除成功"})
}
