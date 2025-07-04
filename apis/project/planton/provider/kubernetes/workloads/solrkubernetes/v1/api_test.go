package solrkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestSolrKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolrKubernetes Suite")
}

var _ = Describe("SolrKubernetes Custom Validation Tests", func() {
	var input *SolrKubernetes

	BeforeEach(func() {
		input = &SolrKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "SolrKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-solr",
			},
			Spec: &SolrKubernetesSpec{
				SolrContainer: &SolrKubernetesSolrContainer{
					Replicas: 1,
					DiskSize: "10Gi",
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
					Image: &kubernetes.ContainerImage{
						Repo: "solr",
						Tag:  "8.7.0",
					},
				},
				ZookeeperContainer: &SolrKubernetesZookeeperContainer{
					Replicas: 1,
					DiskSize: "10Gi",
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
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("solr_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
