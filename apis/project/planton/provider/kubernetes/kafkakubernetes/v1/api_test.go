package kafkakubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestKafkaKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KafkaKubernetesSpec Suite")
}

var _ = Describe("KafkaKubernetesSpec", func() {
	Context("when the spec is fully valid", func() {
		It("should pass validation without errors", func() {
			spec := &KafkaKubernetesSpec{
				KafkaTopics: []*KafkaTopic{
					{
						Name:       "validTopicName",
						Partitions: 3,
						Replicas:   3,
						Config: map[string]string{
							"cleanup.policy": "compact",
						},
					},
				},
				BrokerContainer: &KafkaKubernetesBrokerContainer{
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
					DiskSize: "10Gi",
				},
				ZookeeperContainer: &KafkaKubernetesZookeeperContainer{
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
					DiskSize: "10Gi",
				},
				SchemaRegistryContainer: &KafkaKubernetesSchemaRegistryContainer{
					IsEnabled: true,
					Replicas:  1,
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
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "kafka.example.com",
				},
				IsDeployKafkaUi: true,
			}
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors, got some")
		})
	})

	Context("when the broker disk size is invalid", func() {
		It("should fail validation with `[spec.broker_container.disk_size.format]`", func() {
			spec := &KafkaKubernetesSpec{
				BrokerContainer: &KafkaKubernetesBrokerContainer{
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
					DiskSize: "invalidDiskSize", // Invalid format
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "Expected validation error for invalid broker disk size")
			Expect(err.Error()).To(ContainSubstring("[spec.broker_container.disk_size.format]"),
				"Expected validation error with constraint id `spec.broker_container.disk_size.format`")
		})
	})

	Context("when the Zookeeper disk size is invalid", func() {
		It("should fail validation with `[spec.broker_container.disk_size.format]`", func() {
			spec := &KafkaKubernetesSpec{
				ZookeeperContainer: &KafkaKubernetesZookeeperContainer{
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
					DiskSize: "123abc", // Invalid format
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).NotTo(BeNil(), "Expected validation error for invalid zookeeper disk size")
			Expect(err.Error()).To(ContainSubstring("[spec.broker_container.disk_size.format]"),
				"Expected validation error with constraint id `spec.broker_container.disk_size.format`")
		})
	})

	Context("when the topic name is invalid", func() {
		When("it starts with a non-alphanumeric character", func() {
			It("should fail validation with a message about starting with an alphanumeric char", func() {
				spec := &KafkaKubernetesSpec{
					KafkaTopics: []*KafkaTopic{
						{
							Name:       ".invalidName", // starts with non-alphanumeric
							Partitions: 1,
							Replicas:   1,
						},
					},
				}
				err := protovalidate.Validate(spec)
				Expect(err).NotTo(BeNil(), "Expected validation error for invalid topic name")
				Expect(err.Error()).To(ContainSubstring("Should start with an alphanumeric character"),
					"Expected error message about topic name start")
			})
		})

		When("it ends with a non-alphanumeric character", func() {
			It("should fail validation with a message about ending with alphanumeric", func() {
				spec := &KafkaKubernetesSpec{
					KafkaTopics: []*KafkaTopic{
						{
							Name:       "validName-",
							Partitions: 1,
							Replicas:   1,
						},
					},
				}
				err := protovalidate.Validate(spec)
				Expect(err).NotTo(BeNil(), "Expected validation error for topic name ending")
				Expect(err.Error()).To(ContainSubstring("Should end with an alphanumeric character"),
					"Expected error about ending with alphanumeric")
			})
		})

		When("it contains non-ASCII characters", func() {
			It("should fail validation with a message about non-ASCII characters", func() {
				spec := &KafkaKubernetesSpec{
					KafkaTopics: []*KafkaTopic{
						{
							Name:       "invalidName✓",
							Partitions: 1,
							Replicas:   1,
						},
					},
				}
				err := protovalidate.Validate(spec)
				Expect(err).NotTo(BeNil(), "Expected validation error for non-ASCII character in topic name")
				Expect(err.Error()).To(ContainSubstring("Must not contain non-ASCII characters"),
					"Expected error about non-ASCII")
			})
		})

		When("it contains '..'", func() {
			It("should fail validation with a message about '..'", func() {
				spec := &KafkaKubernetesSpec{
					KafkaTopics: []*KafkaTopic{
						{
							Name:       "invalid..name",
							Partitions: 1,
							Replicas:   1,
						},
					},
				}
				err := protovalidate.Validate(spec)
				Expect(err).NotTo(BeNil(), "Expected validation error for '..' in topic name")
				Expect(err.Error()).To(ContainSubstring("Must not contain '..'"),
					"Expected error about '..'")
			})
		})

		When("it contains invalid characters", func() {
			It("should fail validation with a message about allowed characters", func() {
				spec := &KafkaKubernetesSpec{
					KafkaTopics: []*KafkaTopic{
						{
							Name:       "invalid#name",
							Partitions: 1,
							Replicas:   1,
						},
					},
				}
				err := protovalidate.Validate(spec)
				Expect(err).NotTo(BeNil(), "Expected validation error for invalid characters in topic name")
				Expect(err.Error()).To(ContainSubstring("Only alphanumeric and ('.', '_' and '-') characters are allowed"),
					"Expected error about allowed characters")
			})
		})
	})

	Context("when schema registry is disabled", func() {
		It("should still pass validation", func() {
			spec := &KafkaKubernetesSpec{
				BrokerContainer: &KafkaKubernetesBrokerContainer{
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
				ZookeeperContainer: &KafkaKubernetesZookeeperContainer{
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
					DiskSize: "1Gi",
				},
				SchemaRegistryContainer: &KafkaKubernetesSchemaRegistryContainer{
					IsEnabled: false,
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors when schema registry is disabled")
		})
	})

	Context("when KafkaTopics is empty", func() {
		It("should pass validation with an empty list of topics", func() {
			spec := &KafkaKubernetesSpec{
				BrokerContainer: &KafkaKubernetesBrokerContainer{
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
					DiskSize: "10Gi",
				},
				ZookeeperContainer: &KafkaKubernetesZookeeperContainer{
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
					DiskSize: "10Gi",
				},
			}
			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "Expected no validation errors with empty KafkaTopics")
		})
	})
})
