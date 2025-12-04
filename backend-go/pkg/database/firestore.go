package database

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// InitFirestore menginisialisasi aplikasi Firebase dan mengembalikan client Firestore
func InitFirestore(ctx context.Context) (*firestore.Client, error) {
	// Best Practice: Gunakan environment variable untuk path key
	// Tapi untuk sekarang kita hardcode path relatif dulu agar mudah dicoba
	serviceAccountPath := "serviceAccountKey.json"
	
	// Pastikan path absolut (opsional, untuk keamanan jika run dari folder berbeda)
	absPath, err := filepath.Abs(serviceAccountPath)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan path absolut: %v", err)
	}

	opt := option.WithCredentialsFile(absPath)
	
	// Inisialisasi Firebase App
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error inisialisasi firebase app: %v", err)
	}

	// Inisialisasi Firestore Client
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error inisialisasi firestore client: %v", err)
	}

	return client, nil
}