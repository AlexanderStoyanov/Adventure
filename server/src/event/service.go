package event

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	CreateEvent(ctx context.Context, event *Event) error
	GetEventList(ctx context.Context) ([]Event, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService creates new service
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateEvent(ctx context.Context, event *Event) error {
	logger := log.With(s.logger, "method", "CreateEventService")
	if err := s.repository.CreateEvent(ctx, event); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

func (s *service) GetEventList(ctx context.Context) ([]Event, error) {
	logger := log.With(s.logger, "method", "GetEventListService")
	events, err := s.repository.GetEventList(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, ErrQueryRepository
	}
	return events, nil
}
