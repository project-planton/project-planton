package kubernetesredisv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesRedis(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesRedis Suite")
}

var _ = ginkgo.Describe("KubernetesRedis Custom Validation Tests", func() {
	var input *KubernetesRedis

	ginkgo.BeforeEach(func() {
		input = &KubernetesRedis{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesRedis",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-redis",
			},
			Spec: &KubernetesRedisSpec{
				Container: &KubernetesRedisContainer{
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
				Ingress: &KubernetesRedisIngress{
					Enabled:  true,
					Hostname: "redis.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("redis_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
