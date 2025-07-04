package rediskubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestRedisKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisKubernetes Suite")
}

var _ = Describe("RedisKubernetes Custom Validation Tests", func() {
	var input *RedisKubernetes

	BeforeEach(func() {
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
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "redis.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("redis_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
