package handler

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"backend-go/internal/service"
)


func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

	
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		
		
		token, err := authClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

	
		
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