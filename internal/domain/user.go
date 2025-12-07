package domain

import "context"

type User struct {
	ID        string `json:"id"`        // Ini akan diisi UID dari Firebase Auth
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`      // "admin" atau "member"
	JobDesc   string `json:"job_desc"`  // Tugas khusus dari admin
	CreatedAt string `json:"created_at"`
}

type UserRepository interface {
	// Create atau Update user saat login pertama kali
	Save(ctx context.Context, user *User) error
	
	// Cari user berdasarkan ID (Firebase UID)
	FindByID(ctx context.Context, id string) (*User, error)
	
	// Khusus Admin: Lihat semua member
	GetAllMembers(ctx context.Context) ([]User, error)
}