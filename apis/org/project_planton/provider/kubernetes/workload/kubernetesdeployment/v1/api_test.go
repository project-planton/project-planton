package kubernetesdeploymentv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesDeployment(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesDeployment Suite")
}

var _ = ginkgo.Describe("KubernetesDeployment Custom Validation Tests", func() {
	var input *KubernetesDeployment

	ginkgo.BeforeEach(func() {
		input = &KubernetesDeployment{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesDeployment",
			Metadata: &shared.CloudResourceMetadata{
				Name: "sample-k8sms",
			},
			Spec: &KubernetesDeploymentSpec{
				Version: "review-123", // Valid according to the custom regex checks
				Container: &KubernetesDeploymentContainer{
					App: &KubernetesDeploymentContainerApp{
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
						Env: &KubernetesDeploymentContainerAppEnv{
							Variables: map[string]string{"KEY": "value"},
							Secrets:   map[string]string{"SECRET_KEY": "secret_value"},
						},
						Ports: []*KubernetesDeploymentContainerAppPort{
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
				Availability: &KubernetesDeploymentAvailability{
					MinReplicas: 1,
					HorizontalPodAutoscaling: &KubernetesDeploymentAvailabilityHpa{
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
