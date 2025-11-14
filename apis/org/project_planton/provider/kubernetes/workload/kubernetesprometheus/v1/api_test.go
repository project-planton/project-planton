package kubernetesprometheusv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesPrometheus(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesPrometheus Suite")
}

var _ = ginkgo.Describe("KubernetesPrometheus Custom Validation Tests", func() {
	var input *KubernetesPrometheus

	ginkgo.BeforeEach(func() {
		input = &KubernetesPrometheus{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesPrometheus",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-prom",
			},
			Spec: &KubernetesPrometheusSpec{
				Container: &KubernetesPrometheusContainer{
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
