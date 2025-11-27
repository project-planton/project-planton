package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB wraps a MongoDB client and provides database access.
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Connect establishes a connection to MongoDB.
func Connect(ctx context.Context, uri, databaseName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)

	// Set connection timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	logrus.Info("Successfully connected to MongoDB")

	db := client.Database(databaseName)

	return &MongoDB{
		Client:   client,
		Database: db,
	}, nil
}

// Disconnect closes the MongoDB connection.
func (m *MongoDB) Disconnect(ctx context.Context) error {
	if m.Client != nil {
		if err := m.Client.Disconnect(ctx); err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
		}
		logrus.Info("Disconnected from MongoDB")
	}
	return nil
}

