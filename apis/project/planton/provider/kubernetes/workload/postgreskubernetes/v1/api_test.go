package postgreskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestPostgresKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "PostgresKubernetes Suite")
}

var _ = ginkgo.Describe("PostgresKubernetes Custom Validation Tests", func() {
	var input *PostgresKubernetes

	ginkgo.BeforeEach(func() {
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
				Ingress: &PostgresKubernetesIngress{
					Enabled:  true,
					Hostname: "postgres.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("postgres_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
