package solrkubernetesv1

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestSolrKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolrKubernetesSpec Suite")
}

var _ = Describe("SolrKubernetesSpec", func() {
	Context("with a fully valid spec", func() {
		It("should pass validation with no errors", func() {
			spec := &SolrKubernetesSpec{
				SolrContainer: &SolrKubernetesSolrContainer{
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
					DiskSize: "1Gi", // Valid disk size format
					Image: &kubernetes.ContainerImage{
						Repo: "solr",
						Tag:  "8.7.0",
					},
				},
				Config: &SolrKubernetesSolrConfig{
					JavaMem:                 "-Xmx512m",
					Opts:                    "-Dsolr.autoSoftCommit.maxTime=10000",
					GarbageCollectionTuning: "-XX:SurvivorRatio=4",
				},
				ZookeeperContainer: &SolrKubernetesZookeeperContainer{
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
					DiskSize: "1Gi", // Valid disk size
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "solr.example.com",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when the Solr container disk size is invalid", func() {
		It("should fail validation, expecting constraint id `spec.container.disk_size.required`", func() {
			spec := &SolrKubernetesSpec{
				SolrContainer: &SolrKubernetesSolrContainer{
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
					DiskSize: "abc", // Invalid format
					Image: &kubernetes.ContainerImage{
						Repo: "solr",
						Tag:  "8.7.0",
					},
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for invalid disk size")
			Expect(strings.Contains(err.Error(), "[spec.container.disk_size.required]")).
				To(BeTrue(), "expected constraint id `spec.container.disk_size.required` in error message")
		})
	})

	Context("when the Zookeeper container disk size is invalid", func() {
		It("should fail validation, expecting constraint id `spec.container.disk_size.required`", func() {
			spec := &SolrKubernetesSpec{
				ZookeeperContainer: &SolrKubernetesZookeeperContainer{
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
					DiskSize: "100", // Missing unit, invalid format
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "expected validation error for invalid disk size")
			Expect(strings.Contains(err.Error(), "[spec.container.disk_size.required]")).
				To(BeTrue(), "expected constraint id `spec.container.disk_size.required` in error message")
		})
	})

	Context("when the SolrContainer is missing entirely", func() {
		It("should not return an error if SolrContainer is optional (adjust if required)", func() {
			spec := &SolrKubernetesSpec{
				ZookeeperContainer: &SolrKubernetesZookeeperContainer{
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
			}

			// If SolrContainer is optional, we expect no error. If it's required, adjust your expectation:
			// e.g., Expect(err).NotTo(BeNil()) if it must always be present.
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "did not expect a validation error when SolrContainer is missing (assuming optional)")
		})
	})
})
