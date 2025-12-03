package database

import (
	"context"
	"fmt"
	"time"

	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// StackJobCollectionName is the name of the MongoDB collection for stack jobs.
	StackJobCollectionName = "stackjobs"
)

// StackJobRepository provides data access methods for stack jobs.
type StackJobRepository struct {
	collection *mongo.Collection
}

// NewStackJobRepository creates a new repository instance.
func NewStackJobRepository(db *MongoDB) *StackJobRepository {
	return &StackJobRepository{
		collection: db.Database.Collection(StackJobCollectionName),
	}
}

// Create inserts a new stack job into MongoDB.
func (r *StackJobRepository) Create(ctx context.Context, job *models.StackJob) (*models.StackJob, error) {
	now := time.Now()
	job.ID = primitive.NewObjectID()
	job.CreatedAt = now
	job.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("failed to insert stack job: %w", err)
	}

	job.ID = result.InsertedID.(primitive.ObjectID)
	return job, nil
}

// FindByID retrieves a stack job by ID.
func (r *StackJobRepository) FindByID(ctx context.Context, id string) (*models.StackJob, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	var job models.StackJob
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&job)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Not found, but not an error
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query stack job by ID: %w", err)
	}
	return &job, nil
}

// FindByCloudResourceID retrieves stack jobs by cloud resource ID, sorted by created_at descending.
func (r *StackJobRepository) FindByCloudResourceID(ctx context.Context, cloudResourceID string) ([]*models.StackJob, error) {
	filter := bson.M{"cloud_resource_id": cloudResourceID}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}}) // Newest first

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query stack jobs: %w", err)
	}
	defer cursor.Close(ctx)

	var jobs []*models.StackJob
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, fmt.Errorf("failed to decode stack jobs: %w", err)
	}

	return jobs, nil
}

// Update updates an existing stack job in MongoDB.
func (r *StackJobRepository) Update(ctx context.Context, id string, job *models.StackJob) (*models.StackJob, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	job.ID = objectID
	job.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{},
	}

	// Only update fields that are provided (non-empty for strings)
	if job.Status != "" {
		update["$set"].(bson.M)["status"] = job.Status
	}
	if job.Output != "" {
		update["$set"].(bson.M)["output"] = job.Output
	}
	update["$set"].(bson.M)["updated_at"] = job.UpdatedAt

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedJob models.StackJob
	err = r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		update,
		opts,
	).Decode(&updatedJob)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("stack job not found")
		}
		return nil, fmt.Errorf("failed to update stack job: %w", err)
	}

	return &updatedJob, nil
}

// List retrieves stack jobs with optional filters.
func (r *StackJobRepository) List(ctx context.Context, cloudResourceID *string, status *string) ([]*models.StackJob, error) {
	filter := bson.M{}

	if cloudResourceID != nil && *cloudResourceID != "" {
		filter["cloud_resource_id"] = *cloudResourceID
	}

	if status != nil && *status != "" {
		filter["status"] = *status
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}}) // Newest first

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query stack jobs: %w", err)
	}
	defer cursor.Close(ctx)

	var jobs []*models.StackJob
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, fmt.Errorf("failed to decode stack jobs: %w", err)
	}

	return jobs, nil
}
