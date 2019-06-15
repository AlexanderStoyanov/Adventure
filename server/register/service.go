package register

import (
	"errors"

	"github.com/AlexanderStoyanov/Adventure/server/user"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides register methods.
type Service interface {
	RegisterNewUser(user user.User) (string, error)
}

type service struct{}

func (s *service) RegisterNewUser(user user.User) (string, error) {
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"location": user.Location,
	})
}
