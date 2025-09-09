package elasticsearchkubernetesv1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestElasticsearchKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ElasticsearchKubernetes Suite")
}

var _ = Describe("ElasticsearchKubernetes Custom Validation Tests", func() {
	var input *ElasticsearchKubernetes

	BeforeEach(func() {
		input = &ElasticsearchKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ElasticsearchKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-es",
			},
			Spec: &ElasticsearchKubernetesSpec{
				ElasticsearchContainer: &ElasticsearchKubernetesElasticsearchContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					DiskSize:             "10Gi", // valid format
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
				KibanaContainer: &ElasticsearchKubernetesKibanaContainer{
					IsEnabled: true,
					Replicas:  1,
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
					DnsDomain: "elasticsearch.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("elasticsearch_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
