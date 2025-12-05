package service

import (
	"backend-go/internal/domain"
	"context"
	"errors"
)

type EventService struct {
	repo domain.EventRepository
}

func NewEventService(repo domain.EventRepository) *EventService {
	return &EventService{repo: repo}
}

// --- INI METHOD YANG HILANG DAN MENYEBABKAN ERROR ---
func (s *EventService) GetEvents(ctx context.Context) ([]domain.Event, error) {
	// Memanggil repository GetAll
	return s.repo.GetAll(ctx)
}
// ---------------------------------------------------

func (s *EventService) CreateNewEvent(ctx context.Context, event *domain.Event, userRole string) error {
	// Logika Bisnis: Hanya Core Team yang bisa buat event
	if userRole != "core" && userRole != "lead" {
		return errors.New("unauthorized: only core team can create events")
	}

	// Validasi tambahan
	if event.Title == "" {
		return errors.New("event title is required")
	}

	return s.repo.Create(ctx, event)
}

// 1. Get Event By ID
func (s *EventService) GetEventByID(ctx context.Context, id string) (*domain.Event, error) {
	return s.repo.GetByID(ctx, id)
}

// ... (code sebelumnya) ...

// UPDATE Event Logic
func (s *EventService) UpdateEvent(ctx context.Context, id string, event *domain.Event) error {
	if id == "" {
		return errors.New("id is required")
	}
    // Pastikan ID di struct sama dengan ID di URL
    event.ID = id 
	return s.repo.Update(ctx, id, event)
}

// DELETE Event Logic
func (s *EventService) DeleteEvent(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.repo.Delete(ctx, id)
}