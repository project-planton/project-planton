package kubernetes

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	microservicekubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/microservicekubernetes/v1"
	"github.com/project-planton/project-planton/apis/go/project/planton/shared/kubernetes"
)

// TestMicroserviceKubernetesSpec_ValidSpec ensures a fully valid spec passes validation.
func TestMicroserviceKubernetesSpec_ValidSpec(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-123",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
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
				Env: &microservicekubernetesv1.MicroserviceKubernetesContainerAppEnv{
					Variables: map[string]string{
						"ENV_VAR": "value",
					},
					Secrets: map[string]string{
						"SECRET_VAR": "secretValue",
					},
				},
				Ports: []*microservicekubernetesv1.MicroserviceKubernetesContainerAppPort{
					{
						Name:            "http",
						ContainerPort:   8080,
						NetworkProtocol: "TCP",
						AppProtocol:     "http",
						ServicePort:     80,
						IsIngressPort:   true,
					},
					{
						Name:            "admin",
						ContainerPort:   9090,
						NetworkProtocol: "TCP",
						AppProtocol:     "http",
						ServicePort:     9090,
					},
				},
			},
			Sidecars: []*kubernetes.Container{
				{
					Image: "sidecar-repo",
				},
			},
		},
		Ingress: &kubernetes.IngressSpec{
			DnsDomain: "myapp.example.com",
		},
		Availability: &microservicekubernetesv1.MicroserviceKubernetesAvailability{
			MinReplicas: 1,
			HorizontalPodAutoscaling: &microservicekubernetesv1.MicroserviceKubernetesAvailabilityHpa{
				IsEnabled:                   true,
				TargetCpuUtilizationPercent: 60.0,
				TargetMemoryUtilization:     "1Gi",
			},
		},
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors, got: %v", err)
	}
}

// TestMicroserviceKubernetesSpec_InvalidVersion checks that invalid version formats fail validation.
func TestMicroserviceKubernetesSpec_InvalidVersion(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "Invalid_Character", // Underscore not allowed
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid version, got none")
	} else {
		if !strings.Contains(err.Error(), "Only lowercase letters, numbers, and hyphens are allowed") {
			t.Errorf("expected error about allowed version chars, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_VersionEndsWithHyphen checks that version ending with a hyphen fails validation.
func TestMicroserviceKubernetesSpec_VersionEndsWithHyphen(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-123-",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for version ending with a hyphen, got none")
	} else {
		if !strings.Contains(err.Error(), "Must not end with a hyphen") {
			t.Errorf("expected error about ending hyphen, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_VersionTooLong checks that version longer than 20 chars fails validation.
func TestMicroserviceKubernetesSpec_VersionTooLong(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "thisiswaytoolongforarule",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for overly long version, got none")
	} else {
		if !strings.Contains(err.Error(), "Field length must be at most 20") {
			t.Errorf("expected an error about max length, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_MissingContainer checks that missing container field fails validation.
func TestMicroserviceKubernetesSpec_MissingContainer(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-1",
		// Container missing
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing container, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for container, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_MissingApp checks that missing app field in container fails validation.
func TestMicroserviceKubernetesSpec_MissingApp(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version:   "review-1",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			// App missing
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing container.app, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected a 'Field is required' error for container.app, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_InvalidPortName checks that invalid port names fail validation.
func TestMicroserviceKubernetesSpec_InvalidPortName(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-2",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
				Ports: []*microservicekubernetesv1.MicroserviceKubernetesContainerAppPort{
					{
						Name:            "-invalid", // starts with a hyphen
						ContainerPort:   8080,
						NetworkProtocol: "TCP",
						AppProtocol:     "http",
						ServicePort:     80,
					},
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid port name, got none")
	} else {
		if !strings.Contains(err.Error(), "Name for ports must only contain lowercase alphanumeric characters and hyphens") {
			t.Errorf("expected error about port name format, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_InvalidNetworkProtocol checks that invalid network protocols fail validation.
func TestMicroserviceKubernetesSpec_InvalidNetworkProtocol(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-2",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
				Ports: []*microservicekubernetesv1.MicroserviceKubernetesContainerAppPort{
					{
						Name:            "web",
						ContainerPort:   8080,
						NetworkProtocol: "INVALID",
						AppProtocol:     "http",
						ServicePort:     80,
					},
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for invalid network protocol, got none")
	} else {
		if !strings.Contains(err.Error(), "The network protocol must be one of \"SCTP\", \"TCP\", or \"UDP\"") {
			t.Errorf("expected error about network protocol, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_MissingPortFields checks that missing required fields in ports fail validation.
func TestMicroserviceKubernetesSpec_MissingPortFields(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-2",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
				Ports: []*microservicekubernetesv1.MicroserviceKubernetesContainerAppPort{
					{
						// Missing name, container_port, network_protocol, app_protocol, service_port
					},
				},
			},
		},
	}

	err := protovalidate.Validate(spec)
	if err == nil {
		t.Errorf("expected validation error for missing port fields, got none")
	} else {
		if !strings.Contains(err.Error(), "Field is required") {
			t.Errorf("expected 'Field is required' errors for port fields, got: %v", err)
		}
	}
}

// TestMicroserviceKubernetesSpec_NoAvailability checks if availability is optional and valid if not provided.
func TestMicroserviceKubernetesSpec_NoAvailability(t *testing.T) {
	spec := &microservicekubernetesv1.MicroserviceKubernetesSpec{
		Version: "review-2",
		Container: &microservicekubernetesv1.MicroserviceKubernetesContainer{
			App: &microservicekubernetesv1.MicroserviceKubernetesContainerApp{
				Image: &kubernetes.ContainerImage{
					Repo: "my-repo",
					Tag:  "latest",
				},
				Ports: []*microservicekubernetesv1.MicroserviceKubernetesContainerAppPort{
					{
						Name:            "web",
						ContainerPort:   8080,
						NetworkProtocol: "TCP",
						AppProtocol:     "http",
						ServicePort:     80,
					},
				},
			},
		},
		// No availability provided
	}

	if err := protovalidate.Validate(spec); err != nil {
		t.Errorf("expected no validation errors without availability if it's optional, got: %v", err)
	}
}
