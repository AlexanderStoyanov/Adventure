package firestoredb

import (
	"context"
	"fmt"
	userpkg "user"

	"github.com/mitchellh/mapstructure"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
	"github.com/mmcloughlin/geohash"
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

func (repo *userRepository) RegisterUser(ctx context.Context, user userpkg.User) error {
	DocumentRef := repo.db.Collection("users").Doc(user.Username)
	if dsnap, _ := DocumentRef.Get(ctx); dsnap.Data() != nil {
		return userpkg.ErrInsertInRepository
	}
	_, err := DocumentRef.Set(ctx, map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"geohash":  geohash.EncodeWithPrecision(user.Latitude, user.Longitude, 12),
	})
	return err
}

func (repo *userRepository) GetUserByID(ctx context.Context, username string) (userpkg.User, error) {
	dsnap, err := repo.db.Collection("users").Doc(username).Get(ctx)
	if err != nil {
		return userpkg.User{}, err
	}
	var user userpkg.User
	data := dsnap.Data()
	err = mapstructure.Decode(data, &user)
	user.Latitude, user.Longitude = geohash.Decode(data["geohash"].(string))
	fmt.Println(user)
	return user, nil
}
