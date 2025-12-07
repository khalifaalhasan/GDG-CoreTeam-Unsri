package main

import (
	"context"
	"log"

	"backend-go/internal/handler"
	"backend-go/internal/repository"
	"backend-go/internal/router" 
	"backend-go/internal/service"
	"backend-go/pkg/database"
)

func main() {
	
	ctx := context.Background()
	firestoreClient, authClient, err := database.InitFirebase(ctx)
	if err != nil {
		log.Fatalf("Gagal connect ke Firebase: %v", err)
	}
	defer firestoreClient.Close()
	log.Println("✅ Sukses terhubung ke Firebase!")

	
	eventRepo := repository.NewFirebaseRepo(firestoreClient)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)


	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authMiddleware := handler.AuthMiddleware(authClient)

	r := router.SetupRouter(eventHandler, userHandler, authMiddleware, userService)

	
	log.Println("🚀 Server berjalan di http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}