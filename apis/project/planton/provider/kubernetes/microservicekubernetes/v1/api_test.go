package microservicekubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestMicroserviceKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MicroserviceKubernetes Suite")
}

var _ = Describe("MicroserviceKubernetes Custom Validation Tests", func() {
	var input *MicroserviceKubernetes

	BeforeEach(func() {
		input = &MicroserviceKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "MicroserviceKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "sample-msk8s",
			},
			Spec: &MicroserviceKubernetesSpec{
				Version: "review-123", // Valid according to the custom regex checks
				Container: &MicroserviceKubernetesContainer{
					App: &MicroserviceKubernetesContainerApp{
						Image: &kubernetes.ContainerImage{
							Repo: "example",
							Tag:  "latest",
						},
						Resources: &kubernetes.ContainerResources{
							Limits: &kubernetes.CpuMemory{
								Cpu:    "500m",
								Memory: "512Mi",
							},
							Requests: &kubernetes.CpuMemory{
								Cpu:    "250m",
								Memory: "256Mi",
							},
						},
						Env: &MicroserviceKubernetesContainerAppEnv{
							Variables: map[string]string{"KEY": "value"},
							Secrets:   map[string]string{"SECRET_KEY": "secret_value"},
						},
						Ports: []*MicroserviceKubernetesContainerAppPort{
							{
								Name:            "web1",
								ContainerPort:   8080,
								NetworkProtocol: "TCP",
								AppProtocol:     "http",
								ServicePort:     80,
								IsIngressPort:   true,
							},
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
						TargetCpuUtilizationPercent: 70,
						TargetMemoryUtilization:     "512Mi",
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("microservice_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
