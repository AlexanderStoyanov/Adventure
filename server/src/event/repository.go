package event

import (
	"context"
	"errors"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/go-kit/kit/log"
	"google.golang.org/api/iterator"
)

var (
	ErrInsertInRepository = errors.New("User was not inserted in the repository")
	ErrQueryRepository    = errors.New("User was not found in the repository")
	ErrLoginUser          = errors.New("Authentication failed")
)

// Repository for users
type Repository interface {
	CreateEvent(ctx context.Context, event *Event) error
	GetEventList(ctx context.Context) ([]Event, error)
}

type eventRepository struct {
	firebaseApp *firebase.App
	logger      log.Logger
}

// NewEventRepository creates a new instance of event repo
func NewEventRepository(firebaseApp *firebase.App, logger log.Logger) (Repository, error) {
	return &eventRepository{
		firebaseApp: firebaseApp,
		logger:      log.With(logger, "userRep", "firestore"),
	}, nil
}

func (repo *eventRepository) CreateEvent(ctx context.Context, event *Event) error {
	client, err := repo.firebaseApp.Firestore(ctx)
	fmt.Println(event)
	_, _, err = client.Collection("events").Add(ctx, event)
	if err != nil {
		repo.logger.Log("error", err)
		return err
	}
	return nil
}

func (repo *eventRepository) GetEventList(ctx context.Context) ([]Event, error) {
	client, err := repo.firebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	var event Event
	var events []Event
	iter := client.Collection("events").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.DataTo(&event)
		events = append(events, event)
	}
	fmt.Println(events)
	return events, nil
}
