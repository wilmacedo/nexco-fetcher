package handler

import (
	"context"
	"os"

	"github.com/wilmacedo/nexco-fetcher/db"
	"github.com/wilmacedo/nexco-fetcher/types"
)

func HandleRequest(ctx context.Context) (*types.Response, error) {
	_, err := db.NewMongoClient(ctx, os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Message: "ok",
	}, nil
}
