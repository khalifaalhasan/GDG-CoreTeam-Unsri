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

//get /events/:id
func (h *EventHandler) GetEventByID(c *gin.Context){
	id := c.Param("id")
	event, err := h.service.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data" : event})
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



// PUT /events/:id
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari URL
	var event domain.Event

    // Bind JSON body ke struct
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data: " + err.Error()})
		return
	}

	if err := h.service.UpdateEvent(c.Request.Context(), id, &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// DELETE /events/:id
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari URL
	if err := h.service.DeleteEvent(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}