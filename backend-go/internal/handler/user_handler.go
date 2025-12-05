package handler

import (
	"backend-go/internal/domain"
	"backend-go/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GettAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.service.GetUserProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users/register
// Endpoint ini dipanggil Frontend SETELAH berhasil login di Firebase
// Gunanya untuk menyimpan data nama/email/role ke database kita
// internal/handler/user_handler.go

func (h *UserHandler) Register(c *gin.Context) {
	// Definisikan struct input yang HANYA berisi field yang AMAN
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		// Kita sengaja TIDAK memasukkan field Role di sini
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Ambil UID dari Token (yang diset oleh Middleware)
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UID not found in token"})
		return
	}
	
	// 2. Buat objek User akhir dengan data dari input dan default role
	user := domain.User{
		ID:    uid.(string),
		Name:  input.Name,
		Email: input.Email,
		// user.Role akan otomatis bernilai "" (empty string)
	}

	if err := h.service.RegisterUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "data": user})
}

// GET /users/me (Dashboard Pribadi)
func (h *UserHandler) GetMyProfile(c *gin.Context) {
	// Ambil UID dari Token Middleware
	uid := c.MustGet("uid").(string)

	user, err := h.service.GetUserProfile(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//PUT users/me
func (h *UserHandler) UpdateMyProfile(c *gin.Context){
	uid := c.MustGet("uid").(string)
	
	var input struct {
		Name string `json:"name" binding:"required"`
		
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	if err := h.service.UpdateUserProfile(c.Request.Context(), uid, input.Name, ""); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

//asign job to user (admin feature)
func (h *UserHandler) AssignJob(c *gin.Context) {
	targetUserID := c.Param("id")
	var input struct {
		Job string `json:"job_desc" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job description is required"})
		return
	}

	if err := h.service.AssignJob(c.Request.Context(), targetUserID, input.Job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job assigned successfully"})
}