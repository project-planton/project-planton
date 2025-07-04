package neo4jkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestNeo4JKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Neo4JKubernetes Suite")
}

var _ = Describe("Neo4JKubernetes Custom Validation Tests", func() {
	var input *Neo4JKubernetes

	BeforeEach(func() {
		input = &Neo4JKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "Neo4jKubernetes",
			Metadata: &shared.ApiResourceMetadata{
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
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "neo4j.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("neo4j_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
