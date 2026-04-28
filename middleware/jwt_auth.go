package middleware

import (
	"strings"

	"log/slog"
	"main/common"
	"main/util"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			common.Unauthorized(c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			common.Fail(c, "Token格式错误")
			c.Abort()
			return
		}

		claims, err := util.ParseToken(parts[1])
		if err != nil {
			common.Unauthorized(c)
			c.Abort()
			return
		}
		slog.Info("jwtAuth claims: ", claims)
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.UserRole)
		c.Next()
	}
}
