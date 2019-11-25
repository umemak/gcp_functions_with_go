package functions

import (
	"context"
	"fmt"
	"log"
	"time"
	"cloud.google.com/go/functions/metadata"
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name string `json:"name"`
	Metageneration string `json:"metageneration"`
	ResourceState string `json:"resoutceState"`
	TimeCreated time.Time `json:"timeCreated"`
	Updated time.Time `json:"updated"`
}

func TriggerStorage(ctx context.Context, e GCSEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Event ID: %v¥n", meta.EventID)
	log.Printf("Event type: %v¥n", meta.EventType)
	log.Printf("Bucket Name: %v¥n", e.Bucket)
	log.Printf("File Name: %v¥n", e.Name)
	log.Printf("Resource State: %v¥n", e.ResourceState)
	log.Printf("Created at: %v¥n", e.TimeCreated)
	log.Printf("Updated at: %v¥n", e.Updated)
	return nil
}

