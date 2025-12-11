package service

import (
	"context"
	"fmt"
	"time"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"go.mongodb.org/mongo-driver/bson"
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

// GetCredential retrieves a credential by ID.
func (s *CredentialService) GetCredential(
	ctx context.Context,
	req *connect.Request[backendv1.GetCredentialRequest],
) (*connect.Response[backendv1.GetCredentialResponse], error) {
	if req.Msg.Id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("id is required"))
	}

	doc, err := s.credentialRepo.FindByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get credential: %w", err))
	}
	if doc == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("credential with ID '%s' not found", req.Msg.Id))
	}

	// Convert to proto based on provider
	providerStr := doc["provider"].(string)
	var protoCredential *backendv1.Credential

	switch providerStr {
	case "gcp":
		gcpCred, err := convertBsonToGcpCredential(doc)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert credential: %w", err))
		}
		protoCredential = &backendv1.Credential{
			Id:       gcpCred.ID.Hex(),
			Name:     gcpCred.Name,
			Provider: backendv1.CredentialProvider_GCP,
			CredentialData: &backendv1.Credential_Gcp{
				Gcp: &backendv1.GcpCredentialSpec{
					ServiceAccountKeyBase64: gcpCred.ServiceAccountKeyBase64,
				},
			},
		}
		if !gcpCred.CreatedAt.IsZero() {
			protoCredential.CreatedAt = timestamppb.New(gcpCred.CreatedAt)
		}
		if !gcpCred.UpdatedAt.IsZero() {
			protoCredential.UpdatedAt = timestamppb.New(gcpCred.UpdatedAt)
		}
	case "aws":
		awsCred, err := convertBsonToAwsCredential(doc)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert credential: %w", err))
		}
		protoCredential = &backendv1.Credential{
			Id:       awsCred.ID.Hex(),
			Name:     awsCred.Name,
			Provider: backendv1.CredentialProvider_AWS,
			CredentialData: &backendv1.Credential_Aws{
				Aws: &backendv1.AwsCredentialSpec{
					AccountId:       awsCred.AccountID,
					AccessKeyId:     awsCred.AccessKeyID,
					SecretAccessKey: awsCred.SecretAccessKey,
				},
			},
		}
		if awsCred.Region != "" {
			region := awsCred.Region
			protoCredential.GetAws().Region = &region
		}
		if awsCred.SessionToken != "" {
			sessionToken := awsCred.SessionToken
			protoCredential.GetAws().SessionToken = &sessionToken
		}
		if !awsCred.CreatedAt.IsZero() {
			protoCredential.CreatedAt = timestamppb.New(awsCred.CreatedAt)
		}
		if !awsCred.UpdatedAt.IsZero() {
			protoCredential.UpdatedAt = timestamppb.New(awsCred.UpdatedAt)
		}
	case "azure":
		azureCred, err := convertBsonToAzureCredential(doc)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert credential: %w", err))
		}
		protoCredential = &backendv1.Credential{
			Id:       azureCred.ID.Hex(),
			Name:     azureCred.Name,
			Provider: backendv1.CredentialProvider_AZURE,
			CredentialData: &backendv1.Credential_Azure{
				Azure: &backendv1.AzureCredentialSpec{
					ClientId:       azureCred.ClientID,
					ClientSecret:   azureCred.ClientSecret,
					TenantId:       azureCred.TenantID,
					SubscriptionId: azureCred.SubscriptionID,
				},
			},
		}
		if !azureCred.CreatedAt.IsZero() {
			protoCredential.CreatedAt = timestamppb.New(azureCred.CreatedAt)
		}
		if !azureCred.UpdatedAt.IsZero() {
			protoCredential.UpdatedAt = timestamppb.New(azureCred.UpdatedAt)
		}
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported provider: %s", providerStr))
	}

	return connect.NewResponse(&backendv1.GetCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// UpdateCredential updates an existing credential.
func (s *CredentialService) UpdateCredential(
	ctx context.Context,
	req *connect.Request[backendv1.UpdateCredentialRequest],
) (*connect.Response[backendv1.UpdateCredentialResponse], error) {
	if req.Msg.Id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("id is required"))
	}
	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("name is required"))
	}
	if req.Msg.Provider == backendv1.CredentialProvider_CREDENTIAL_PROVIDER_UNSPECIFIED {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("provider is required"))
	}

	// Handle based on provider type
	switch req.Msg.Provider {
	case backendv1.CredentialProvider_GCP:
		return s.updateGcpCredential(ctx, req.Msg.Id, req.Msg.Name, req.Msg.GetGcp())
	case backendv1.CredentialProvider_AWS:
		return s.updateAwsCredential(ctx, req.Msg.Id, req.Msg.Name, req.Msg.GetAws())
	case backendv1.CredentialProvider_AZURE:
		return s.updateAzureCredential(ctx, req.Msg.Id, req.Msg.Name, req.Msg.GetAzure())
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unsupported provider: %v", req.Msg.Provider))
	}
}

// updateGcpCredential updates a GCP credential.
func (s *CredentialService) updateGcpCredential(
	ctx context.Context,
	id, name string,
	spec *backendv1.GcpCredentialSpec,
) (*connect.Response[backendv1.UpdateCredentialResponse], error) {
	if spec == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("gcp credential spec is required"))
	}
	if spec.ServiceAccountKeyBase64 == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("service_account_key_base64 is required"))
	}

	updatedCredential, err := s.credentialRepo.UpdateGcp(ctx, id, name, spec.ServiceAccountKeyBase64)
	if err != nil {
		if err.Error() == fmt.Sprintf("GCP credential with ID '%s' not found", id) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update GCP credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       updatedCredential.ID.Hex(),
		Name:     updatedCredential.Name,
		Provider: backendv1.CredentialProvider_GCP,
		CredentialData: &backendv1.Credential_Gcp{
			Gcp: &backendv1.GcpCredentialSpec{
				ServiceAccountKeyBase64: updatedCredential.ServiceAccountKeyBase64,
			},
		},
	}

	if !updatedCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(updatedCredential.CreatedAt)
	}
	if !updatedCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(updatedCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.UpdateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// updateAwsCredential updates an AWS credential.
func (s *CredentialService) updateAwsCredential(
	ctx context.Context,
	id, name string,
	spec *backendv1.AwsCredentialSpec,
) (*connect.Response[backendv1.UpdateCredentialResponse], error) {
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

	updatedCredential, err := s.credentialRepo.UpdateAws(ctx, id, name, spec.AccountId, spec.AccessKeyId, spec.SecretAccessKey, region, sessionToken)
	if err != nil {
		if err.Error() == fmt.Sprintf("AWS credential with ID '%s' not found", id) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update AWS credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       updatedCredential.ID.Hex(),
		Name:     updatedCredential.Name,
		Provider: backendv1.CredentialProvider_AWS,
		CredentialData: &backendv1.Credential_Aws{
			Aws: &backendv1.AwsCredentialSpec{
				AccountId:       updatedCredential.AccountID,
				AccessKeyId:     updatedCredential.AccessKeyID,
				SecretAccessKey: updatedCredential.SecretAccessKey,
			},
		},
	}

	if updatedCredential.Region != "" {
		region := updatedCredential.Region
		protoCredential.GetAws().Region = &region
	}
	if updatedCredential.SessionToken != "" {
		sessionToken := updatedCredential.SessionToken
		protoCredential.GetAws().SessionToken = &sessionToken
	}

	if !updatedCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(updatedCredential.CreatedAt)
	}
	if !updatedCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(updatedCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.UpdateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// updateAzureCredential updates an Azure credential.
func (s *CredentialService) updateAzureCredential(
	ctx context.Context,
	id, name string,
	spec *backendv1.AzureCredentialSpec,
) (*connect.Response[backendv1.UpdateCredentialResponse], error) {
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

	updatedCredential, err := s.credentialRepo.UpdateAzure(ctx, id, name, spec.ClientId, spec.ClientSecret, spec.TenantId, spec.SubscriptionId)
	if err != nil {
		if err.Error() == fmt.Sprintf("Azure credential with ID '%s' not found", id) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update Azure credential: %w", err))
	}

	protoCredential := &backendv1.Credential{
		Id:       updatedCredential.ID.Hex(),
		Name:     updatedCredential.Name,
		Provider: backendv1.CredentialProvider_AZURE,
		CredentialData: &backendv1.Credential_Azure{
			Azure: &backendv1.AzureCredentialSpec{
				ClientId:       updatedCredential.ClientID,
				ClientSecret:   updatedCredential.ClientSecret,
				TenantId:       updatedCredential.TenantID,
				SubscriptionId: updatedCredential.SubscriptionID,
			},
		},
	}

	if !updatedCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(updatedCredential.CreatedAt)
	}
	if !updatedCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(updatedCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.UpdateCredentialResponse{
		Credential: protoCredential,
	}), nil
}

// DeleteCredential deletes a credential by ID.
func (s *CredentialService) DeleteCredential(
	ctx context.Context,
	req *connect.Request[backendv1.DeleteCredentialRequest],
) (*connect.Response[backendv1.DeleteCredentialResponse], error) {
	if req.Msg.Id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("id is required"))
	}

	err := s.credentialRepo.Delete(ctx, req.Msg.Id)
	if err != nil {
		if err.Error() == fmt.Sprintf("credential with ID '%s' not found", req.Msg.Id) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to delete credential: %w", err))
	}

	return connect.NewResponse(&backendv1.DeleteCredentialResponse{
		Message: fmt.Sprintf("Credential with ID '%s' deleted successfully", req.Msg.Id),
	}), nil
}

// Helper functions to convert bson.M to typed credentials
func convertBsonToGcpCredential(doc bson.M) (*models.GcpCredential, error) {
	id, ok := doc["_id"].(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid _id field")
	}

	var createdAt, updatedAt time.Time
	if dt, ok := doc["created_at"].(primitive.DateTime); ok {
		createdAt = dt.Time()
	} else if t, ok := doc["created_at"].(time.Time); ok {
		createdAt = t
	}
	if dt, ok := doc["updated_at"].(primitive.DateTime); ok {
		updatedAt = dt.Time()
	} else if t, ok := doc["updated_at"].(time.Time); ok {
		updatedAt = t
	}

	return &models.GcpCredential{
		ID:                      id,
		Name:                    doc["name"].(string),
		ServiceAccountKeyBase64: doc["service_account_key_base64"].(string),
		CreatedAt:               createdAt,
		UpdatedAt:               updatedAt,
	}, nil
}

func convertBsonToAwsCredential(doc bson.M) (*models.AwsCredential, error) {
	id, ok := doc["_id"].(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid _id field")
	}

	var createdAt, updatedAt time.Time
	if dt, ok := doc["created_at"].(primitive.DateTime); ok {
		createdAt = dt.Time()
	} else if t, ok := doc["created_at"].(time.Time); ok {
		createdAt = t
	}
	if dt, ok := doc["updated_at"].(primitive.DateTime); ok {
		updatedAt = dt.Time()
	} else if t, ok := doc["updated_at"].(time.Time); ok {
		updatedAt = t
	}

	cred := &models.AwsCredential{
		ID:              id,
		Name:            doc["name"].(string),
		AccountID:       doc["account_id"].(string),
		AccessKeyID:     doc["access_key_id"].(string),
		SecretAccessKey: doc["secret_access_key"].(string),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}

	if region, ok := doc["region"].(string); ok {
		cred.Region = region
	}
	if sessionToken, ok := doc["session_token"].(string); ok {
		cred.SessionToken = sessionToken
	}

	return cred, nil
}

func convertBsonToAzureCredential(doc bson.M) (*models.AzureCredential, error) {
	id, ok := doc["_id"].(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid _id field")
	}

	var createdAt, updatedAt time.Time
	if dt, ok := doc["created_at"].(primitive.DateTime); ok {
		createdAt = dt.Time()
	} else if t, ok := doc["created_at"].(time.Time); ok {
		createdAt = t
	}
	if dt, ok := doc["updated_at"].(primitive.DateTime); ok {
		updatedAt = dt.Time()
	} else if t, ok := doc["updated_at"].(time.Time); ok {
		updatedAt = t
	}

	return &models.AzureCredential{
		ID:             id,
		Name:           doc["name"].(string),
		ClientID:       doc["client_id"].(string),
		ClientSecret:   doc["client_secret"].(string),
		TenantID:       doc["tenant_id"].(string),
		SubscriptionID: doc["subscription_id"].(string),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}
