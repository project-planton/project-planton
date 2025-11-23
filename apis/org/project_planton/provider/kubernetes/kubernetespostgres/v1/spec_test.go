package kubernetespostgresv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesPostgres(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesPostgres Suite")
}

var _ = ginkgo.Describe("KubernetesPostgres Custom Validation Tests", func() {
	var input *KubernetesPostgres

	ginkgo.BeforeEach(func() {
		input = &KubernetesPostgres{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesPostgres",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-pg",
			},
			Spec: &KubernetesPostgresSpec{
				Container: &KubernetesPostgresContainer{
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
				Ingress: &KubernetesPostgresIngress{
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
