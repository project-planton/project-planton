package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/pkg/crkreflect"
)

// CredentialResolver resolves provider credentials from the database based on provider.
type CredentialResolver struct {
	awsCredentialRepo   *database.AwsCredentialRepository
	gcpCredentialRepo   *database.GcpCredentialRepository
	azureCredentialRepo *database.AzureCredentialRepository
}

// NewCredentialResolver creates a new credential resolver instance.
func NewCredentialResolver(
	awsCredentialRepo *database.AwsCredentialRepository,
	gcpCredentialRepo *database.GcpCredentialRepository,
	azureCredentialRepo *database.AzureCredentialRepository,
) *CredentialResolver {
	return &CredentialResolver{
		awsCredentialRepo:   awsCredentialRepo,
		gcpCredentialRepo:   gcpCredentialRepo,
		azureCredentialRepo: azureCredentialRepo,
	}
}

// ResolveProviderConfig resolves provider credentials from the database based on the provider from cloud resource kind.
// Returns a ProviderConfig proto message that can be used for deployment.
func (r *CredentialResolver) ResolveProviderConfig(
	ctx context.Context,
	kindName string,
) (*backendv1.ProviderConfig, error) {
	// Step 1: Get the CloudResourceKind enum from kind name
	kindEnum, err := crkreflect.KindByKindName(kindName)
	if err != nil {
		return nil, fmt.Errorf("failed to get kind enum for '%s': %w", kindName, err)
	}

	// Step 2: Get the provider from the kind
	providerEnum := crkreflect.GetProvider(kindEnum)
	if providerEnum == cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified {
		return nil, fmt.Errorf("provider not configured for cloud resource kind '%s'", kindName)
	}

	// Step 3: Convert provider enum to string (e.g., "aws", "gcp", "azure")
	providerString := providerEnumToString(providerEnum)
	if providerString == "" {
		return nil, fmt.Errorf("unsupported provider: %v", providerEnum)
	}

	// Step 4: Fetch the first credential from the appropriate repository based on provider
	switch providerString {
	case "aws":
		awsCred, err := r.awsCredentialRepo.FindFirst(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch AWS credential: %w", err)
		}
		if awsCred == nil {
			return nil, fmt.Errorf("no AWS credential found. Please create an AWS credential first")
		}
		return &backendv1.ProviderConfig{
			Config: &backendv1.ProviderConfig_Aws{
				Aws: &backendv1.AwsProviderConfig{
					AccountId:       awsCred.AccountID,
					AccessKeyId:     awsCred.AccessKeyID,
					SecretAccessKey: awsCred.SecretAccessKey,
					Region:          &awsCred.Region,
					SessionToken:    &awsCred.SessionToken,
				},
			},
		}, nil

	case "gcp":
		gcpCred, err := r.gcpCredentialRepo.FindFirst(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch GCP credential: %w", err)
		}
		if gcpCred == nil {
			return nil, fmt.Errorf("no GCP credential found. Please create a GCP credential first")
		}
		return &backendv1.ProviderConfig{
			Config: &backendv1.ProviderConfig_Gcp{
				Gcp: &backendv1.GcpProviderConfig{
					ServiceAccountKeyBase64: gcpCred.ServiceAccountKeyBase64,
				},
			},
		}, nil

	case "azure":
		azureCred, err := r.azureCredentialRepo.FindFirst(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch Azure credential: %w", err)
		}
		if azureCred == nil {
			return nil, fmt.Errorf("no Azure credential found. Please create an Azure credential first")
		}
		return &backendv1.ProviderConfig{
			Config: &backendv1.ProviderConfig_Azure{
				Azure: &backendv1.AzureProviderConfig{
					ClientId:       azureCred.ClientID,
					ClientSecret:   azureCred.ClientSecret,
					TenantId:       azureCred.TenantID,
					SubscriptionId: azureCred.SubscriptionID,
				},
			},
		}, nil

	default:
		return nil, fmt.Errorf("provider '%s' is not yet supported for automatic credential resolution", providerString)
	}
}

// providerEnumToString converts CloudResourceProvider enum to a lowercase string.
// The enum String() method returns values like "aws", "gcp", "azure", etc.
func providerEnumToString(provider cloudresourcekind.CloudResourceProvider) string {
	name := provider.String()
	// Handle special case for unspecified
	if name == "cloud_resource_provider_unspecified" {
		return ""
	}
	// For other values, String() returns the enum name directly (e.g., "aws", "gcp")
	return strings.ToLower(name)
}
