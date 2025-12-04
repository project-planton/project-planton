package service

import (
	"context"
	"fmt"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"
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

	// Check if a resource with the same name already exists
	existingResource, err := s.repo.FindByName(ctx, name)
	if err != nil {
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

// ListCloudResources retrieves cloud resources with optional pagination.
func (s *CloudResourceService) ListCloudResources(
	ctx context.Context,
	req *connect.Request[backendv1.ListCloudResourcesRequest],
) (*connect.Response[backendv1.ListCloudResourcesResponse], error) {
	opts := &database.CloudResourceListOptions{}
	if req.Msg.Kind != nil {
		kind := *req.Msg.Kind
		opts.Kind = &kind
	}

	// Apply pagination with defaults (page=0, size=20) if not provided
	var pageNum int32 = 0
	var pageSize int32 = 20
	if req.Msg.PageInfo != nil {
		pageNum = req.Msg.PageInfo.Num
		pageSize = req.Msg.PageInfo.Size
	}
	opts.PageNum = &pageNum
	opts.PageSize = &pageSize

	// Calculate total pages
	totalCount, err := s.repo.Count(ctx, opts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to count cloud resources: %w", err))
	}

	var totalPages int32
	if pageSize > 0 {
		totalPages = int32((totalCount + int64(pageSize) - 1) / int64(pageSize))
	}

	resources, err := s.repo.List(ctx, opts)
	if err != nil {
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

	response := &backendv1.ListCloudResourcesResponse{
		Resources:  protoResources,
		TotalPages: totalPages,
	}

	return connect.NewResponse(response), nil
}

// GetCloudResource retrieves a cloud resource by ID.
func (s *CloudResourceService) GetCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.GetCloudResourceRequest],
) (*connect.Response[backendv1.GetCloudResourceResponse], error) {
	id := req.Msg.Id
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("resource ID cannot be empty"))
	}

	resource, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get cloud resource: %w", err))
	}

	if resource == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("cloud resource with ID '%s' not found", id))
	}

	protoResource := &backendv1.CloudResource{
		Id:       resource.ID.Hex(),
		Name:     resource.Name,
		Kind:     resource.Kind,
		Manifest: resource.Manifest,
	}

	if !resource.CreatedAt.IsZero() {
		protoResource.CreatedAt = timestamppb.New(resource.CreatedAt)
	}
	if !resource.UpdatedAt.IsZero() {
		protoResource.UpdatedAt = timestamppb.New(resource.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.GetCloudResourceResponse{
		Resource: protoResource,
	}), nil
}

// UpdateCloudResource updates an existing cloud resource from a YAML manifest.
func (s *CloudResourceService) UpdateCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.UpdateCloudResourceRequest],
) (*connect.Response[backendv1.UpdateCloudResourceResponse], error) {
	id := req.Msg.Id
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("resource ID cannot be empty"))
	}

	manifest := req.Msg.Manifest
	if manifest == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest cannot be empty"))
	}

	// Check if resource exists
	existingResource, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find cloud resource: %w", err))
	}
	if existingResource == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("cloud resource with ID '%s' not found", id))
	}

	// Parse YAML to extract kind and metadata.name
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal([]byte(manifest), &yamlData); err != nil {
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

	// Validate name and kind match existing resource
	if name != existingResource.Name {
		return nil, connect.NewError(connect.CodeInvalidArgument,
			fmt.Errorf("manifest name '%s' does not match existing resource name '%s'", name, existingResource.Name))
	}

	if kind != existingResource.Kind {
		return nil, connect.NewError(connect.CodeInvalidArgument,
			fmt.Errorf("manifest kind '%s' does not match existing resource kind '%s'", kind, existingResource.Kind))
	}

	// Create updated domain model
	cloudResource := &models.CloudResource{
		Name:     name,
		Kind:     kind,
		Manifest: manifest,
	}

	// Update in database
	updatedResource, err := s.repo.Update(ctx, id, cloudResource)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update cloud resource: %w", err))
	}

	// Convert to proto
	protoResource := &backendv1.CloudResource{
		Id:       updatedResource.ID.Hex(),
		Name:     updatedResource.Name,
		Kind:     updatedResource.Kind,
		Manifest: updatedResource.Manifest,
	}

	if !updatedResource.CreatedAt.IsZero() {
		protoResource.CreatedAt = timestamppb.New(updatedResource.CreatedAt)
	}
	if !updatedResource.UpdatedAt.IsZero() {
		protoResource.UpdatedAt = timestamppb.New(updatedResource.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.UpdateCloudResourceResponse{
		Resource: protoResource,
	}), nil
}

// DeleteCloudResource deletes a cloud resource by ID.
func (s *CloudResourceService) DeleteCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.DeleteCloudResourceRequest],
) (*connect.Response[backendv1.DeleteCloudResourceResponse], error) {
	id := req.Msg.Id
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("resource ID cannot be empty"))
	}

	// Check if resource exists first
	resource, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find cloud resource: %w", err))
	}
	if resource == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("cloud resource with ID '%s' not found", id))
	}

	// Delete from database
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to delete cloud resource: %w", err))
	}

	return connect.NewResponse(&backendv1.DeleteCloudResourceResponse{
		Message: fmt.Sprintf("Cloud resource '%s' deleted successfully", resource.Name),
	}), nil
}

// ApplyCloudResource creates or updates a cloud resource (upsert operation).
func (s *CloudResourceService) ApplyCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.ApplyCloudResourceRequest],
) (*connect.Response[backendv1.ApplyCloudResourceResponse], error) {
	manifest := req.Msg.Manifest
	if manifest == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("manifest cannot be empty"))
	}

	// Parse YAML to extract kind and metadata.name
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal([]byte(manifest), &yamlData); err != nil {
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

	// Check if resource exists by name and kind
	existingResource, err := s.repo.FindByNameAndKind(ctx, name, kind)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to check for existing cloud resource: %w", err))
	}

	var resultResource *models.CloudResource
	var created bool

	if existingResource != nil {
		// Resource exists - perform update

		cloudResource := &models.CloudResource{
			Name:     name,
			Kind:     kind,
			Manifest: manifest,
		}

		resultResource, err = s.repo.Update(ctx, existingResource.ID.Hex(), cloudResource)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update cloud resource: %w", err))
		}
		created = false
	} else {
		// Resource doesn't exist - perform create

		cloudResource := &models.CloudResource{
			Name:     name,
			Kind:     kind,
			Manifest: manifest,
		}

		resultResource, err = s.repo.Create(ctx, cloudResource)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create cloud resource: %w", err))
		}
		created = true
	}

	// Convert to proto
	protoResource := &backendv1.CloudResource{
		Id:       resultResource.ID.Hex(),
		Name:     resultResource.Name,
		Kind:     resultResource.Kind,
		Manifest: resultResource.Manifest,
	}

	if !resultResource.CreatedAt.IsZero() {
		protoResource.CreatedAt = timestamppb.New(resultResource.CreatedAt)
	}
	if !resultResource.UpdatedAt.IsZero() {
		protoResource.UpdatedAt = timestamppb.New(resultResource.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.ApplyCloudResourceResponse{
		Resource: protoResource,
		Created:  created,
	}), nil
}

// CountCloudResources returns the total count of cloud resources.
func (s *CloudResourceService) CountCloudResources(
	ctx context.Context,
	req *connect.Request[backendv1.CountCloudResourcesRequest],
) (*connect.Response[backendv1.CountCloudResourcesResponse], error) {
	opts := &database.CloudResourceListOptions{}
	if req.Msg.Kind != nil {
		kind := *req.Msg.Kind
		opts.Kind = &kind
	}

	count, err := s.repo.Count(ctx, opts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to count cloud resources: %w", err))
	}

	return connect.NewResponse(&backendv1.CountCloudResourcesResponse{
		Count: count,
	}), nil
}
