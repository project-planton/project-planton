package database

import (
	"context"
	"fmt"

	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// CollectionName is the name of the MongoDB collection for deployment components.
	CollectionName = "deployment_components"
)

// DeploymentComponentRepository provides data access methods for deployment components.
type DeploymentComponentRepository struct {
	collection *mongo.Collection
}

// NewDeploymentComponentRepository creates a new repository instance.
func NewDeploymentComponentRepository(db *MongoDB) *DeploymentComponentRepository {
	return &DeploymentComponentRepository{
		collection: db.Database.Collection(CollectionName),
	}
}

// ListOptions contains options for listing deployment components.
type ListOptions struct {
	Provider *string
	Kind     *string
}

// List retrieves deployment components from MongoDB with optional filters.
func (r *DeploymentComponentRepository) List(ctx context.Context, opts *ListOptions) ([]*models.DeploymentComponent, error) {
	filter := bson.M{}

	if opts != nil {
		if opts.Provider != nil && *opts.Provider != "" {
			filter["provider"] = *opts.Provider
		}
		if opts.Kind != nil && *opts.Kind != "" {
			filter["kind"] = *opts.Kind
		}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query deployment components: %w", err)
	}
	defer cursor.Close(ctx)

	var components []*models.DeploymentComponent
	if err := cursor.All(ctx, &components); err != nil {
		return nil, fmt.Errorf("failed to decode deployment components: %w", err)
	}

	return components, nil
}

