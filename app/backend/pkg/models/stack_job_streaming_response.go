package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StackUpdateStreamingResponse represents a streaming output chunk from a Pulumi deployment.
// Each chunk is stored separately to enable real-time monitoring and complete log history.
type StackUpdateStreamingResponse struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StackUpdateID string             `bson:"stack_update_id" json:"stack_update_id"` // Foreign key to stackupdates collection
	Content       string             `bson:"content" json:"content"`                 // The actual output content (line or chunk)
	StreamType    string             `bson:"stream_type" json:"stream_type"`         // "stdout" or "stderr"
	SequenceNum   int                `bson:"sequence_num" json:"sequence_num"`       // Order of this chunk (for reconstruction)
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`           // Timestamp when this chunk was received
}
