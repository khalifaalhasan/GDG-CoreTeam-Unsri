package handler

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"backend-go/internal/service"
)

// AuthMiddleware memvalidasi token dari Firebase
func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil Header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// 2. Format harus "Bearer <TOKEN>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		
		// 3. Verifikasi Token ke Firebase
		token, err := authClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// 4. Token valid! Simpan UID user ke context Gin
		// Supaya bisa dipakai di Handler nanti (misal: c.MustGet("uid"))
		c.Set("uid", token.UID)
		c.Next()
	}
}

func RoleMiddleware(userService *service.UserService, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.MustGet("uid").(string)

		user, err := userService.GetUserProfile(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error" : "User data not found or role error"})
			c.Abort()
			return
		}

		if user.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error":"Acces denied : required role " + requiredRole})
			c.Abort()
			return
		}

		c.Next()
	}
}