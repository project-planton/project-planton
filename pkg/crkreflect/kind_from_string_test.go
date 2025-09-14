package crkreflect

import (
	"testing"

	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
)

func TestKindFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected cloudresourcekind.CloudResourceKind
	}{
		// MicroserviceKubernetes tests
		{
			name:     "MicroserviceKubernetes - PascalCase",
			input:    "MicroserviceKubernetes",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
		},
		{
			name:     "MicroserviceKubernetes - kebab-case",
			input:    "microservice-kubernetes",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
		},
		{
			name:     "MicroserviceKubernetes - snake_case",
			input:    "microservice_kubernetes",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
		},
		{
			name:     "MicroserviceKubernetes - UPPERCASE",
			input:    "MICROSERVICEKUBERNETES",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
		},
		{
			name:     "MicroserviceKubernetes - mixed case with hyphens",
			input:    "Microservice-Kubernetes",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
		},
		{
			name:     "MicroserviceKubernetes - mixed case with underscores",
			input:    "Microservice_Kubernetes",
			expected: cloudresourcekind.CloudResourceKind_MicroserviceKubernetes,
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
			"MicroserviceKubernetes",
			"microservicekubernetes",
			"MICROSERVICEKUBERNETES",
			"microservice-kubernetes",
			"microservice_kubernetes",
			"Microservice-Kubernetes",
			"Microservice_Kubernetes",
			"MICROSERVICE-KUBERNETES",
			"MICROSERVICE_KUBERNETES",
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
