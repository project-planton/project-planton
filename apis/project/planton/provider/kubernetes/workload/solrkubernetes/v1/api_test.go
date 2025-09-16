package solrkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestSolrKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SolrKubernetes Suite")
}

var _ = ginkgo.Describe("SolrKubernetes Custom Validation Tests", func() {
	var input *SolrKubernetes

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("solr_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
