package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	openfgakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/openfgakubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

// TestOpenfgaKubernetesSpec_ValidSpec ensures that a fully valid spec passes validation.
func TestOpenfgaKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &openfgakubernetesv1.OpenfgaKubernetesSpec{
		Container: &openfgakubernetesv1.OpenfgaKubernetesContainer{
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
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "openfga.example.com",
		},
		Datastore: &openfgakubernetesv1.OpenfgaKubernetesDataStore{
			Engine: "postgres",
			Uri:    "postgres://user:pass@localhost:5432/mydb",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestOpenfgaKubernetesDataStore_InvalidEngine checks that invalid engine values fail validation.
func TestOpenfgaKubernetesDataStore_InvalidEngine(t *testing.T) {
	spec := &openfgakubernetesv1.OpenfgaKubernetesSpec{
		Container: &openfgakubernetesv1.OpenfgaKubernetesContainer{
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
		},
		Datastore: &openfgakubernetesv1.OpenfgaKubernetesDataStore{
			Engine: "sqlite", // Invalid
			Uri:    "sqlite://path/to/db",
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid engine, got none")
	} else {
		if !strings.Contains(err.Error(), "The datastore engine must be one of \"postgres\" and \"mysql\".") {
			t.Errorf("expected error about allowed engine values, got: %v", err)
		}
	}
}

// TestOpenfgaKubernetesSpec_EmptyIngress ensures that if ingress is optional and empty, it's still valid.
func TestOpenfgaKubernetesSpec_EmptyIngress(t *testing.T) {
	spec := &openfgakubernetesv1.OpenfgaKubernetesSpec{
		Container: &openfgakubernetesv1.OpenfgaKubernetesContainer{
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
		},
		Datastore: &openfgakubernetesv1.OpenfgaKubernetesDataStore{
			Engine: "postgres",
			Uri:    "postgres://user:pass@localhost:5432/mydb",
		},
		// No ingress provided
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors for omitted optional ingress, got: %v", err)
	}
}
