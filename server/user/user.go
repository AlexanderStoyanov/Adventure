package user

import "context"

// User template structure
type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name,omitempty"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Location [2]float32 `json:"location,omitempty"`
}

// Repository for users
type Repository interface {
	CreateUser(ctx context.Context, user User) error
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
