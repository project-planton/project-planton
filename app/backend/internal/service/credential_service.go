package service

import (
	"context"
	"fmt"
	"time"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CredentialService implements the CredentialService RPC.
type CredentialService struct {
	gcpCredentialRepo *database.GcpCredentialRepository
}

// NewCredentialService creates a new service instance.
func NewCredentialService(gcpCredentialRepo *database.GcpCredentialRepository) *CredentialService {
	return &CredentialService{
		gcpCredentialRepo: gcpCredentialRepo,
	}
}

// CreateGcpCredential creates a new GCP credential.
func (s *CredentialService) CreateGcpCredential(
	ctx context.Context,
	req *connect.Request[backendv1.CreateGcpCredentialRequest],
) (*connect.Response[backendv1.CreateGcpCredentialResponse], error) {
	// Validate request
	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("name is required"))
	}
	if req.Msg.ServiceAccountKeyBase64 == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("service_account_key_base64 is required"))
	}

	// Create credential model
	now := time.Now()
	credential := &models.GcpCredential{
		Name:                    req.Msg.Name,
		ServiceAccountKeyBase64: req.Msg.ServiceAccountKeyBase64,
		CreatedAt:               now,
		UpdatedAt:               now,
	}

	// Save to database
	createdCredential, err := s.gcpCredentialRepo.Create(ctx, credential)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create GCP credential: %w", err))
	}

	// Convert to proto
	protoCredential := &backendv1.GcpCredential{
		Id:                      createdCredential.ID.Hex(),
		Name:                    createdCredential.Name,
		ServiceAccountKeyBase64: createdCredential.ServiceAccountKeyBase64,
	}

	if !createdCredential.CreatedAt.IsZero() {
		protoCredential.CreatedAt = timestamppb.New(createdCredential.CreatedAt)
	}
	if !createdCredential.UpdatedAt.IsZero() {
		protoCredential.UpdatedAt = timestamppb.New(createdCredential.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.CreateGcpCredentialResponse{
		Credential: protoCredential,
	}), nil
}
