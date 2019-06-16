package register

import (
	"context"
	"errors"

	"github.com/AlexanderStoyanov/Adventure/server/user"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	RegisterNewUser(ctx context.Context, user user.User) (string, error)
}

type service struct {
	repo user.Repository
}

func (s *service) RegisterNewUser(ctx context.Context, user user.User) (string, error) {
	_, _, _ = client.Collection("users").Add(ctx, map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"location": user.Location,
	})
	return "", nil
}

// NewService creates new service
func NewService() Service {
	return &service{}
}
