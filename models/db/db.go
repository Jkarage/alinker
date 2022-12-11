package db

import (
	"context"
	"time"

	"github.com/jkarage/alinker/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect tries making a connection to the database.
// Returns the connection, context, cancel if Ok
// Returns an error if found one
func Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	connectionString, err := env.Env("DB_CONNECTION", "")
	if err != nil {
		return nil, ctx, cancel, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	return client, ctx, cancel, err
}
