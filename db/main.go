package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseService interface {
	Connect(ctx context.Context) error
	Ping(ctx context.Context) error
}

type MongoClient struct {
	Client *mongo.Client
}

func (c *MongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.Client.Ping(ctx, rp)
}

func NewMongoClient(ctx context.Context, url string) (*MongoClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client: client,
	}, err
}
