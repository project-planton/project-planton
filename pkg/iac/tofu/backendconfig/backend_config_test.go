package backendconfig

import (
	"testing"

	awsvpcv1 "github.com/plantonhq/project-planton/apis/org/project_planton/provider/aws/awsvpc/v1"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared"
	"github.com/plantonhq/project-planton/pkg/iac/tofu/tofulabels"
	"github.com/stretchr/testify/assert"
)

func TestExtractFromManifest(t *testing.T) {
	tests := []struct {
		name      string
		manifest  *awsvpcv1.AwsVpc
		want      *TofuBackendConfig
		wantError bool
		errorMsg  string
	}{
		{
			name: "valid s3 backend",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "s3",
						tofulabels.BackendObjectLabelKey: "my-terraform-state/aws-vpc/dev",
					},
				},
			},
			want: &TofuBackendConfig{
				BackendType:   "s3",
				BackendObject: "my-terraform-state/aws-vpc/dev",
			},
			wantError: false,
		},
		{
			name: "valid gcs backend",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "gcs",
						tofulabels.BackendObjectLabelKey: "my-gcs-bucket/terraform/state",
					},
				},
			},
			want: &TofuBackendConfig{
				BackendType:   "gcs",
				BackendObject: "my-gcs-bucket/terraform/state",
			},
			wantError: false,
		},
		{
			name: "valid azurerm backend",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "azurerm",
						tofulabels.BackendObjectLabelKey: "my-container/terraform/state",
					},
				},
			},
			want: &TofuBackendConfig{
				BackendType:   "azurerm",
				BackendObject: "my-container/terraform/state",
			},
			wantError: false,
		},
		{
			name: "valid local backend",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "local",
						tofulabels.BackendObjectLabelKey: "/tmp/terraform.tfstate",
					},
				},
			},
			want: &TofuBackendConfig{
				BackendType:   "local",
				BackendObject: "/tmp/terraform.tfstate",
			},
			wantError: false,
		},
		{
			name: "no backend labels - returns nil without error",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						"other.label": "value",
					},
				},
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "missing backend object",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey: "s3",
						// Missing backend object
					},
				},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "both",
		},
		{
			name: "missing backend type",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendObjectLabelKey: "my-bucket/state",
						// Missing backend type
					},
				},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "both",
		},
		{
			name: "empty backend type",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "",
						tofulabels.BackendObjectLabelKey: "my-bucket/state",
					},
				},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "cannot be empty",
		},
		{
			name: "empty backend object",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "s3",
						tofulabels.BackendObjectLabelKey: "",
					},
				},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "cannot be empty",
		},
		{
			name: "unsupported backend type",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{
					Labels: map[string]string{
						tofulabels.BackendTypeLabelKey:   "unsupported",
						tofulabels.BackendObjectLabelKey: "some/path",
					},
				},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "unsupported backend type",
		},
		{
			name: "no labels",
			manifest: &awsvpcv1.AwsVpc{
				Metadata: &shared.CloudResourceMetadata{},
			},
			want:      nil,
			wantError: true,
			errorMsg:  "no labels found in manifest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFromManifest(tt.manifest)

			if tt.wantError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
