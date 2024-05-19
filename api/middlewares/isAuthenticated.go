package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// クッキーからauthTokenを取得
		tokenString, err := c.Cookie("authToken")

		// エラーメッセージをまとめる関数
		unauthorized := func() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "権限がありません。"})
			c.Abort()
		}

		if err != nil {
			unauthorized()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			unauthorized()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if idFloat, ok := claims["id"].(float64); ok {
				userId := uint(idFloat)
				c.Set("userID", userId)
				fmt.Println("UserID set in middleware:", userId)
				c.Next()
				return
			}
		}

		// 何らかの理由で認証が失敗した場合
		unauthorized()
	}
}
