package service

import (
	"context"
	"fmt"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
)

// DeploymentComponentService implements the DeploymentComponentService RPC.
type DeploymentComponentService struct {
	repo *database.DeploymentComponentRepository
}

// NewDeploymentComponentService creates a new service instance.
func NewDeploymentComponentService(repo *database.DeploymentComponentRepository) *DeploymentComponentService {
	return &DeploymentComponentService{
		repo: repo,
	}
}

// ListDeploymentComponents retrieves a list of deployment components with optional filters.
func (s *DeploymentComponentService) ListDeploymentComponents(
	ctx context.Context,
	req *connect.Request[backendv1.ListDeploymentComponentsRequest],
) (*connect.Response[backendv1.ListDeploymentComponentsResponse], error) {
	logrus.WithFields(logrus.Fields{
		"provider": req.Msg.Provider,
		"kind":     req.Msg.Kind,
	}).Info("Listing deployment components")

	opts := &database.ListOptions{}
	if req.Msg.Provider != nil {
		provider := *req.Msg.Provider
		opts.Provider = &provider
	}
	if req.Msg.Kind != nil {
		kind := *req.Msg.Kind
		opts.Kind = &kind
	}

	components, err := s.repo.List(ctx, opts)
	if err != nil {
		logrus.WithError(err).Error("Failed to list deployment components")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list deployment components: %w", err))
	}

	protoComponents := make([]*backendv1.DeploymentComponent, 0, len(components))
	for _, comp := range components {
		protoComp := &backendv1.DeploymentComponent{
			Id:            comp.ID.Hex(),
			Kind:          comp.Kind,
			Provider:      comp.Provider,
			Name:          comp.Name,
			Version:       comp.Version,
			IdPrefix:      comp.IDPrefix,
			IsServiceKind: comp.IsServiceKind,
		}

		if !comp.CreatedAt.IsZero() {
			protoComp.CreatedAt = timestamppb.New(comp.CreatedAt)
		}
		if !comp.UpdatedAt.IsZero() {
			protoComp.UpdatedAt = timestamppb.New(comp.UpdatedAt)
		}

		protoComponents = append(protoComponents, protoComp)
	}

	return connect.NewResponse(&backendv1.ListDeploymentComponentsResponse{
		Components: protoComponents,
	}), nil
}
