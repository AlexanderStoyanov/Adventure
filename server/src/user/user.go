package user

import (
	"context"
	"errors"
)

// User template structure
type User struct {
	Name      string  `firestore:"name,omitempty"`
	Username  string  `firestore:"username"`
	Email     string  `firestore:"email"`
	Password  string  `firestore:"password"`
	Latitude  float64 `firestore:"latitude,omitempty"`
	Longitude float64 `firestore:"longitude,omitempty"`
}

var (
	ErrInsertInRepository = errors.New("User was not inserted in the repository")
	ErrQueryRepository    = errors.New("User was not found in the repository")
	ErrLoginUser          = errors.New("Authentication failed")
)

// Repository for users
type Repository interface {
	RegisterUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, id string) (User, error)
	//LoginUser(ctx context.Context, username, password string) error
}

// New creates a new user
func New(name, username, email, password string, latitude, longitude float64) *User {
	return &User{
		Name:      name,
		Username:  username,
		Email:     email,
		Password:  password,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
