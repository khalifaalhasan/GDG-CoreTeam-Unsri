package repository

import (
	"context"
	"log"

	"backend-go/internal/domain"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type firebaseRepo struct {
	client *firestore.Client
}

func NewFirebaseRepo(client *firestore.Client) domain.EventRepository {
	return &firebaseRepo{
		client: client,
	}
}

// Method 1: GetAll (Sudah ada)
func (r *firebaseRepo) GetAll(ctx context.Context) ([]domain.Event, error) {
	var events []domain.Event

	iter := r.client.Collection("events").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error iterating documents: %v", err)
			return nil, err
		}

		var event domain.Event
		if err := doc.DataTo(&event); err != nil {
			log.Printf("Error mapping data: %v", err)
			continue
		}
		// Opsional: Set ID dari doc ID firestore jika field ID kosong
		if event.ID == "" {
			event.ID = doc.Ref.ID
		}
		events = append(events, event)
	}

	return events, nil
}

// Method 2: Create (TAMBAHKAN INI UNTUK FIX ERROR)
func (r *firebaseRepo) Create(ctx context.Context, event *domain.Event) error {
	// Menambahkan data ke collection "events" di Firestore
	// Firestore akan otomatis generate ID jika pakai .Add
	// Jika ingin set ID manual, gunakan .Doc("id").Set(...)
	
	ref, _, err := r.client.Collection("events").Add(ctx, event)
	if err != nil {
		log.Printf("Error creating event: %v", err)
		return err
	}
	
	// Update struct ID dengan ID yang baru digenerate Firestore (opsional, tapi berguna)
	event.ID = ref.ID
	
	return nil
}