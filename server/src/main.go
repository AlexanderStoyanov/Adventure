package main

import (
	"context"
	"os"
	"user"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"

	// "github.com/AlexanderStoyanov/Adventure/server/register"
	// "github.com/AlexanderStoyanov/Adventure/server/user"
	"firestoredb"
	"register"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	projectID := "elliptical-city-243420"

	// Get a Firestore client.
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, projectID)

	// Close client when done.
	defer client.Close()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var repo, _ = firestoredb.NewUserRepository(client, logger)

	var rs register.Service
	rs = register.NewService(repo, logger)
	rs.RegisterUser(ctx, user.User{"Nastya", "san", "bla@bla123", "bla", 40.7128, -74.0060})
	rs.GetUserByID(ctx, "san")
}
