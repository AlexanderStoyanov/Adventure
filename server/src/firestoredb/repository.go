package firestoredb

import (
	"context"
	userpkg "user"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
)

type userRepository struct {
	db     *firestore.Client
	logger log.Logger
}

// NewUserRepository creates a new instance of user repo
func NewUserRepository(db *firestore.Client, logger log.Logger) (userpkg.Repository, error) {
	return &userRepository{
		db:     db,
		logger: log.With(logger, "userRep", "firestore"),
	}, nil
}

func (repo *userRepository) RegisterNewUser(ctx context.Context, user userpkg.User) error {
	_, _, err := repo.db.Collection("users").Add(ctx, map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"location": user.Location,
	})
	return err
}

func (repo *userRepository) GetUserByID(ctx context.Context, id string) (userpkg.User, error) {
	//stub
	return userpkg.User{}, nil
}
