package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTトークンを検証するミドルウェア
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Authorizationヘッダーを取得
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		// トークンを取得
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // "Bearer "がない場合
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			ctx.Abort()
			return
		}

		// トークンを解析
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// JWT署名の方法がHS256であることを確認
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// トークンが有効な場合、次のハンドラに進む
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("userID", claims["sub"])
			ctx.Set("email", claims["email"])
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}

		// リクエストを次のハンドラに渡す
		ctx.Next()
	}
}
