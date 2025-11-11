package neo4jkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared/kubernetes"
)

func TestNeo4JKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Neo4JKubernetes Suite")
}

var _ = ginkgo.Describe("Neo4JKubernetes Custom Validation Tests", func() {
	var input *Neo4JKubernetes

	ginkgo.BeforeEach(func() {
		input = &Neo4JKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "Neo4jKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-neo4j",
			},
			Spec: &Neo4JKubernetesSpec{
				Container: &Neo4JKubernetesContainer{
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
				Ingress: &Neo4JKubernetesIngress{
					Enabled:  true,
					Hostname: "neo4j.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("neo4j_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
