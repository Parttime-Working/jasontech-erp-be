package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供授權標頭"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "授權標頭格式錯誤"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 檢查簽名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("非預期的簽名方法: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "無效的 token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 將使用者資訊存入 context
			c.Set("user_id", claims["user_id"])
			c.Set("username", claims["username"])
			c.Set("level", claims["level"]) // 添加等級信息
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "無效的 token"})
		}
	}
}

// LevelMiddleware 檢查使用者等級是否為管理員或超級管理員
func LevelMiddleware(requiredLevels ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		level, exists := c.Get("level")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "無法獲取使用者等級"})
			return
		}

		levelStr, ok := level.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "使用者等級格式錯誤"})
			return
		}

		// 如果沒有指定需要的等級，默認需要 admin 或 super_admin
		if len(requiredLevels) == 0 {
			requiredLevels = []string{"admin", "super_admin"}
		}

		for _, reqLevel := range requiredLevels {
			if levelStr == reqLevel {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "需要更高的權限"})
	}
}

// AdminMiddleware 向後相容，檢查使用者是否為管理員或超級管理員
func AdminMiddleware() gin.HandlerFunc {
	return LevelMiddleware("admin", "super_admin")
}
