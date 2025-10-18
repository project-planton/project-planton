package rediskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestRedisKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RedisKubernetes Suite")
}

var _ = ginkgo.Describe("RedisKubernetes Custom Validation Tests", func() {
	var input *RedisKubernetes

	ginkgo.BeforeEach(func() {
		input = &RedisKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "RedisKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-redis",
			},
			Spec: &RedisKubernetesSpec{
				Container: &RedisKubernetesContainer{
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
				Ingress: &RedisKubernetesIngress{
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
