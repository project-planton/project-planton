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

// Connect establishes a connection to MongoDB with retry logic.
// This is especially useful in containerized environments where MongoDB might not be immediately available.
func Connect(ctx context.Context, uri, databaseName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)

	var client *mongo.Client
	var err error

	// Retry connection up to 10 times (30 seconds total)
	maxRetries := 10
	retryDelay := 3 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		logrus.WithFields(logrus.Fields{
			"attempt": attempt,
			"max":     maxRetries,
		}).Info("Attempting to connect to MongoDB")

		// Set connection timeout for this attempt
		connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

		client, err = mongo.Connect(connectCtx, clientOptions)
		if err != nil {
			cancel()
			logrus.WithError(err).Warnf("MongoDB connection attempt %d/%d failed", attempt, maxRetries)

			if attempt < maxRetries {
				logrus.Infof("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			return nil, fmt.Errorf("failed to connect to MongoDB after %d attempts: %w", maxRetries, err)
		}

		// Ping the database to verify connection
		if err := client.Ping(connectCtx, nil); err != nil {
			cancel()
			logrus.WithError(err).Warnf("MongoDB ping attempt %d/%d failed", attempt, maxRetries)

			if attempt < maxRetries {
				logrus.Infof("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			return nil, fmt.Errorf("failed to ping MongoDB after %d attempts: %w", maxRetries, err)
		}

		cancel()
		logrus.Info("Successfully connected to MongoDB")
		break
	}

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

