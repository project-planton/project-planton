package postgreskubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestPostgresKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PostgresKubernetes Suite")
}

var _ = Describe("PostgresKubernetes Custom Validation Tests", func() {
	var input *PostgresKubernetes

	BeforeEach(func() {
		input = &PostgresKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "PostgresKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-pg",
			},
			Spec: &PostgresKubernetesSpec{
				Container: &PostgresKubernetesContainer{
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
					DiskSize: "2Gi", // valid disk size
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "postgres.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("postgres_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
