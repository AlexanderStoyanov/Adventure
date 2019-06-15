package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	projectID := "elliptical-city-243420"

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
}
