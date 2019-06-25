package auth

import (
	"context"
	"errors"

	"firebase.google.com/go/auth"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	RegisterUser(ctx context.Context, user *auth.UserToCreate) error
	GetUserByID(ctx context.Context, id string) (*auth.UserRecord, error)
	LoginUser(ctx context.Context, username, password string) (*auth.UserRecord, error)
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

func (s *service) RegisterUser(ctx context.Context, user *auth.UserToCreate) error {
	logger := log.With(s.logger, "method", "RegisterUserService")

	if _, err := s.repository.RegisterUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return ErrInsertInRepository
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*auth.UserRecord, error) {
	if id == "" {
		return nil, ErrInvalidArgument
	}

	logger := log.With(s.logger, "method", "GetUserByIDService")
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, ErrQueryRepository
	}
	return user, nil
}

func (s *service) LoginUser(ctx context.Context, username, password string) (*auth.UserRecord, error) {
	if username == "" || password == "" {
		return nil, ErrInvalidArgument
	}

	logger := log.With(s.logger, "method", "LoginUser")
	user, err := s.repository.GetUserByID(ctx, username)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, ErrLoginUser
	}
	return user, nil
}
