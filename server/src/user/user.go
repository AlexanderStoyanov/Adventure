package user

import (
	"context"
	"errors"
)

// User template structure
type User struct {
	ID       string     `firestore:"id"`
	Name     string     `firestore:"name,omitempty"`
	Username string     `firestore:"username"`
	Email    string     `firestore:"email"`
	Password string     `firestore:"password"`
	Location [2]float32 `firestore:"location,omitempty"`
}

var (
	ErrInsertInRepository = errors.New("User was not inserted in the repository")
	ErrQueryRepository    = errors.New("User was not found in the repository")
)

// Repository for users
type Repository interface {
	RegisterNewUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, id string) (User, error)
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
