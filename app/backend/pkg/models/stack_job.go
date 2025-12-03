package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StackJob represents a Pulumi stack deployment job in MongoDB.
type StackJob struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CloudResourceID string             `bson:"cloud_resource_id" json:"cloud_resource_id"`
	Status          string             `bson:"status" json:"status"`                     // success, failed, in_progress
	Output          string             `bson:"output,omitempty" json:"output,omitempty"` // JSON string containing Pulumi output
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
