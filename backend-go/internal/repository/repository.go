package repository

import (
	"backend-go/internal/domain"
	"context"
)

// EventRepository mendefinisikan kontrak apa saja yg bisa dilakukan ke DB
type EventRepository interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
	GetEvents(ctx context.Context) ([]domain.Event, error)
	GetEventByID(ctx context.Context, id string) (*domain.Event, error)
}