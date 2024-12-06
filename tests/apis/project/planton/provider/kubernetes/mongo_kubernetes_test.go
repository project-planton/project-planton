package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	mongodbkubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
)

// TestMongodbKubernetesSpec_ValidSpec ensures that a fully valid spec passes validation.
func TestMongodbKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &mongodbkubernetesv1.MongodbKubernetesSpec{
		Container: &mongodbkubernetesv1.MongodbKubernetesContainer{
			Replicas: 3,
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
			IsPersistenceEnabled: true,
			DiskSize:             "10Gi",
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "mongo.example.com",
		},
		HelmValues: map[string]string{
			"auth.rootPassword": "secret",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestMongodbKubernetesSpec_InvalidDiskSizeFormat checks that an invalid disk size fails validation.
func TestMongodbKubernetesSpec_InvalidDiskSizeFormat(t *testing.T) {
	spec := &mongodbkubernetesv1.MongodbKubernetesSpec{
		Container: &mongodbkubernetesv1.MongodbKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
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
			DiskSize: "abc", // Invalid format
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid disk size format, got none")
	} else {
		if !strings.Contains(err.Error(), "[spec.container.disk_size.required]") {
			t.Errorf("expected disk size format error, got: %v", err)
		}
	}
}

// TestMongodbKubernetesSpec_NoDiskSizeWhenPersistenceEnabled checks that absence of disk_size fails validation if persistence is enabled.
func TestMongodbKubernetesSpec_NoDiskSizeWhenPersistenceEnabled(t *testing.T) {
	spec := &mongodbkubernetesv1.MongodbKubernetesSpec{
		Container: &mongodbkubernetesv1.MongodbKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
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
			// Missing disk_size
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing disk size when persistence is enabled, got none")
	} else {
		if !strings.Contains(err.Error(), "Disk size is required") {
			t.Errorf("expected error about required disk size, got: %v", err)
		}
	}
}

// TestMongodbKubernetesSpec_PersistenceDisabledNoDiskSize ensures that if persistence is disabled, no disk_size is required.
func TestMongodbKubernetesSpec_PersistenceDisabledNoDiskSize(t *testing.T) {
	spec := &mongodbkubernetesv1.MongodbKubernetesSpec{
		Container: &mongodbkubernetesv1.MongodbKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: false, // No persistence, so disk_size isn't required
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
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors without persistence and no disk_size, got: %v", err)
	}
}

// TestMongodbKubernetesSpec_EmptyHelmValues ensures it's valid to have empty helm_values if optional.
func TestMongodbKubernetesSpec_EmptyHelmValues(t *testing.T) {
	spec := &mongodbkubernetesv1.MongodbKubernetesSpec{
		Container: &mongodbkubernetesv1.MongodbKubernetesContainer{
			Replicas:             1,
			IsPersistenceEnabled: true,
			DiskSize:             "1Gi",
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
		},
		// helm_values not provided
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors for empty helm_values, got: %v", err)
	}
}
