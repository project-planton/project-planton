package kubernetes

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	locustkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/locustkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

// TestLocustKubernetesSpec_ValidSpec ensures a fully valid spec passes validation.
func TestLocustKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &locustkubernetesv1.LocustKubernetesSpec{
		MasterContainer: &locustkubernetesv1.LocustKubernetesContainer{
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
		WorkerContainer: &locustkubernetesv1.LocustKubernetesContainer{
			Replicas: 5,
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
			DnsDomain: "locust.example.com",
		},
		LoadTest: &locustkubernetesv1.LocustKubernetesLoadTest{
			Name:          "my_load_test",
			MainPyContent: "from locust import HttpUser, task",
			LibFilesContent: map[string]string{
				"utils.py": "def helper(): pass",
			},
			PipPackages: []string{"requests", "locustio"},
		},
		HelmValues: map[string]string{
			"image.tag": "latest",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestLocustKubernetesSpec_NoHelmValues is valid if helm_values is optional.
func TestLocustKubernetesSpec_NoHelmValues(t *testing.T) {
	spec := &locustkubernetesv1.LocustKubernetesSpec{
		MasterContainer: &locustkubernetesv1.LocustKubernetesContainer{
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
		WorkerContainer: &locustkubernetesv1.LocustKubernetesContainer{
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
		LoadTest: &locustkubernetesv1.LocustKubernetesLoadTest{
			Name:          "my_test",
			MainPyContent: "from locust import HttpUser, task",
			LibFilesContent: map[string]string{
				"utils.py": "def helper(): pass",
			},
		},
		// No helm_values provided, which should be fine if optional
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}
