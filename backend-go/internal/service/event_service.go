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