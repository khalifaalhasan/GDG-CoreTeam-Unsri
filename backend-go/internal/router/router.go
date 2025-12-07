package router

import (
	"backend-go/internal/handler"
	"backend-go/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter adalah fungsi utama yang menyiapkan dan mengembalikan Gin Engine
func SetupRouter(
	eventHandler *handler.EventHandler,
	userHandler *handler.UserHandler,
	authMiddleware gin.HandlerFunc,
	userService *service.UserService,
) *gin.Engine {
	r := gin.Default()

	// 1. Definisikan Middleware Khusus
	// Middleware ini perlu akses ke userService untuk cek role di Firestore
	adminOnlyMiddleware := handler.RoleMiddleware(userService, "admin")

	// 2. Buat Group Utama /api/v1
	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong", "database": "connected"})
		})

		// 3. Panggil Sub-Router (modularitas)
		AddEventRoutes(v1, eventHandler, authMiddleware, adminOnlyMiddleware)
		AddUserRoutes(v1, userHandler, authMiddleware, adminOnlyMiddleware)
	}

	return r
}