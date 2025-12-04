package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	// Import package internal kita
	"backend-go/internal/handler"
	"backend-go/internal/repository"
	"backend-go/internal/service"
	"backend-go/pkg/database" // Import package database yang baru dibuat
)

func main() {
	// 1. Setup Context
	ctx := context.Background()

	// 2. Inisialisasi Database (Firestore)
	firestoreClient, err := database.InitFirestore(ctx)
	if err != nil {
		log.Fatalf("Gagal connect ke Firestore: %v", err)
	}
	// Penting: Tutup koneksi saat aplikasi mati
	defer firestoreClient.Close()
	
	log.Println("✅ Sukses terhubung ke Firebase Firestore!")

	// 3. Wiring Dependency Injection (Lapisan Bawang)
	// Repo butuh DB -> Service butuh Repo -> Handler butuh Service
	
	eventRepo := repository.NewFirebaseRepo(firestoreClient)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	// 4. Setup Router (Gin)
	r := gin.Default()

	// Grouping routes API v1
	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong", "database": "connected"})
		})

		// Event Routes
		events := v1.Group("/events")
		{
			events.GET("/", eventHandler.GetAllEvents)
			events.POST("/", eventHandler.CreateEvent) // Nanti tambahkan Auth Middleware disini
		}
	}

	// 5. Jalankan Server
	log.Println("🚀 Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}