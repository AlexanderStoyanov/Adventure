package auth

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/crypto/bcrypt"

	userpkg "github.com/AlexanderStoyanov/Adventure/server/src/user"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	RegisterUser(ctx context.Context, user userpkg.User) error
	GetUserByID(ctx context.Context, id string) (userpkg.User, error)
	LoginUser(ctx context.Context, username, password string) (userpkg.User, error)
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

func (s *service) RegisterUser(ctx context.Context, user userpkg.User) error {
	logger := log.With(s.logger, "method", "RegisterUserService")

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		user.Password = string(hashedPassword)
	} else if err != nil {
		level.Error(logger).Log("err", err)
	}

	if err := s.repository.RegisterUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return userpkg.ErrInsertInRepository
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (userpkg.User, error) {
	if id == "" {
		return userpkg.User{}, ErrInvalidArgument
	}

	logger := log.With(s.logger, "method", "GetUserByIDService")
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, userpkg.ErrQueryRepository
	}
	return user, nil
}

func (s *service) LoginUser(ctx context.Context, username, password string) (userpkg.User, error) {
	if username == "" || password == "" {
		return userpkg.User{}, ErrInvalidArgument
	}

	logger := log.With(s.logger, "method", "LoginUser")
	user, err := s.repository.GetUserByID(ctx, username)
	if err != nil {
		level.Error(logger).Log("err", err)
		return userpkg.User{}, userpkg.ErrLoginUser
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		level.Error(logger).Log("err", err)
		return userpkg.User{}, userpkg.ErrLoginUser
	}
	return user, nil
}
