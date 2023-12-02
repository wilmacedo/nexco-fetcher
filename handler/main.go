package handler

import (
	"context"
	"os"

	"github.com/wilmacedo/nexco-fetcher/db"
)

func HandleRequest(ctx context.Context) (*string, error) {
	_, err := db.NewMongoClient(ctx, os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, err
	}

	message := "ok"
	return &message, nil
}
