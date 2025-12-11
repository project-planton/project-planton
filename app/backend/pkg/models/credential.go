package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Credential represents a base credential document in MongoDB.
// Each provider has its own collection (e.g., aws_credentials, gcp_credentials).
type Credential struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Provider  string             `bson:"provider" json:"provider"` // "aws", "gcp", "azure", etc.
	Spec      interface{}        `bson:"spec" json:"spec"`         // Provider-specific credential spec
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// AwsCredential represents AWS credentials.
type AwsCredential struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	AccountID       string             `bson:"account_id" json:"account_id"`
	AccessKeyID     string             `bson:"access_key_id" json:"access_key_id"`
	SecretAccessKey string             `bson:"secret_access_key" json:"secret_access_key"`
	Region          string             `bson:"region,omitempty" json:"region,omitempty"`
	SessionToken    string             `bson:"session_token,omitempty" json:"session_token,omitempty"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

// GcpCredential represents GCP credentials.
type GcpCredential struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name                    string             `bson:"name" json:"name"`
	ServiceAccountKeyBase64 string             `bson:"service_account_key_base64" json:"service_account_key_base64"`
	CreatedAt               time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at" json:"updated_at"`
}

// AzureCredential represents Azure credentials.
type AzureCredential struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	ClientID       string             `bson:"client_id" json:"client_id"`
	ClientSecret   string             `bson:"client_secret" json:"client_secret"`
	TenantID       string             `bson:"tenant_id" json:"tenant_id"`
	SubscriptionID string             `bson:"subscription_id" json:"subscription_id"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

// AtlasCredential represents MongoDB Atlas credentials.
type AtlasCredential struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	PublicKey  string             `bson:"public_key" json:"public_key"`
	PrivateKey string             `bson:"private_key" json:"private_key"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// CloudflareCredential represents Cloudflare credentials.
type CloudflareCredential struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	AuthScheme int32              `bson:"auth_scheme" json:"auth_scheme"` // 1 = api_token, 2 = legacy_api_key
	APIToken   string             `bson:"api_token,omitempty" json:"api_token,omitempty"`
	APIKey     string             `bson:"api_key,omitempty" json:"api_key,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// ConfluentCredential represents Confluent Cloud credentials.
type ConfluentCredential struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	APIKey    string             `bson:"api_key" json:"api_key"`
	APISecret string             `bson:"api_secret" json:"api_secret"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// SnowflakeCredential represents Snowflake credentials.
type SnowflakeCredential struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Account   string             `bson:"account" json:"account"`
	Region    string             `bson:"region" json:"region"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// KubernetesCredential represents Kubernetes cluster credentials.
type KubernetesCredential struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Provider  int32              `bson:"provider" json:"provider"` // 1 = gcp_gke, 2 = aws_eks, 3 = azure_aks, 4 = digital_ocean_doks
	Spec      interface{}        `bson:"spec" json:"spec"`         // Provider-specific k8s config
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
