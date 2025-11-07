package signozkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	"google.golang.org/protobuf/proto"
)

func TestSignozKubernetesSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SignozKubernetesSpec Validation Suite")
}

var _ = ginkgo.Describe("SignozKubernetesSpec validations", func() {
	var spec *SignozKubernetesSpec

	ginkgo.BeforeEach(func() {
		spec = &SignozKubernetesSpec{
			SignozContainer: &SignozKubernetesContainer{
				Replicas: 1,
				Resources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "1000m",
						Memory: "2Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "200m",
						Memory: "512Mi",
					},
				},
			},
			OtelCollectorContainer: &SignozKubernetesContainer{
				Replicas: 2,
				Resources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "2000m",
						Memory: "4Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "500m",
						Memory: "1Gi",
					},
				},
			},
			Database: &SignozKubernetesDatabaseConfig{
				IsExternal: false,
				ManagedDatabase: &SignozKubernetesManagedClickhouse{
					Container: &SignozKubernetesClickhouseContainer{
						Replicas:           1,
						PersistenceEnabled: true,
						DiskSize:           "20Gi",
						Resources: &kubernetes.ContainerResources{
							Limits: &kubernetes.CpuMemory{
								Cpu:    "2000m",
								Memory: "4Gi",
							},
							Requests: &kubernetes.CpuMemory{
								Cpu:    "500m",
								Memory: "1Gi",
							},
						},
					},
					Cluster: &SignozKubernetesClickhouseCluster{
						IsEnabled: false,
					},
					Zookeeper: &SignozKubernetesZookeeperConfig{
						IsEnabled: false,
					},
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("spec with self-managed ClickHouse", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with self-managed ClickHouse and clustering enabled", func() {
			ginkgo.It("should not return a validation error with valid shard and replica counts", func() {
				spec.Database.ManagedDatabase.Cluster = &SignozKubernetesClickhouseCluster{
					IsEnabled:    true,
					ShardCount:   2,
					ReplicaCount: 2,
				}
				spec.Database.ManagedDatabase.Zookeeper = &SignozKubernetesZookeeperConfig{
					IsEnabled: true,
					Container: &SignozKubernetesZookeeperContainer{
						Replicas: 3,
						DiskSize: "8Gi",
						Resources: &kubernetes.ContainerResources{
							Limits: &kubernetes.CpuMemory{
								Cpu:    "500m",
								Memory: "512Mi",
							},
							Requests: &kubernetes.CpuMemory{
								Cpu:    "100m",
								Memory: "256Mi",
							},
						},
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with external ClickHouse", func() {
			ginkgo.It("should not return a validation error with valid connection details", func() {
				spec.Database = &SignozKubernetesDatabaseConfig{
					IsExternal: true,
					ExternalDatabase: &SignozKubernetesExternalClickhouse{
						Host:        "clickhouse.database.svc.cluster.local",
						HttpPort:    proto.Int32(8123),
						TcpPort:     proto.Int32(9000),
						ClusterName: proto.String("cluster"),
						IsSecure:    false,
						Username:    "signoz",
						Password:    "my-password",
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with ClickHouse persistence disabled", func() {
			ginkgo.It("should not return a validation error when disk_size is empty", func() {
				spec.Database.ManagedDatabase.Container.PersistenceEnabled = false
				spec.Database.ManagedDatabase.Container.DiskSize = ""
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("spec with zero replicas for SigNoz container", func() {
			ginkgo.It("should return a validation error", func() {
				spec.SignozContainer.Replicas = 0
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with zero replicas for OTel Collector", func() {
			ginkgo.It("should return a validation error", func() {
				spec.OtelCollectorContainer.Replicas = 0
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with zero replicas for ClickHouse", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database.ManagedDatabase.Container.Replicas = 0
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with external database but no connection details", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database = &SignozKubernetesDatabaseConfig{
					IsExternal: true,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with external database and invalid port", func() {
			ginkgo.It("should return a validation error for http_port", func() {
				spec.Database = &SignozKubernetesDatabaseConfig{
					IsExternal: true,
					ExternalDatabase: &SignozKubernetesExternalClickhouse{
						Host:        "clickhouse.database.svc.cluster.local",
						HttpPort:    proto.Int32(0),
						TcpPort:     proto.Int32(9000),
						ClusterName: proto.String("cluster"),
						Username:    "signoz",
						Password:    "my-password",
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for tcp_port", func() {
				spec.Database = &SignozKubernetesDatabaseConfig{
					IsExternal: true,
					ExternalDatabase: &SignozKubernetesExternalClickhouse{
						Host:        "clickhouse.database.svc.cluster.local",
						HttpPort:    proto.Int32(8123),
						TcpPort:     proto.Int32(0),
						ClusterName: proto.String("cluster"),
						Username:    "signoz",
						Password:    "my-password",
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with ClickHouse persistence enabled but no disk_size", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database.ManagedDatabase.Container.PersistenceEnabled = true
				spec.Database.ManagedDatabase.Container.DiskSize = ""
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with invalid ClickHouse disk_size format", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database.ManagedDatabase.Container.PersistenceEnabled = true
				spec.Database.ManagedDatabase.Container.DiskSize = "invalid"
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with ClickHouse clustering enabled but invalid counts", func() {
			ginkgo.It("should return a validation error when shard_count is zero", func() {
				spec.Database.ManagedDatabase.Cluster = &SignozKubernetesClickhouseCluster{
					IsEnabled:    true,
					ShardCount:   0,
					ReplicaCount: 2,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error when replica_count is zero", func() {
				spec.Database.ManagedDatabase.Cluster = &SignozKubernetesClickhouseCluster{
					IsEnabled:    true,
					ShardCount:   2,
					ReplicaCount: 0,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with ClickHouse clustering disabled", func() {
			ginkgo.It("should not validate cluster counts when clustering is disabled", func() {
				spec.Database.ManagedDatabase.Cluster = &SignozKubernetesClickhouseCluster{
					IsEnabled:    false,
					ShardCount:   0,
					ReplicaCount: 0,
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with Zookeeper invalid disk_size", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database.ManagedDatabase.Zookeeper = &SignozKubernetesZookeeperConfig{
					IsEnabled: true,
					Container: &SignozKubernetesZookeeperContainer{
						Replicas: 3,
						DiskSize: "invalid",
						Resources: &kubernetes.ContainerResources{
							Limits: &kubernetes.CpuMemory{
								Cpu:    "500m",
								Memory: "512Mi",
							},
							Requests: &kubernetes.CpuMemory{
								Cpu:    "100m",
								Memory: "256Mi",
							},
						},
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("spec with Zookeeper zero replicas", func() {
			ginkgo.It("should return a validation error", func() {
				spec.Database.ManagedDatabase.Zookeeper = &SignozKubernetesZookeeperConfig{
					IsEnabled: true,
					Container: &SignozKubernetesZookeeperContainer{
						Replicas: 0,
						DiskSize: "8Gi",
						Resources: &kubernetes.ContainerResources{
							Limits: &kubernetes.CpuMemory{
								Cpu:    "500m",
								Memory: "512Mi",
							},
							Requests: &kubernetes.CpuMemory{
								Cpu:    "100m",
								Memory: "256Mi",
							},
						},
					},
				}
				err := protovalidate.Validate(spec)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
