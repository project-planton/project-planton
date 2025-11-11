package clickhousekubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared/kubernetes"
)

func TestClickHouseKubernetesSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ClickHouseKubernetesSpec Validation Suite")
}

var _ = ginkgo.Describe("ClickHouseKubernetesSpec validations", func() {
	var spec *ClickHouseKubernetesSpec

	ginkgo.BeforeEach(func() {
		spec = &ClickHouseKubernetesSpec{
			ClusterName: "test-cluster",
			Container: &ClickHouseKubernetesContainer{
				Replicas:           1,
				PersistenceEnabled: true,
				DiskSize:           "8Gi",
				Resources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "1000m",
						Memory: "2Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "100m",
						Memory: "256Mi",
					},
				},
			},
			Logging: &ClickHouseKubernetesLoggingConfig{
				Level: ClickHouseKubernetesLoggingConfig_information,
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("spec with persistence enabled", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with persistence disabled", func() {
			ginkgo.It("should not return a validation error when disk_size is empty", func() {
				spec.Container.PersistenceEnabled = false
				spec.Container.DiskSize = ""
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with clustering enabled", func() {
			ginkgo.It("should not return a validation error with valid shard and replica counts", func() {
				spec.Cluster = &ClickHouseKubernetesClusterConfig{
					IsEnabled:    true,
					ShardCount:   2,
					ReplicaCount: 2,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("spec with persistence enabled but no disk_size", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Container.PersistenceEnabled = true
				spec.Container.DiskSize = ""
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with invalid disk_size format", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Container.PersistenceEnabled = true
				spec.Container.DiskSize = "invalid"
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with zero replicas", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Container.Replicas = 0
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with clustering enabled but invalid counts", func() {
			ginkgo.It("should return a validation error when shard_count is zero", func() {
				spec.Cluster = &ClickHouseKubernetesClusterConfig{
					IsEnabled:    true,
					ShardCount:   0,
					ReplicaCount: 2,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error when replica_count is zero", func() {
				spec.Cluster = &ClickHouseKubernetesClusterConfig{
					IsEnabled:    true,
					ShardCount:   2,
					ReplicaCount: 0,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with clustering disabled", func() {
			ginkgo.It("should not validate cluster counts when clustering is disabled", func() {
				spec.Cluster = &ClickHouseKubernetesClusterConfig{
					IsEnabled:    false,
					ShardCount:   0,
					ReplicaCount: 0,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
