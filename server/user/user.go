package user

import (
	"context"
	"errors"
)

// User template structure
type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name,omitempty"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Location [2]float32 `json:"location,omitempty"`
}

var (
	ErrInsertInRepository = errors.New("User was not inserted in the repository")
	ErrQueryRepository    = errors.New("User was not found in the repository")
)

// Repository for users
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	//GetUserByID(ctx context.Context, id string) (User, error)
}

// New creates a new user
func New(name string, username string, email string, password string, location [2]float32) *User {
	return &User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
		Location: location,
	}
}
