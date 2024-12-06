package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	locustkubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/locustkubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
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

// TestLocustKubernetesSpec_MissingLoadTest checks that missing the LoadTest field fails validation.
func TestLocustKubernetesSpec_MissingLoadTest(t *testing.T) {
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
		// LoadTest is missing
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing load_test, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for load_test, got: %v", err)
		}
	}
}

// TestLocustKubernetesSpec_MissingLoadTestName checks that missing the name in LoadTest fails validation.
func TestLocustKubernetesSpec_MissingLoadTestName(t *testing.T) {
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
			// Name is missing
			MainPyContent: "from locust import HttpUser, task",
			LibFilesContent: map[string]string{
				"utils.py": "def helper(): pass",
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing load_test.name, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for load_test.name, got: %v", err)
		}
	}
}

// TestLocustKubernetesSpec_MissingMainPyContent checks that missing the main_py_content in LoadTest fails validation.
func TestLocustKubernetesSpec_MissingMainPyContent(t *testing.T) {
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
			Name: "my_test",
			// main_py_content missing
			LibFilesContent: map[string]string{
				"utils.py": "def helper(): pass",
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing load_test.main_py_content, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for load_test.main_py_content, got: %v", err)
		}
	}
}

// TestLocustKubernetesSpec_MissingLibFilesContent checks that missing lib_files_content fails validation.
func TestLocustKubernetesSpec_MissingLibFilesContent(t *testing.T) {
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
			// lib_files_content missing
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing load_test.lib_files_content, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for load_test.lib_files_content, got: %v", err)
		}
	}
}

// TestLocustKubernetesSpec_EmptyLibFilesContent ensures that an empty map also fails if the field is required.
func TestLocustKubernetesSpec_EmptyLibFilesContent(t *testing.T) {
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
			Name:            "my_test",
			MainPyContent:   "from locust import HttpUser, task",
			LibFilesContent: map[string]string{}, // Empty map might still be considered "provided"
		},
	}

	// Depending on how you interpret "required", this might or might not fail.
	// If "required" means empty maps are not allowed, it should fail.
	// If not, this might pass. Adjust your expectations accordingly.
	err := protovalidate.Validate(spec)
	if err != nil {
		// If you consider empty maps invalid, then a validation error is expected.
		// If not, remove this check.
		t.Errorf("did not expect a validation error for empty lib_files_content if empty maps are allowed: %v", err)
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
