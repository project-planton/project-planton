package crkreflect

import (
	"testing"

	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
)

func TestKindFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected cloudresourcekind.CloudResourceKind
	}{
		// KubernetesMicroservice tests
		{
			name:     "KubernetesMicroservice - PascalCase",
			input:    "KubernetesMicroservice",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},
		{
			name:     "KubernetesMicroservice - kebab-case",
			input:    "kubernetes-microservice",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},
		{
			name:     "KubernetesMicroservice - snake_case",
			input:    "kubernetes_microservice",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},
		{
			name:     "KubernetesMicroservice - UPPERCASE",
			input:    "KUBERNETESMICROSERVICE",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},
		{
			name:     "KubernetesMicroservice - mixed case with hyphens",
			input:    "Kubernetes-Microservice",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},
		{
			name:     "KubernetesMicroservice - mixed case with underscores",
			input:    "Kubernetes_Microservice",
			expected: cloudresourcekind.CloudResourceKind_KubernetesMicroservice,
		},

		// AwsEcsService tests
		{
			name:     "AwsEcsService - PascalCase",
			input:    "AwsEcsService",
			expected: cloudresourcekind.CloudResourceKind_AwsEcsService,
		},
		{
			name:     "AwsEcsService - kebab-case",
			input:    "aws-ecs-service",
			expected: cloudresourcekind.CloudResourceKind_AwsEcsService,
		},
		{
			name:     "AwsEcsService - snake_case",
			input:    "aws_ecs_service",
			expected: cloudresourcekind.CloudResourceKind_AwsEcsService,
		},
		{
			name:     "AwsEcsService - UPPERCASE",
			input:    "AWSECSSERVICE",
			expected: cloudresourcekind.CloudResourceKind_AwsEcsService,
		},

		// GcpGkeCluster tests
		{
			name:     "GcpGkeCluster - PascalCase",
			input:    "GcpGkeCluster",
			expected: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
		},
		{
			name:     "GcpGkeCluster - kebab-case",
			input:    "gcp-gke-cluster",
			expected: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
		},
		{
			name:     "GcpGkeCluster - snake_case",
			input:    "gcp_gke_cluster",
			expected: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
		},

		// Invalid input tests
		{
			name:     "Empty string",
			input:    "",
			expected: cloudresourcekind.CloudResourceKind_unspecified,
		},
		{
			name:     "Unknown resource",
			input:    "UnknownResource",
			expected: cloudresourcekind.CloudResourceKind_unspecified,
		},
		{
			name:     "Random string",
			input:    "random-string-123",
			expected: cloudresourcekind.CloudResourceKind_unspecified,
		},
	}

	// Also test if aliases in AliasMap work (if any are defined)
	// Note: This test assumes AliasMap might be populated elsewhere
	if len(AliasMap) > 0 {
		for kind, aliases := range AliasMap {
			for _, alias := range aliases {
				tests = append(tests, struct {
					name     string
					input    string
					expected cloudresourcekind.CloudResourceKind
				}{
					name:     "Alias for " + kind.String() + ": " + alias,
					input:    alias,
					expected: kind,
				})
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KindFromString(tt.input)
			if result != tt.expected {
				t.Errorf("KindFromString(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestKindFromStringNormalization specifically tests the normalization logic
func TestKindFromStringNormalization(t *testing.T) {
	// Test that all different representations of the same kind resolve to the same value
	testCases := [][]string{
		{
			"KubernetesMicroservice",
			"kubernetesmicroservice",
			"KUBERNETESMICROSERVICE",
			"kubernetes-microservice",
			"kubernetes_microservice",
			"Kubernetes-Microservice",
			"Kubernetes_Microservice",
			"KUBERNETES-MICROSERVICE",
			"KUBERNETES_MICROSERVICE",
		},
		{
			"AwsEcsService",
			"awsecsservice",
			"AWSECSSERVICE",
			"aws-ecs-service",
			"aws_ecs_service",
			"Aws-Ecs-Service",
			"AWS_ECS_SERVICE",
		},
	}

	for _, variants := range testCases {
		if len(variants) == 0 {
			continue
		}

		// Get the expected kind from the first variant (PascalCase)
		expected := KindFromString(variants[0])
		if expected == cloudresourcekind.CloudResourceKind_unspecified {
			t.Errorf("Failed to get expected kind for %s", variants[0])
			continue
		}

		// Test that all variants resolve to the same kind
		for _, variant := range variants {
			result := KindFromString(variant)
			if result != expected {
				t.Errorf("KindFromString(%q) = %v, want %v (same as %q)",
					variant, result, expected, variants[0])
			}
		}
	}
}
