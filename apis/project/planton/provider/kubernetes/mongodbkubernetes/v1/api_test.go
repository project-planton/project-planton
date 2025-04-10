package mongodbkubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestMongodbKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MongodbKubernetesSpec Suite")
}

var _ = Describe("MongodbKubernetesSpec", func() {
	Context("with a fully valid spec", func() {
		It("should pass validation with no errors", func() {
			spec := &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas: 3,
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
					IsPersistenceEnabled: true,
					DiskSize:             "10Gi",
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "mongo.example.com",
				},
				HelmValues: map[string]string{
					"auth.rootPassword": "secret",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when disk size format is invalid", func() {
		It("should fail validation with `[spec.container.disk_size.required]` in the error", func() {
			spec := &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
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
					DiskSize: "abc", // Invalid format
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for invalid disk size format")
			Expect(err.Error()).To(ContainSubstring("[spec.container.disk_size.required]"),
				"expected disk size format error substring in the validation message")
		})
	})

	Context("when persistence is enabled but disk_size is missing", func() {
		It("should fail validation with an error mentioning 'Disk size is required'", func() {
			spec := &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
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
					// No diskSize
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for missing disk size")
			Expect(err.Error()).To(ContainSubstring("Disk size is required"),
				"expected error message about required disk size")
		})
	})

	Context("when persistence is disabled and disk_size is missing", func() {
		It("should pass validation without errors", func() {
			spec := &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: false,
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
					// diskSize omitted
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors without persistence and no diskSize")
		})
	})

	Context("when helm_values is empty", func() {
		It("should pass validation if optional", func() {
			spec := &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					DiskSize:             "1Gi",
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
				// HelmValues not provided
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors for empty helm_values")
		})
	})
})
