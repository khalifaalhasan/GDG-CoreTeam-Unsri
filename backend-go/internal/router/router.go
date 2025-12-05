package router

import (
	"backend-go/internal/handler"
	"backend-go/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	eventHandler *handler.EventHandler,
	userHandler *handler.UserHandler,
	authMiddleware gin.HandlerFunc,
	userService *service.UserService,
) *gin.Engine {
	r := gin.Default()

	adminOnlyMiddleware := handler.RoleMiddleware(userService, "admin")

	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong", "database": "connected"})
		})

		// --- EVENT ROUTES ---
		events := v1.Group("/events")
		{
			//public routes
			events.GET("", eventHandler.GetAllEvents)
			events.GET("/:id", eventHandler.GetEventByID)

			//protected routes
			protected := events.Group("")
			protected.Use(authMiddleware, adminOnlyMiddleware)
			{
				protected.POST("", eventHandler.CreateEvent)
				protected.PUT("/:id", eventHandler.UpdateEvent)
				protected.DELETE("/:id", eventHandler.DeleteEvent)
			}
		}

		// --- USER ROUTES ---
		baseUsers := v1.Group("/users")
		baseUsers.Use(authMiddleware)
		{
			baseUsers.POST("/register", userHandler.Register)
			baseUsers.GET("/me", userHandler.GetMyProfile)
			baseUsers.PUT("/me", userHandler.UpdateMyProfile)

			// Admin-only routes
			adminUsers := baseUsers.Group("")
			adminUsers.Use(adminOnlyMiddleware)
			{
				adminUsers.GET("", userHandler.GetAllUsers)
				adminUsers.GET("/:id", userHandler.GetUserByID)
				adminUsers.PUT("/:id/job", userHandler.AssignJob)
			}

		}
	}
		return r

}


