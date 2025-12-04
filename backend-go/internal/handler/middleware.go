package handler

import (
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "No token provided"})
			return
		}

		// Verifikasi ID Token dari Firebase
		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := authClient.VerifyIDToken(c, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		// Simpan UID ke context
		c.Set("uid", token.UID)
		c.Next()
	}
}