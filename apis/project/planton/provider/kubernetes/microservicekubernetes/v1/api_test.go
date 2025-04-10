package microservicekubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestMicroserviceKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MicroserviceKubernetesSpec Suite")
}

var _ = Describe("MicroserviceKubernetesSpec", func() {
	Context("when the spec is fully valid", func() {
		It("should pass validation without errors", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "review-123",
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
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
						Env: &MicroserviceKubernetesContainerAppEnv{
							Variables: map[string]string{
								"ENV_VAR": "value",
							},
							Secrets: map[string]string{
								"SECRET_VAR": "secretValue",
							},
						},
						Ports: []*MicroserviceKubernetesContainerAppPort{
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
				Availability: &MicroserviceKubernetesAvailability{
					MinReplicas: 1,
					HorizontalPodAutoscaling: &MicroserviceKubernetesAvailabilityHpa{
						IsEnabled:                   true,
						TargetCpuUtilizationPercent: 60.0,
						TargetMemoryUtilization:     "1Gi",
					},
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors")
		})
	})

	Context("when the version is invalid", func() {
		It("should fail validation if it has disallowed characters", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "Invalid_Character", // underscore not allowed
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "my-repo",
							Tag:  "latest",
						},
					},
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "Expected validation error for invalid version")
			Expect(err.Error()).To(ContainSubstring("Only lowercase letters, numbers, and hyphens are allowed"),
				"Expected error about allowed version characters")
		})

		It("should fail validation if it ends with a hyphen", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "review-123-",
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "my-repo",
							Tag:  "latest",
						},
					},
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "Expected validation error for version ending with a hyphen")
			Expect(err.Error()).To(ContainSubstring("Must not end with a hyphen"),
				"Expected error about ending hyphen")
		})
	})

	Context("when the port name is invalid", func() {
		It("should fail validation if it starts with a hyphen", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "review-2",
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "my-repo",
							Tag:  "latest",
						},
						Ports: []*MicroserviceKubernetesContainerAppPort{
							{
								Name:            "-invalid",
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
			Expect(err).NotTo(BeNil(), "Expected validation error for invalid port name")
			Expect(err.Error()).To(ContainSubstring("Name for ports must only contain lowercase alphanumeric characters and hyphens"),
				"Expected error about port name format")
		})
	})

	Context("when the network protocol is invalid", func() {
		It("should fail validation if not one of SCTP, TCP, or UDP", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "review-2",
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "my-repo",
							Tag:  "latest",
						},
						Ports: []*MicroserviceKubernetesContainerAppPort{
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
			Expect(err).NotTo(BeNil(), "Expected validation error for invalid network protocol")
			Expect(err.Error()).To(ContainSubstring("The network protocol must be one of \"SCTP\", \"TCP\", or \"UDP\""),
				"Expected error about network protocol")
		})
	})

	Context("when availability is not provided", func() {
		It("should pass validation if availability is optional", func() {
			spec := &MicroserviceKubernetesSpec{
				Version: "review-2",
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "my-repo",
							Tag:  "latest",
						},
						Ports: []*MicroserviceKubernetesContainerAppPort{
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
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors without availability if it's optional")
		})
	})
})
