package domain

import "context"

// 1. Perbaiki Struct: Ubah/Pastikan ada field "Title"
type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"` // Sebelumnya mungkin 'Name', ubah jadi 'Title'
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}

// 2. Perbaiki Interface: Tambahkan method Create
type EventRepository interface {
	GetAll(ctx context.Context) ([]Event, error)
	// Tambahkan ini karena kamu mau bikin fitur CreateNewEvent
	Create(ctx context.Context, event *Event) error 
}