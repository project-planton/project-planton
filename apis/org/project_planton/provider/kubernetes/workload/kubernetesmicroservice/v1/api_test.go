package kubernetesmicroservicev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesMicroservice(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesMicroservice Suite")
}

var _ = ginkgo.Describe("KubernetesMicroservice Custom Validation Tests", func() {
	var input *KubernetesMicroservice

	ginkgo.BeforeEach(func() {
		input = &KubernetesMicroservice{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesMicroservice",
			Metadata: &shared.CloudResourceMetadata{
				Name: "sample-k8sms",
			},
			Spec: &KubernetesMicroserviceSpec{
				Version: "review-123", // Valid according to the custom regex checks
				Container: &KubernetesMicroserviceContainer{
					App: &KubernetesMicroserviceContainerApp{
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
						Env: &KubernetesMicroserviceContainerAppEnv{
							Variables: map[string]string{"KEY": "value"},
							Secrets:   map[string]string{"SECRET_KEY": "secret_value"},
						},
						Ports: []*KubernetesMicroserviceContainerAppPort{
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
				Availability: &KubernetesMicroserviceAvailability{
					MinReplicas: 1,
					HorizontalPodAutoscaling: &KubernetesMicroserviceAvailabilityHpa{
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
