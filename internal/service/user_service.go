package service

import (
	"backend-go/internal/domain"
	"context"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// 1. Register/Sync User: Dipanggil saat user pertama kali login
// internal/service/user_service.go

func (s *UserService) RegisterUser(ctx context.Context, user *domain.User) error {
	// Logika ini akan selalu berjalan dan mengeset role=member 
    // karena input dari handler sudah difilter
	if user.Role == "" {
		user.Role = "member"
	}
	return s.repo.Save(ctx, user)
}

// 2. Get User Profile: Dipanggil untuk dashboard user
func (s *UserService) GetUserProfile(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) GettAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAllMembers(ctx)
}

// 3. Update User Job (Khusus Admin)
func (s *UserService) AssignJob(ctx context.Context, userID string, job string) error {
	// Ambil data user dulu
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	
	// Update job-nya
	user.JobDesc = job
	
	// Simpan balik ke database
	return s.repo.Save(ctx, user)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, id string, newName string, email string) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if newName != "" {
		user.Name = newName
	}

	return s.repo.Save(ctx, user)
}

// 4. Get All Members (Khusus Admin)
func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAllMembers(ctx)
}