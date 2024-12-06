package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/postgreskubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
)

// TestPostgresKubernetesSpec_ValidSpec ensures that a fully valid spec passes validation.
func TestPostgresKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &postgreskubernetesv1.PostgresKubernetesSpec{
		Container: &postgreskubernetesv1.PostgresKubernetesContainer{
			Replicas: 1,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "1000m",
					Memory: "1Gi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "50m",
					Memory: "100Mi",
				},
			},
			DiskSize: "5Gi",
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "postgres.example.com",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestPostgresKubernetesSpec_InvalidDiskSize checks that an invalid disk size format fails validation.
func TestPostgresKubernetesSpec_InvalidDiskSize(t *testing.T) {
	spec := &postgreskubernetesv1.PostgresKubernetesSpec{
		Container: &postgreskubernetesv1.PostgresKubernetesContainer{
			Replicas: 1,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "1000m",
					Memory: "1Gi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "50m",
					Memory: "100Mi",
				},
			},
			DiskSize: "invalid-size", // Invalid format
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid disk size, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.format]") {
			t.Errorf("expected disk size format error, got: %v", err)
		}
	}
}

// TestPostgresKubernetesSpec_EmptyDiskSize checks that using the default disk size is still valid if optional.
// If the field is optional and has a default, this should pass.
func TestPostgresKubernetesSpec_EmptyDiskSize(t *testing.T) {
	spec := &postgreskubernetesv1.PostgresKubernetesSpec{
		Container: &postgreskubernetesv1.PostgresKubernetesContainer{
			Replicas: 1,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "1000m",
					Memory: "1Gi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "50m",
					Memory: "100Mi",
				},
			},
			// DiskSize not provided, should fall back to default if allowed
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors when disk_size is omitted if default is allowed, got: %v", err)
	}
}

// TestPostgresKubernetesSpec_NoIngress checks if ingress is optional and can be omitted without validation errors.
func TestPostgresKubernetesSpec_NoIngress(t *testing.T) {
	spec := &postgreskubernetesv1.PostgresKubernetesSpec{
		Container: &postgreskubernetesv1.PostgresKubernetesContainer{
			Replicas: 1,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "1000m",
					Memory: "1Gi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "50m",
					Memory: "100Mi",
				},
			},
			DiskSize: "1Gi",
		},
		// No ingress provided
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors when ingress is omitted if it's optional, got: %v", err)
	}
}

// TestPostgresKubernetesSpec_LargeReplicas ensures that large replicas count is allowed if there's no explicit constraint.
func TestPostgresKubernetesSpec_LargeReplicas(t *testing.T) {
	spec := &postgreskubernetesv1.PostgresKubernetesSpec{
		Container: &postgreskubernetesv1.PostgresKubernetesContainer{
			Replicas: 10,
			Resources: &kubernetes.ContainerResources{
				Limits: &kubernetes.CpuMemory{
					Cpu:    "2000m",
					Memory: "2Gi",
				},
				Requests: &kubernetes.CpuMemory{
					Cpu:    "100m",
					Memory: "200Mi",
				},
			},
			DiskSize: "10Gi",
		},
	}

	// Assuming no explicit replica constraints, this should pass.
	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors for large replica count, got: %v", err)
	}
}
