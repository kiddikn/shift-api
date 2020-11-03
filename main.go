// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	// 環境変数から取得
	//	projectID := "otakeslsc-215813"
	projectID := "xx"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the id for the new topic.
	topicID := "my-topic"

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}
