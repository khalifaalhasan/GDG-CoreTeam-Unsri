package router

import (
	"backend-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// AddUserRoutes mengatur semua rute yang berkaitan dengan User
func AddUserRoutes(
	v1 *gin.RouterGroup,
	userHandler *handler.UserHandler,
	authMiddleware gin.HandlerFunc,
	adminOnlyMiddleware gin.HandlerFunc,
) {
	baseUsers := v1.Group("/users")
	baseUsers.Use(authMiddleware) // Semua rute user wajib Login
	{
		// PUBLIC ACCESS (Setiap user bisa akses dirinya sendiri)
		baseUsers.POST("/register", userHandler.Register)
		baseUsers.GET("/{id}", userHandler.GetMyProfile)
		baseUsers.PUT("/{id}", userHandler.UpdateMyProfile)

		// ADMIN ACCESS (Hanya untuk Role Admin)
		adminUsers := baseUsers.Group("")
		adminUsers.Use(adminOnlyMiddleware) // Wajib Role Admin
		{
			adminUsers.GET("", userHandler.GetAllUsers)
			adminUsers.GET("/:id", userHandler.GetUserByID)
			adminUsers.POST("/:id/job", userHandler.AssignJob) // Endpoint Admin
		}
	}
}