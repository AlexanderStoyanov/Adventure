package auth

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/go-kit/kit/log"
)

var (
	ErrInsertInRepository = errors.New("User was not inserted in the repository")
	ErrQueryRepository    = errors.New("User was not found in the repository")
	ErrLoginUser          = errors.New("Authentication failed")
)

// Repository for users
type Repository interface {
	RegisterUser(ctx context.Context, user *auth.UserToCreate) (*auth.UserRecord, error)
	GetUserByID(ctx context.Context, id string) (*auth.UserRecord, error)
}

type userRepository struct {
	firebaseApp *firebase.App
	logger      log.Logger
}

// NewUserRepository creates a new instance of user repo
func NewUserRepository(firebaseApp *firebase.App, logger log.Logger) (Repository, error) {
	return &userRepository{
		firebaseApp: firebaseApp,
		logger:      log.With(logger, "userRep", "firestore"),
	}, nil
}

func (repo *userRepository) RegisterUser(ctx context.Context, user *auth.UserToCreate) (*auth.UserRecord, error) {
	client, err := repo.firebaseApp.Auth(ctx)
	u, err := client.CreateUser(ctx, user)
	if err != nil {
		repo.logger.Log("error", err)
		return nil, err
	}
	return u, nil
}

func (repo *userRepository) GetUserByID(ctx context.Context, uid string) (*auth.UserRecord, error) {
	client, err := repo.firebaseApp.Auth(ctx)
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		repo.logger.Log("error", err)
		return nil, err
	}
	return u, nil
}
