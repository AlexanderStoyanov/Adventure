package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/go-kit/kit/log"

	"cloud.google.com/go/firestore"

	// "github.com/AlexanderStoyanov/Adventure/server/register"
	// "github.com/AlexanderStoyanov/Adventure/server/user"
	"firestoredb"
	"register"
	"user"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	projectID := "elliptical-city-243420"

	// Get a Firestore client.
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, projectID)
	fmt.Println(reflect.TypeOf(client))

	// Close client when done.
	defer client.Close()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var repo, _ = firestoredb.NewUserRepository(client, logger)

	var rs register.Service
	rs = register.NewService(repo, logger)
	arr := [2]float32{1.523, 2.2132}
	rs.Create(ctx, user.User{"ID", "Nastya", "san", "bla@bla", "bla", arr})
}
