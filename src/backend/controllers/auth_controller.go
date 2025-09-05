package controllers

import (
	"erp/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 使用 Repository 查詢使用者
	user, err := GetUserRepo().GetByUsername(credentials.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 驗證密碼
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 更新最後登入時間
	now := time.Now()
	user.LastLoginAt = &now
	err = GetUserRepo().Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update last login time"})
		return
	}

	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role, // 添加角色信息到 token
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// 使用密鑰簽署令牌
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// 返回成功響應和令牌
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user": models.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Role:        user.Role,
			LastLoginAt: user.LastLoginAt,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	})
}
