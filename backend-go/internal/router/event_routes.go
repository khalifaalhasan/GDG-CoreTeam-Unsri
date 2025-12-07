package router

import (
	"backend-go/internal/handler"
	"github.com/gin-gonic/gin"
)

// AddEventRoutes mengatur semua rute yang berkaitan dengan Event
func AddEventRoutes(
	v1 *gin.RouterGroup,
	eventHandler *handler.EventHandler,
	authMiddleware gin.HandlerFunc,
	adminOnlyMiddleware gin.HandlerFunc,
) {
	events := v1.Group("/events")
	{
		// PUBLIC ROUTES: Siapa pun bisa melihat event
		events.GET("", eventHandler.GetAllEvents)
		events.GET("/:id", eventHandler.GetEventByID)

		// PROTECTED ROUTES: Hanya Admin yang bisa CUD event
		protected := events.Group("")
		protected.Use(authMiddleware, adminOnlyMiddleware) // Wajib Login + Role Admin
		{
			protected.POST("", eventHandler.CreateEvent)
			protected.PUT("/:id", eventHandler.UpdateEvent)
			protected.DELETE("/:id", eventHandler.DeleteEvent)
		}
	}
}