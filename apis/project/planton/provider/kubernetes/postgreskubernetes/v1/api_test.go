package postgreskubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestPostgresKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PostgresKubernetesSpec Suite")
}

var _ = Describe("PostgresKubernetesSpec", func() {
	Context("with a fully valid spec", func() {
		It("should pass validation with no errors", func() {
			spec := &PostgresKubernetesSpec{
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
					DiskSize: "5Gi",
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "postgres.example.com",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when disk size is invalid", func() {
		It("should fail validation with `[spec.container.disk_size.required]` in the error message", func() {
			spec := &PostgresKubernetesSpec{
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
					DiskSize: "invalid-size", // Invalid format
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for invalid disk size")
			Expect(err.Error()).To(ContainSubstring("[spec.container.disk_size.required]"),
				"expected disk size format error substring in the validation message")
		})
	})

	Context("when disk size is not provided", func() {
		It("should fail validation if it’s required despite any default mechanism", func() {
			spec := &PostgresKubernetesSpec{
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
					// No DiskSize
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for missing disk size")
			Expect(err.Error()).To(ContainSubstring("[spec.container.disk_size.required]"),
				"expected disk size format error in the message")
		})
	})

	Context("when ingress is omitted", func() {
		It("should pass validation if ingress is optional", func() {
			spec := &PostgresKubernetesSpec{
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
					DiskSize: "1Gi",
				},
				// No ingress
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors when ingress is omitted")
		})
	})

	Context("when replicas are large", func() {
		It("should pass validation if there is no constraint on maximum replicas", func() {
			spec := &PostgresKubernetesSpec{
				Container: &PostgresKubernetesContainer{
					Replicas: 10,
					Resources: &kubernetes.ContainerResources{
						Limits: &kubernetes.CpuMemory{
							Cpu:    "2000m",
							Memory: "2Gi",
						},
						Requests: &kubernetes.CpuMemory{
							Cpu:    "100m",
							Memory: "200Mi",
						},
					},
					DiskSize: "10Gi",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors for large replica count")
		})
	})
})
