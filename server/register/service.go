package register

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log/level"

	"github.com/AlexanderStoyanov/Adventure/server/user"
	"github.com/go-kit/kit/log"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	Create(ctx context.Context, user user.User) error
	GetByID(ctx context.Context, id string) (user.User, error)
}

type service struct {
	repository user.Repository
	logger     log.Logger
}

// NewService creates new service
func NewService(rep user.Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, user user.User) error {
	return nil
}

func (s *service) GetByID(ctx context.Context, id string) (user.User, error) {
	logger := log.With(s.logger, "method", "GetByID")
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, Service.ErrQueryRepository
	}
	return user, nil
}
