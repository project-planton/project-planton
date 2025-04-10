package rediskubernetesv1

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestRedisKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisKubernetesSpec Suite")
}

var _ = Describe("RedisKubernetesSpec", func() {
	Context("with a valid spec", func() {
		It("should not return validation errors", func() {
			spec := &RedisKubernetesSpec{
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
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when persistence is enabled but no disk_size is provided", func() {
		It("should fail validation, expecting `spec.container.disk_size.required`", func() {
			spec := &RedisKubernetesSpec{
				Container: &RedisKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					// No disk_size
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
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected a validation error for missing disk_size when persistence is enabled")
			Expect(strings.Contains(err.Error(), "[spec.container.disk_size.required]")).
				To(BeTrue(), "expected error with constraint id `spec.container.disk_size.required`")
		})
	})

	Context("when disk_size has an invalid format", func() {
		It("should fail validation, expecting `spec.container.disk_size.required`", func() {
			spec := &RedisKubernetesSpec{
				Container: &RedisKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					DiskSize:             "abc", // invalid format
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
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected a validation error for invalid disk_size format")
			Expect(strings.Contains(err.Error(), "[spec.container.disk_size.required]")).
				To(BeTrue(), "expected error with constraint id `spec.container.disk_size.required`")
		})
	})

	Context("when persistence is disabled and disk_size is empty", func() {
		It("should not return a validation error", func() {
			spec := &RedisKubernetesSpec{
				Container: &RedisKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: false,
					// disk_size is empty, but persistence is disabled
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
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "did not expect a validation error when persistence is disabled and disk_size is empty")
		})
	})
})
