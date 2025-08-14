package controllers

import (
	"erp/models"
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

	result := database.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立使用者"})
		return
	}

	// 回傳新建立的使用者資訊
	c.JSON(http.StatusCreated, models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// GetUsers 取得所有使用者
func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取使用者列表"})
		return
	}

	// 轉換為 UserResponse 以隱藏密碼等敏感資訊
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, userResponses)
}

// GetUserByID 根據 ID 取得特定使用者
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// UpdateUser 更新使用者資訊
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// 檢查使用者是否存在
	if err := database.First(&user, id).Error; err != nil {
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

	// 儲存變更
	if err := database.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法更新使用者"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// DeleteUser 刪除使用者
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// 檢查使用者是否存在
	if err := database.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	// 執行軟刪除
	if err := database.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法刪除使用者"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "使用者已成功刪除"})
}
