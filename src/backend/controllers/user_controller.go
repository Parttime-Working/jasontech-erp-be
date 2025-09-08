package controllers

import (
	"erp/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser 建立新使用者
func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 雜湊密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法處理密碼"})
		return
	}

	// 建立使用者
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// 設置等級：如果沒有指定，默認為 "user"
	if input.Level != "" {
		user.Level = input.Level
	} else {
		user.Level = "user"
	}

	err = GetUserRepo().Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立使用者"})
		return
	}

	// 回傳新建立的使用者資訊
	c.JSON(http.StatusCreated, models.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Level:       user.Level,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}

// GetUsers 取得所有使用者
func GetUsers(c *gin.Context) {
	users, err := GetUserRepo().GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取使用者列表"})
		return
	}

	// 轉換為 UserResponse 以隱藏密碼等敏感資訊
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Level:       user.Level,
			LastLoginAt: user.LastLoginAt,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, userResponses)
}

// GetUserByID 根據 ID 取得特定使用者
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// 將 string ID 轉換為 uint
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的使用者 ID"})
		return
	}

	user, err := GetUserRepo().GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Level:       user.Level,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}

// UpdateUser 更新使用者資訊
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// 將 string ID 轉換為 uint
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的使用者 ID"})
		return
	}

	// 檢查使用者是否存在
	user, err := GetUserRepo().GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	// 綁定更新資料
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新使用者欄位
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Password != nil {
		// 如果提供了新密碼，則進行雜湊處理
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "無法處理密碼"})
			return
		}
		user.Password = string(hashedPassword)
	}
	if input.Level != nil {
		// 檢查權限：只有管理員或超級管理員可以修改等級
		currentUserLevel, exists := c.Get("level")
		if !exists || (currentUserLevel != "admin" && currentUserLevel != "super_admin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "只有管理員或超級管理員可以修改使用者等級"})
			return
		}

		// 禁止修改最高管理員的等級 (ID=1 且等級為 super_admin)
		if user.ID == 1 && user.Level == "super_admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "無法修改系統最高管理員的等級"})
			return
		}

		// 驗證等級值
		if *input.Level != "super_admin" && *input.Level != "admin" && *input.Level != "user" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的等級值，只能是 'super_admin', 'admin' 或 'user'"})
			return
		}

		user.Level = *input.Level
	}

	// 儲存變更
	if err := GetUserRepo().Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法更新使用者"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Level:       user.Level,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}

// DeleteUser 刪除使用者
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// 將 string ID 轉換為 uint
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的使用者 ID"})
		return
	}

	// 檢查使用者是否存在
	user, err := GetUserRepo().GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	// 檢查權限：禁止刪除最高管理員 (ID=1 且角色為 admin)
	if user.ID == 1 && user.Level == "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "無法刪除系統最高管理員"})
		return
	}

	// 執行軟刪除
	if err := GetUserRepo().Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法刪除使用者"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "使用者已成功刪除"})
}
