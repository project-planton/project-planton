package database

import (
	"context"
	"fmt"
	"time"

	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// CloudResourceCollectionName is the name of the MongoDB collection for cloud resources.
	CloudResourceCollectionName = "cloud_resources"
)

// CloudResourceRepository provides data access methods for cloud resources.
type CloudResourceRepository struct {
	collection *mongo.Collection
}

// NewCloudResourceRepository creates a new repository instance.
func NewCloudResourceRepository(db *MongoDB) *CloudResourceRepository {
	return &CloudResourceRepository{
		collection: db.Database.Collection(CloudResourceCollectionName),
	}
}

// FindByName retrieves a cloud resource by name.
func (r *CloudResourceRepository) FindByName(ctx context.Context, name string) (*models.CloudResource, error) {
	var resource models.CloudResource
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&resource)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Not found, but not an error
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query cloud resource by name: %w", err)
	}
	return &resource, nil
}

// Create inserts a new cloud resource into MongoDB.
func (r *CloudResourceRepository) Create(ctx context.Context, resource *models.CloudResource) (*models.CloudResource, error) {
	now := time.Now()
	resource.ID = primitive.NewObjectID()
	resource.CreatedAt = now
	resource.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, resource)
	if err != nil {
		return nil, fmt.Errorf("failed to insert cloud resource: %w", err)
	}

	resource.ID = result.InsertedID.(primitive.ObjectID)
	return resource, nil
}

// CloudResourceListOptions contains options for listing cloud resources.
type CloudResourceListOptions struct {
	Kind *string
}

// List retrieves cloud resources from MongoDB with optional filters.
func (r *CloudResourceRepository) List(ctx context.Context, opts *CloudResourceListOptions) ([]*models.CloudResource, error) {
	filter := bson.M{}

	if opts != nil {
		if opts.Kind != nil && *opts.Kind != "" {
			filter["kind"] = *opts.Kind
		}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query cloud resources: %w", err)
	}
	defer cursor.Close(ctx)

	var resources []*models.CloudResource
	if err := cursor.All(ctx, &resources); err != nil {
		return nil, fmt.Errorf("failed to decode cloud resources: %w", err)
	}

	return resources, nil
}

