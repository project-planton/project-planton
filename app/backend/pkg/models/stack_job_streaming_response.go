package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StackJobStreamingResponse represents a streaming output chunk from a Pulumi deployment.
// Each chunk is stored separately to enable real-time monitoring and complete log history.
type StackJobStreamingResponse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StackJobID  string             `bson:"stack_job_id" json:"stack_job_id"` // Foreign key to stackjobs collection
	Content     string             `bson:"content" json:"content"`           // The actual output content (line or chunk)
	StreamType  string             `bson:"stream_type" json:"stream_type"`   // "stdout" or "stderr"
	SequenceNum int                `bson:"sequence_num" json:"sequence_num"` // Order of this chunk (for reconstruction)
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`     // Timestamp when this chunk was received
}
