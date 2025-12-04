package handler

import (
	"backend-go/internal/domain"
	"backend-go/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{service: s}
}

// GET /events
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	events, err := h.service.GetEvents(c.Request.Context()) // Panggil service
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}

// POST /events
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event domain.Event
	
	// Binding JSON body ke Struct
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data: " + err.Error()})
		return
	}

	// TODO: Ambil role user dari JWT Token (Middleware)
	// Untuk sementara kita hardcode role "core" agar bisa tes create
	userRole := "core" 

	if err := h.service.CreateNewEvent(c.Request.Context(), &event, userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event berhasil dibuat", "data": event})
}