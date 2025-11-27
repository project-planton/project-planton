package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeploymentComponent represents a deployment component document in MongoDB.
type DeploymentComponent struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Kind          string             `bson:"kind" json:"kind"`
	Provider      string             `bson:"provider" json:"provider"`
	Name          string             `bson:"name" json:"name"`
	Version       string             `bson:"version" json:"version"`
	IDPrefix      string             `bson:"id_prefix" json:"id_prefix"`
	IsServiceKind bool               `bson:"is_service_kind" json:"is_service_kind"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

