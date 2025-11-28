package service

import (
	"context"
	"fmt"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CloudResourceService implements the CloudResourceService RPC.
type CloudResourceService struct {
	repo *database.CloudResourceRepository
}

// NewCloudResourceService creates a new service instance.
func NewCloudResourceService(repo *database.CloudResourceRepository) *CloudResourceService {
	return &CloudResourceService{
		repo: repo,
	}
}

// CreateCloudResource creates a new cloud resource from a YAML manifest.
func (s *CloudResourceService) CreateCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.CreateCloudResourceRequest],
) (*connect.Response[backendv1.CreateCloudResourceResponse], error) {
	manifest := req.Msg.Manifest
	if manifest == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest cannot be empty"))
	}

	// Parse YAML to extract kind and metadata.name
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal([]byte(manifest), &yamlData); err != nil {
		logrus.WithError(err).Error("Failed to parse YAML manifest")
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid YAML format: %w", err))
	}

	// Extract kind
	kind, ok := yamlData["kind"].(string)
	if !ok || kind == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest must contain 'kind' field"))
	}

	// Extract metadata.name
	metadata, ok := yamlData["metadata"].(map[string]interface{})
	if !ok {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest must contain 'metadata' field"))
	}

	name, ok := metadata["name"].(string)
	if !ok || name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest must contain 'metadata.name' field"))
	}

	logrus.WithFields(logrus.Fields{
		"name": name,
		"kind": kind,
	}).Info("Creating cloud resource")

	// Check if a resource with the same name already exists
	existingResource, err := s.repo.FindByName(ctx, name)
	if err != nil {
		logrus.WithError(err).Error("Failed to check for existing cloud resource")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to validate cloud resource name: %w", err))
	}
	if existingResource != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("cloud resource with name '%s' already exists", name))
	}

	// Create domain model
	cloudResource := &models.CloudResource{
		Name:     name,
		Kind:     kind,
		Manifest: manifest,
	}

	// Save to database
	createdResource, err := s.repo.Create(ctx, cloudResource)
	if err != nil {
		logrus.WithError(err).Error("Failed to create cloud resource")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create cloud resource: %w", err))
	}

	// Convert to proto
	protoResource := &backendv1.CloudResource{
		Id:       createdResource.ID.Hex(),
		Name:     createdResource.Name,
		Kind:     createdResource.Kind,
		Manifest: createdResource.Manifest,
	}

	if !createdResource.CreatedAt.IsZero() {
		protoResource.CreatedAt = timestamppb.New(createdResource.CreatedAt)
	}
	if !createdResource.UpdatedAt.IsZero() {
		protoResource.UpdatedAt = timestamppb.New(createdResource.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.CreateCloudResourceResponse{
		Resource: protoResource,
	}), nil
}

// ListCloudResources retrieves all cloud resources.
func (s *CloudResourceService) ListCloudResources(
	ctx context.Context,
	req *connect.Request[backendv1.ListCloudResourcesRequest],
) (*connect.Response[backendv1.ListCloudResourcesResponse], error) {
	opts := &database.CloudResourceListOptions{}
	if req.Msg.Kind != nil {
		kind := *req.Msg.Kind
		opts.Kind = &kind
	}

	logrus.WithFields(logrus.Fields{
		"kind": req.Msg.Kind,
	}).Info("Listing cloud resources")

	resources, err := s.repo.List(ctx, opts)
	if err != nil {
		logrus.WithError(err).Error("Failed to list cloud resources")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list cloud resources: %w", err))
	}

	protoResources := make([]*backendv1.CloudResource, 0, len(resources))
	for _, res := range resources {
		protoRes := &backendv1.CloudResource{
			Id:       res.ID.Hex(),
			Name:     res.Name,
			Kind:     res.Kind,
			Manifest: res.Manifest,
		}

		if !res.CreatedAt.IsZero() {
			protoRes.CreatedAt = timestamppb.New(res.CreatedAt)
		}
		if !res.UpdatedAt.IsZero() {
			protoRes.UpdatedAt = timestamppb.New(res.UpdatedAt)
		}

		protoResources = append(protoResources, protoRes)
	}

	return connect.NewResponse(&backendv1.ListCloudResourcesResponse{
		Resources: protoResources,
	}), nil
}

