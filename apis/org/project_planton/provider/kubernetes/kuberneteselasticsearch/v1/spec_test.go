package kuberneteselasticsearchv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"

	"buf.build/go/protovalidate"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesElasticsearch(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesElasticsearch Suite")
}

var _ = ginkgo.Describe("KubernetesElasticsearch Custom Validation Tests", func() {
	var input *KubernetesElasticsearch

	ginkgo.BeforeEach(func() {
		input = &KubernetesElasticsearch{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesElasticsearch",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-es",
			},
			Spec: &KubernetesElasticsearchSpec{
				Elasticsearch: &KubernetesElasticsearchElasticsearchSpec{
					Container: &KubernetesElasticsearchElasticsearchContainer{
						Replicas:           1,
						PersistenceEnabled: true,
						DiskSize:           "10Gi", // valid format
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
					Ingress: &KubernetesElasticsearchIngress{
						Enabled:  true,
						Hostname: "elasticsearch.example.com",
					},
				},
				Kibana: &KubernetesElasticsearchKibanaSpec{
					Enabled: true,
					Container: &KubernetesElasticsearchKibanaContainer{
						Replicas: 1,
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
					Ingress: &KubernetesElasticsearchIngress{
						Enabled:  true,
						Hostname: "kibana.example.com",
					},
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("elasticsearch_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
