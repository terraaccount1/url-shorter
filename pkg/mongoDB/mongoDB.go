package mongodb

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDatabase(ctx context.Context, dbname string, uri string) *mongo.Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(fmt.Errorf("failed to connect MongoDB : %w", err))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Errorf("failed to ping MongoDB : %w", err))
	}

	slog.Info("Successfully connected and pinged.")

	return client.Database(dbname)
}

func StopDatabase(ctx context.Context, db *mongo.Database) {
	slog.Info("Disconnect from MongoDB ...")
	db.Client().Disconnect(ctx)
	slog.Info("Successfully disconnected from MongoDB.")
}
