package prometheuskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestPrometheusKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "PrometheusKubernetes Suite")
}

var _ = ginkgo.Describe("PrometheusKubernetes Custom Validation Tests", func() {
	var input *PrometheusKubernetes

	ginkgo.BeforeEach(func() {
		input = &PrometheusKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "PrometheusKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-prom",
			},
			Spec: &PrometheusKubernetesSpec{
				Container: &PrometheusKubernetesContainer{
					Replicas:           1,
					PersistenceEnabled: true,
					DiskSize:           "10Gi",
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("prometheus_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
