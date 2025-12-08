package database

import (
	"context"
	"fmt"

	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// Credential collection names
	AwsCredentialCollectionName        = "aws_credentials"
	GcpCredentialCollectionName        = "gcp_credentials"
	AzureCredentialCollectionName      = "azure_credentials"
	AtlasCredentialCollectionName      = "atlas_credentials"
	CloudflareCredentialCollectionName = "cloudflare_credentials"
	ConfluentCredentialCollectionName  = "confluent_credentials"
	SnowflakeCredentialCollectionName  = "snowflake_credentials"
	KubernetesCredentialCollectionName = "kubernetes_credentials"
)

// AwsCredentialRepository provides data access methods for AWS credentials.
type AwsCredentialRepository struct {
	collection *mongo.Collection
}

// NewAwsCredentialRepository creates a new AWS credential repository instance.
func NewAwsCredentialRepository(db *MongoDB) *AwsCredentialRepository {
	return &AwsCredentialRepository{
		collection: db.Database.Collection(AwsCredentialCollectionName),
	}
}

// FindByID retrieves an AWS credential by ID.
func (r *AwsCredentialRepository) FindByID(ctx context.Context, id string) (*models.AwsCredential, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid credential ID: %w", err)
	}

	var credential models.AwsCredential
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find AWS credential: %w", err)
	}
	return &credential, nil
}

// FindFirst retrieves the first AWS credential (for default/provider-based lookup).
func (r *AwsCredentialRepository) FindFirst(ctx context.Context) (*models.AwsCredential, error) {
	var credential models.AwsCredential
	err := r.collection.FindOne(ctx, bson.M{}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find AWS credential: %w", err)
	}
	return &credential, nil
}

// GcpCredentialRepository provides data access methods for GCP credentials.
type GcpCredentialRepository struct {
	collection *mongo.Collection
}

// NewGcpCredentialRepository creates a new GCP credential repository instance.
func NewGcpCredentialRepository(db *MongoDB) *GcpCredentialRepository {
	return &GcpCredentialRepository{
		collection: db.Database.Collection(GcpCredentialCollectionName),
	}
}

// FindByID retrieves a GCP credential by ID.
func (r *GcpCredentialRepository) FindByID(ctx context.Context, id string) (*models.GcpCredential, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid credential ID: %w", err)
	}

	var credential models.GcpCredential
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find GCP credential: %w", err)
	}
	return &credential, nil
}

// FindFirst retrieves the first GCP credential (for default/provider-based lookup).
func (r *GcpCredentialRepository) FindFirst(ctx context.Context) (*models.GcpCredential, error) {
	var credential models.GcpCredential
	err := r.collection.FindOne(ctx, bson.M{}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find GCP credential: %w", err)
	}
	return &credential, nil
}

// Create creates a new GCP credential.
func (r *GcpCredentialRepository) Create(ctx context.Context, credential *models.GcpCredential) (*models.GcpCredential, error) {
	result, err := r.collection.InsertOne(ctx, credential)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCP credential: %w", err)
	}

	credential.ID = result.InsertedID.(primitive.ObjectID)
	return credential, nil
}

// AzureCredentialRepository provides data access methods for Azure credentials.
type AzureCredentialRepository struct {
	collection *mongo.Collection
}

// NewAzureCredentialRepository creates a new Azure credential repository instance.
func NewAzureCredentialRepository(db *MongoDB) *AzureCredentialRepository {
	return &AzureCredentialRepository{
		collection: db.Database.Collection(AzureCredentialCollectionName),
	}
}

// FindByID retrieves an Azure credential by ID.
func (r *AzureCredentialRepository) FindByID(ctx context.Context, id string) (*models.AzureCredential, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid credential ID: %w", err)
	}

	var credential models.AzureCredential
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find Azure credential: %w", err)
	}
	return &credential, nil
}

// FindFirst retrieves the first Azure credential (for default/provider-based lookup).
func (r *AzureCredentialRepository) FindFirst(ctx context.Context) (*models.AzureCredential, error) {
	var credential models.AzureCredential
	err := r.collection.FindOne(ctx, bson.M{}).Decode(&credential)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find Azure credential: %w", err)
	}
	return &credential, nil
}
