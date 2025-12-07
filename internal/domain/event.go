package domain

import "context"


type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"` 
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}


type EventRepository interface {
	GetAll(ctx context.Context) ([]Event, error)
	GetByID(ctx context.Context, id string) (*Event, error)
	Create(ctx context.Context, event *Event) error 
	Update(ctx context.Context, id string, event *Event) error 
	Delete(ctx context.Context, id string) error 
}