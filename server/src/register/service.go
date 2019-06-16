package register

import (
	"context"
	"errors"

	//userpkg "github.com/AlexanderStoyanov/Adventure/server/user"
	userpkg "user"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	Create(ctx context.Context, user userpkg.User) error
	GetByID(ctx context.Context, id string) (userpkg.User, error)
}

type service struct {
	repository userpkg.Repository
	logger     log.Logger
}

// NewService creates new service
func NewService(rep userpkg.Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, user userpkg.User) error {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user.ID = id

	logger := log.With(s.logger, "method", "Create")
	err := s.repository.RegisterNewUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return userpkg.ErrInsertInRepository
	}
	return nil
}

func (s *service) GetByID(ctx context.Context, id string) (userpkg.User, error) {
	if id == "" {
		return userpkg.User{}, ErrInvalidArgument
	}

	logger := log.With(s.logger, "method", "GetByID")
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, userpkg.ErrQueryRepository
	}
	return user, nil
}
