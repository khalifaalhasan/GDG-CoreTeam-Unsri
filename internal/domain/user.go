package domain

import "context"

type User struct {
	ID        string `json:"id"`       
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`      
	JobDesc   string `json:"job_desc"`  
	CreatedAt string `json:"created_at"`
}

type UserRepository interface {
	
	Save(ctx context.Context, user *User) error
	
	FindByID(ctx context.Context, id string) (*User, error)

	GetAllMembers(ctx context.Context) ([]User, error)
}