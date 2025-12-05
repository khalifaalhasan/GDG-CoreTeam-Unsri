package main

import (
	"context"
	"log"

	// Import package internal
	"backend-go/internal/handler"
	"backend-go/internal/repository"
	"backend-go/internal/router" // <-- IMPORT ROUTER YANG BARU
	"backend-go/internal/service"
	"backend-go/pkg/database"
)

func main() {
	// 1. Setup Context & Database
	ctx := context.Background()
	firestoreClient, authClient, err := database.InitFirebase(ctx)
	if err != nil {
		log.Fatalf("Gagal connect ke Firebase: %v", err)
	}
	defer firestoreClient.Close()
	log.Println("✅ Sukses terhubung ke Firebase!")

	// 2. Wiring (Dependency Injection)
	
	// A. Event Module
	eventRepo := repository.NewFirebaseRepo(firestoreClient)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	// B. User Module
	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// C. Middleware
	authMiddleware := handler.AuthMiddleware(authClient)

	r := router.SetupRouter(eventHandler, userHandler, authMiddleware, userService)

	// 4. Jalankan Server
	log.Println("🚀 Server berjalan di http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}