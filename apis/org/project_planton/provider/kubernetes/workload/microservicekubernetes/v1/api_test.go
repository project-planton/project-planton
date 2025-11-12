package microservicekubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestMicroserviceKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "MicroserviceKubernetes Suite")
}

var _ = ginkgo.Describe("MicroserviceKubernetes Custom Validation Tests", func() {
	var input *MicroserviceKubernetes

	ginkgo.BeforeEach(func() {
		input = &MicroserviceKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "MicroserviceKubernetes",
			Metadata: &shared.CloudResourceMetadata{
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("microservice_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
