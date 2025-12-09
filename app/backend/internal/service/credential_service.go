package service

import (
	"context"
	"fmt"
	"time"

	"github.com/project-planton/project-planton/app/backend/internal/database"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CredentialService implements the CredentialService RPC.
type CredentialService struct {
	credentialRepo *database.CredentialRepository
}

// NewCredentialService creates a new service instance.
func NewCredentialService(credentialRepo *database.CredentialRepository) *CredentialService {
	return &CredentialService{
		credentialRepo: credentialRepo,
	}
}

// CreateCredential creates a new cloud provider credential.
func (s *CredentialService) CreateCredential(
	ctx context.Context,
	req *connect.Request[backendv1.CreateCredentialRequest],
) (*connect.Response[backendv1.CreateCredentialResponse], error) {
	// Validate common fields
	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("name is required"))
	}
	if req.Msg.Provider == backendv1.CredentialProvider_CREDENTIAL_PROVIDER_UNSPECIFIED {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("provider is required"))
	}

	now := time.Now()

	// Handle based on provider type
	switch req.Msg.Provider {
	case backendv1.CredentialProvider_GCP:
		return s.createGcpCredential(ctx, req.Msg.Name, req.Msg.GetGcp(), now)
	case backendv1.CredentialProvider_AWS:
		return s.createAwsCredential(ctx, req.Msg.Name, req.Msg.GetAws(), now)
	case backendv1.CredentialProvider_AZURE:
		return s.createAzureCredential(ctx, req.Msg.Name, req.Msg.GetAzure(), now)
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported provider: %v", req.Msg.Provider))
	}
}

// createGcpCredential creates a GCP credential.
func (s *CredentialService) createGcpCredential(
	ctx context.Context,
	name string,
	spec *backendv1.GcpCredentialSpec,
	now time.Time,
) (*connect.Response[backendv1.CreateCredentialResponse], error) {
	if spec == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("gcp credential spec is required"))
	}
	if spec.ServiceAccountKeyBase64 == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("service_account_key_base64 is required"))
	}

	createdCredential, err := s.credentialRepo.CreateGcp(ctx, name, spec.ServiceAccountKeyBase64)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create GCP credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       createdCredential.ID.Hex(),
		Name:     createdCredential.Name,
		Provider: backendv1.CredentialProvider_GCP,
		CredentialData: &backendv1.Credential_Gcp{
			Gcp: &backendv1.GcpCredentialSpec{
				ServiceAccountKeyBase64: createdCredential.ServiceAccountKeyBase64,
			},
		},
	}

	if !createdCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(createdCredential.CreatedAt)
	}
	if !createdCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(createdCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.CreateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// createAwsCredential creates an AWS credential.
func (s *CredentialService) createAwsCredential(
	ctx context.Context,
	name string,
	spec *backendv1.AwsCredentialSpec,
	now time.Time,
) (*connect.Response[backendv1.CreateCredentialResponse], error) {
	if spec == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("aws credential spec is required"))
	}
	if spec.AccountId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("account_id is required"))
	}
	if spec.AccessKeyId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("access_key_id is required"))
	}
	if spec.SecretAccessKey == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("secret_access_key is required"))
	}

	region := ""
	if spec.Region != nil {
		region = *spec.Region
	}
	sessionToken := ""
	if spec.SessionToken != nil {
		sessionToken = *spec.SessionToken
	}

	createdCredential, err := s.credentialRepo.CreateAws(ctx, name, spec.AccountId, spec.AccessKeyId, spec.SecretAccessKey, region, sessionToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create AWS credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       createdCredential.ID.Hex(),
		Name:     createdCredential.Name,
		Provider: backendv1.CredentialProvider_AWS,
		CredentialData: &backendv1.Credential_Aws{
			Aws: &backendv1.AwsCredentialSpec{
				AccountId:       createdCredential.AccountID,
				AccessKeyId:     createdCredential.AccessKeyID,
				SecretAccessKey: createdCredential.SecretAccessKey,
			},
		},
	}

	if createdCredential.Region != "" {
		region := createdCredential.Region
		protoCredential.GetAws().Region = &region
	}
	if createdCredential.SessionToken != "" {
		sessionToken := createdCredential.SessionToken
		protoCredential.GetAws().SessionToken = &sessionToken
	}

	if !createdCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(createdCredential.CreatedAt)
	}
	if !createdCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(createdCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.CreateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// createAzureCredential creates an Azure credential.
func (s *CredentialService) createAzureCredential(
	ctx context.Context,
	name string,
	spec *backendv1.AzureCredentialSpec,
	now time.Time,
) (*connect.Response[backendv1.CreateCredentialResponse], error) {
	if spec == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("azure credential spec is required"))
	}
	if spec.ClientId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("client_id is required"))
	}
	if spec.ClientSecret == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("client_secret is required"))
	}
	if spec.TenantId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("tenant_id is required"))
	}
	if spec.SubscriptionId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("subscription_id is required"))
	}

	createdCredential, err := s.credentialRepo.CreateAzure(ctx, name, spec.ClientId, spec.ClientSecret, spec.TenantId, spec.SubscriptionId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create Azure credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       createdCredential.ID.Hex(),
		Name:     createdCredential.Name,
		Provider: backendv1.CredentialProvider_AZURE,
		CredentialData: &backendv1.Credential_Azure{
			Azure: &backendv1.AzureCredentialSpec{
				ClientId:       createdCredential.ClientID,
				ClientSecret:   createdCredential.ClientSecret,
				TenantId:       createdCredential.TenantID,
				SubscriptionId: createdCredential.SubscriptionID,
			},
		},
	}

	if !createdCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(createdCredential.CreatedAt)
	}
	if !createdCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(createdCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.CreateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// ListCredentials lists all credentials with optional provider filter.
func (s *CredentialService) ListCredentials(
	ctx context.Context,
	req *connect.Request[backendv1.ListCredentialsRequest],
) (*connect.Response[backendv1.ListCredentialsResponse], error) {
	// Convert provider enum to string for database query
	var providerFilter *string
	if req.Msg.Provider != nil && *req.Msg.Provider != backendv1.CredentialProvider_CREDENTIAL_PROVIDER_UNSPECIFIED {
		// Convert CredentialProvider enum to string
		var provider string
		switch *req.Msg.Provider {
		case backendv1.CredentialProvider_GCP:
			provider = "gcp"
		case backendv1.CredentialProvider_AWS:
			provider = "aws"
		case backendv1.CredentialProvider_AZURE:
			provider = "azure"
		}
		if provider != "" {
			providerFilter = &provider
		}
	}

	// Query database
	credentials, err := s.credentialRepo.List(ctx, providerFilter)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list credentials: %w", err))
	}

	// Convert to proto summaries (without sensitive data)
	summaries := make([]*backendv1.CredentialSummary, 0, len(credentials))
	for _, cred := range credentials {
		summary := &backendv1.CredentialSummary{
			Id:   cred["_id"].(primitive.ObjectID).Hex(),
			Name: cred["name"].(string),
		}

		// Convert provider string to enum
		providerStr := cred["provider"].(string)
		switch providerStr {
		case "gcp":
			summary.Provider = backendv1.CredentialProvider_GCP
		case "aws":
			summary.Provider = backendv1.CredentialProvider_AWS
		case "azure":
			summary.Provider = backendv1.CredentialProvider_AZURE
		}

		// Add timestamps if present
		if createdAt, ok := cred["created_at"].(primitive.DateTime); ok {
			summary.CreatedAt = timestamppb.New(createdAt.Time())
		}
		if updatedAt, ok := cred["updated_at"].(primitive.DateTime); ok {
			summary.UpdatedAt = timestamppb.New(updatedAt.Time())
		}

		summaries = append(summaries, summary)
	}

	return connect.NewResponse(&backendv1.ListCredentialsResponse{
		Credentials: summaries,
	}), nil
}

