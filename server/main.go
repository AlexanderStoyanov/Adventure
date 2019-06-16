package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"

	"github.com/AlexanderStoyanov/Adventure/server/register"
	"github.com/AlexanderStoyanov/Adventure/server/user"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	projectID := "elliptical-city-243420"

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	fmt.Println(reflect.TypeOf(client))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()

	var rs = register.Service
	rs = register.NewService()
	rs.RegisterNewUser(ctx, user.New("Alex", "sancho100", "san", "bla@bla", "bla", [1.523, 2.2132]), client)
}
