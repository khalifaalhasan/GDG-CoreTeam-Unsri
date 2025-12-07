package repository

import (
	"backend-go/internal/domain"
	"context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type userRepo struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) domain.UserRepository {
	return &userRepo{client: client}
}

// Save: Menyimpan data user (Upsert - kalau ada update, kalau belum create)
func (r *userRepo) Save(ctx context.Context, user *domain.User) error {
	// Gunakan UID dari Firebase sebagai ID Dokumen
	_, err := r.client.Collection("users").Doc(user.ID).Set(ctx, user)
	return err
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	doc, err := r.client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user domain.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	user.ID = doc.Ref.ID
	return &user, nil
}

func (r *userRepo) GetAllMembers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	// Ambil semua user
	iter := r.client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user domain.User
		doc.DataTo(&user)
		user.ID = doc.Ref.ID
		users = append(users, user)
	}
	return users, nil
}