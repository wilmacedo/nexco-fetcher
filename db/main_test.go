package db

import (
	"context"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := NewMongoClient(ctx, "mongodb://user:pass@0.0.0.0:27017")
	if err != nil {
		t.Errorf("NewMongoClient() error=%v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		t.Errorf("Ping() error=%v", err)
	}
}
