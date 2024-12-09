package kubernetes

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	jenkinskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/jenkinskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestJenkinsKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &jenkinskubernetesv1.JenkinsKubernetesSpec{
		ContainerResources: &kubernetes.ContainerResources{
			Limits: &kubernetes.CpuMemory{
				Cpu:    "2000m",
				Memory: "2Gi",
			},
			Requests: &kubernetes.CpuMemory{
				Cpu:    "100m",
				Memory: "200Mi",
			},
		},
		HelmValues: map[string]string{
			"controller.tag": "lts",
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "jenkins.example.com",
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

func TestJenkinsKubernetesSpec_DefaultResources(t *testing.T) {
	// Not setting container_resources, expecting defaults to be applied without validation errors.
	spec := &jenkinskubernetesv1.JenkinsKubernetesSpec{}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors with default container_resources, got: %v", err)
	}
}

func TestJenkinsKubernetesSpec_EmptyHelmValues(t *testing.T) {
	// Empty helm_values map should be allowed.
	spec := &jenkinskubernetesv1.JenkinsKubernetesSpec{
		ContainerResources: &kubernetes.ContainerResources{
			Limits: &kubernetes.CpuMemory{
				Cpu:    "1000m",
				Memory: "1Gi",
			},
			Requests: &kubernetes.CpuMemory{
				Cpu:    "50m",
				Memory: "100Mi",
			},
		},
		HelmValues: make(map[string]string),
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors for empty helm_values, got: %v", err)
	}
}

func TestJenkinsKubernetesSpec_NoIngress(t *testing.T) {
	// No ingress set, should still pass as it's optional.
	spec := &jenkinskubernetesv1.JenkinsKubernetesSpec{
		ContainerResources: &kubernetes.ContainerResources{
			Limits: &kubernetes.CpuMemory{
				Cpu:    "1000m",
				Memory: "1Gi",
			},
			Requests: &kubernetes.CpuMemory{
				Cpu:    "50m",
				Memory: "100Mi",
			},
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors with no ingress, got: %v", err)
	}
}
