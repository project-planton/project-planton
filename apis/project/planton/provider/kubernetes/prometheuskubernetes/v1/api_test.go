package prometheuskubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestPrometheusKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusKubernetes Suite")
}

var _ = Describe("PrometheusKubernetes Custom Validation Tests", func() {
	var input *PrometheusKubernetes

	BeforeEach(func() {
		input = &PrometheusKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "PrometheusKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-prom",
			},
			Spec: &PrometheusKubernetesSpec{
				Container: &PrometheusKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					DiskSize:             "10Gi",
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
					DnsDomain: "prometheus.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("prometheus_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
