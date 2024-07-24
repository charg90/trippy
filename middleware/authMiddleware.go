package middleware

import (
	"os"
	"strings"
	"trippy/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.ValidateToken(tokenString, []byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}