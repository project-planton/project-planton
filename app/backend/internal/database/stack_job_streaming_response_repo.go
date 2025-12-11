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
	// StackJobStreamingResponseCollectionName is the name of the MongoDB collection for streaming responses.
	StackJobStreamingResponseCollectionName = "stackjob_streaming_responses"
)

// StackJobStreamingResponseRepository provides data access methods for streaming responses.
type StackJobStreamingResponseRepository struct {
	collection *mongo.Collection
}

// NewStackJobStreamingResponseRepository creates a new repository instance.
func NewStackJobStreamingResponseRepository(db *MongoDB) *StackJobStreamingResponseRepository {
	return &StackJobStreamingResponseRepository{
		collection: db.Database.Collection(StackJobStreamingResponseCollectionName),
	}
}

// Create inserts a new streaming response chunk into MongoDB.
func (r *StackJobStreamingResponseRepository) Create(ctx context.Context, response *models.StackJobStreamingResponse) (*models.StackJobStreamingResponse, error) {
	now := time.Now()
	response.ID = primitive.NewObjectID()
	response.CreatedAt = now

	result, err := r.collection.InsertOne(ctx, response)
	if err != nil {
		return nil, fmt.Errorf("failed to insert streaming response: %w", err)
	}

	response.ID = result.InsertedID.(primitive.ObjectID)
	return response, nil
}

// CreateBatch inserts multiple streaming response chunks in a single operation.
func (r *StackJobStreamingResponseRepository) CreateBatch(ctx context.Context, responses []*models.StackJobStreamingResponse) error {
	if len(responses) == 0 {
		return nil
	}

	now := time.Now()
	docs := make([]interface{}, len(responses))
	for i, resp := range responses {
		resp.ID = primitive.NewObjectID()
		resp.CreatedAt = now
		docs[i] = resp
	}

	_, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return fmt.Errorf("failed to insert streaming responses batch: %w", err)
	}

	return nil
}

// FindByStackJobID retrieves all streaming responses for a specific stack job, ordered by sequence number.
func (r *StackJobStreamingResponseRepository) FindByStackJobID(ctx context.Context, stackJobID string) ([]*models.StackJobStreamingResponse, error) {
	filter := bson.M{"stack_job_id": stackJobID}
	// Use bson.D for ordered sort (sequence_num first, then created_at)
	opts := options.Find().SetSort(bson.D{
		{Key: "sequence_num", Value: 1},
		{Key: "created_at", Value: 1},
	})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find streaming responses: %w", err)
	}
	defer cursor.Close(ctx)

	var responses []*models.StackJobStreamingResponse
	if err := cursor.All(ctx, &responses); err != nil {
		return nil, fmt.Errorf("failed to decode streaming responses: %w", err)
	}

	return responses, nil
}

// DeleteByStackJobID deletes all streaming responses for a specific stack job.
func (r *StackJobStreamingResponseRepository) DeleteByStackJobID(ctx context.Context, stackJobID string) error {
	filter := bson.M{"stack_job_id": stackJobID}
	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete streaming responses: %w", err)
	}

	_ = result // Result contains deleted count, but we don't need it for now
	return nil
}

// GetNextSequenceNum returns the next sequence number for a stack job.
func (r *StackJobStreamingResponseRepository) GetNextSequenceNum(ctx context.Context, stackJobID string) (int, error) {
	filter := bson.M{"stack_job_id": stackJobID}
	opts := options.FindOne().SetSort(bson.M{"sequence_num": -1})

	var lastResponse models.StackJobStreamingResponse
	err := r.collection.FindOne(ctx, filter, opts).Decode(&lastResponse)
	if err == mongo.ErrNoDocuments {
		// No previous responses, start at 0
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("failed to get next sequence number: %w", err)
	}

	return lastResponse.SequenceNum + 1, nil
}

// FindByStackJobIDAfterSequence retrieves streaming responses for a stack job after a specific sequence number.
// Used for resuming streams from a specific point.
func (r *StackJobStreamingResponseRepository) FindByStackJobIDAfterSequence(ctx context.Context, stackJobID string, afterSequenceNum int) ([]*models.StackJobStreamingResponse, error) {
	filter := bson.M{
		"stack_job_id": stackJobID,
		"sequence_num": bson.M{"$gt": afterSequenceNum},
	}
	// Use bson.D for ordered sort (sequence_num first, then created_at)
	opts := options.Find().SetSort(bson.D{
		{Key: "sequence_num", Value: 1},
		{Key: "created_at", Value: 1},
	})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find streaming responses: %w", err)
	}
	defer cursor.Close(ctx)

	var responses []*models.StackJobStreamingResponse
	if err := cursor.All(ctx, &responses); err != nil {
		return nil, fmt.Errorf("failed to decode streaming responses: %w", err)
	}

	return responses, nil
}
