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


func (h *EventHandler) GetAllEvents(c *gin.Context) {
	events, err := h.service.GetEvents(c.Request.Context()) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}


func (h *EventHandler) GetEventByID(c *gin.Context){
	id := c.Param("id")
	event, err := h.service.GetEventByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data" : event})
}


func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event domain.Event
	

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data: " + err.Error()})
		return
	}


	userRole := "core" 

	if err := h.service.CreateNewEvent(c.Request.Context(), &event, userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event berhasil dibuat", "data": event})
}




func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id := c.Param("id") 
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


func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id") 
	if err := h.service.DeleteEvent(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}