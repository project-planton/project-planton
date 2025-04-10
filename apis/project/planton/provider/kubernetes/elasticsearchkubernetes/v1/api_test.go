package elasticsearchkubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestElasticsearchKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ElasticsearchKubernetesSpec Suite")
}

var _ = Describe("ElasticsearchKubernetesSpec", func() {
	Context("with a valid spec", func() {
		It("should not return any validation errors", func() {
			spec := &ElasticsearchKubernetesSpec{
				ElasticsearchContainer: &ElasticsearchKubernetesElasticsearchContainer{
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
				KibanaContainer: &ElasticsearchKubernetesKibanaContainer{
					IsEnabled: true,
					Replicas:  1,
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
					DnsDomain: "elasticsearch.example.com",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors, got some")
		})
	})

	Context("when persistence is enabled but no disk_size is provided", func() {
		It("should return a validation error containing `[spec.container.disk_size.required]`", func() {
			spec := &ElasticsearchKubernetesSpec{
				ElasticsearchContainer: &ElasticsearchKubernetesElasticsearchContainer{
					Replicas:             1,
					IsPersistenceEnabled: true,
					// No disk_size provided
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
				KibanaContainer: &ElasticsearchKubernetesKibanaContainer{
					IsEnabled: true,
					Replicas:  1,
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
			Expect(err).NotTo(BeNil(), "Expected an error for missing disk_size but got none")
			Expect(err.Error()).To(ContainSubstring("[spec.container.disk_size.required]"),
				"Expected validation error with constraint id `spec.container.disk_size.required`")
		})
	})

	Context("when disk_size has an invalid format", func() {
		It("should return a validation error containing `[spec.container.disk_size.required]`", func() {
			spec := &ElasticsearchKubernetesSpec{
				ElasticsearchContainer: &ElasticsearchKubernetesElasticsearchContainer{
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
				KibanaContainer: &ElasticsearchKubernetesKibanaContainer{
					IsEnabled: true,
					Replicas:  1,
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
			Expect(err).NotTo(BeNil(), "Expected an error for invalid disk_size but got none")
			Expect(err.Error()).To(ContainSubstring("[spec.container.disk_size.required]"),
				"Expected validation error with constraint id `spec.container.disk_size.required`")
		})
	})

	Context("when persistence is disabled and disk_size is empty", func() {
		It("should not return any validation errors", func() {
			spec := &ElasticsearchKubernetesSpec{
				ElasticsearchContainer: &ElasticsearchKubernetesElasticsearchContainer{
					Replicas:             1,
					IsPersistenceEnabled: false,
					// disk_size is empty, but persistence is disabled, so no error expected
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
				KibanaContainer: &ElasticsearchKubernetesKibanaContainer{
					IsEnabled: true,
					Replicas:  1,
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
			Expect(err).To(BeNil(), "Did not expect a validation error when persistence is disabled and disk_size is empty")
		})
	})
})
