package controllers

import (
	"erp/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRole 建立新角色
func CreateRole(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{
		Name:        input.Name,
		Description: input.Description,
	}

	err := GetRoleRepo().Create(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立角色"})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetRoles 取得所有角色
func GetRoles(c *gin.Context) {
	roles, err := GetRoleRepo().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取角色列表"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GetRoleByID 根據 ID 取得特定角色
func GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	var roleID uint
	if _, err := fmt.Sscanf(id, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	role, err := GetRoleRepo().GetByID(roleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到角色"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var roleID uint
	if _, err := fmt.Sscanf(id, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := GetRoleRepo().GetByID(roleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到角色"})
		return
	}

	// 更新欄位
	if input.Name != nil {
		role.Name = *input.Name
	}
	if input.Description != nil {
		role.Description = *input.Description
	}

	err = GetRoleRepo().Update(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法更新角色"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// DeleteRole 刪除角色
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	var roleID uint
	if _, err := fmt.Sscanf(id, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	err := GetRoleRepo().Delete(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法刪除角色"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色已刪除"})
}

// AssignRoleToUser 為用戶分配角色
func AssignRoleToUser(c *gin.Context) {
	userIDStr := c.Param("userId")
	roleIDStr := c.Param("roleId")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的用戶 ID"})
		return
	}

	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	// 檢查用戶等級權限
	userLevel, exists := c.Get("level")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "無法獲取用戶等級"})
		return
	}

	if userLevel != "admin" && userLevel != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有管理員可以分配角色"})
		return
	}

	// 實作使用者角色分配
	userRole := models.UserRole{UserID: uint(userID), RoleID: uint(roleID)}
	err = GetUserRoleRepo().Create(&userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法分配角色"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色分配成功"})
}

// RemoveRoleFromUser 從用戶移除角色
func RemoveRoleFromUser(c *gin.Context) {
	userIDStr := c.Param("userId")
	roleIDStr := c.Param("roleId")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的用戶 ID"})
		return
	}

	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的角色 ID"})
		return
	}

	// 檢查用戶等級權限
	userLevel, exists := c.Get("level")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "無法獲取用戶等級"})
		return
	}

	if userLevel != "admin" && userLevel != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有管理員可以移除角色"})
		return
	}

	// 實作使用者角色移除
	err = GetUserRoleRepo().Delete(uint(userID), uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法移除角色"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色移除成功"})
}
