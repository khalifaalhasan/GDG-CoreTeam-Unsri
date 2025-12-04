package domain

type User struct {
	ID    string `json:"id" firestore:"id"`
	Name  string `json:"name" firestore:"name"`
	Email string `json:"email" firestore:"email"`
	Role  string `json:"role" firestore:"role"` // 'core' or 'member'
}