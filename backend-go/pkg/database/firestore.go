package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// Kita return Firestore Client DAN Auth Client
func InitFirebase(ctx context.Context) (*firestore.Client, *auth.Client, error) {
	// Pastikan path serviceAccountKey.json benar
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing app: %v", err)
	}

	// 1. Init Firestore
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	// 2. Init Auth (PENTING BUAT LOGIN)
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing auth: %v", err)
	}

	return firestoreClient, authClient, nil
}