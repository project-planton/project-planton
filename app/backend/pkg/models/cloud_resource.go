package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CloudResource represents a cloud resource document in MongoDB.
type CloudResource struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Kind      string             `bson:"kind" json:"kind"`
	Manifest  string             `bson:"manifest" json:"manifest"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

