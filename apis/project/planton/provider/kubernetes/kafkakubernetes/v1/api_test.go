package kafkakubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestKafkaKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KafkaKubernetes Suite")
}

var _ = Describe("KafkaKubernetes Custom Validation Tests", func() {
	var input *KafkaKubernetes

	BeforeEach(func() {
		input = &KafkaKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KafkaKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-kafka",
			},
			Spec: &KafkaKubernetesSpec{
				KafkaTopics: []*KafkaTopic{
					{
						Name: "my-topic1",
					},
					{
						Name: "another_topic-1",
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
				SchemaRegistryContainer: &KafkaKubernetesSchemaRegistryContainer{
					IsEnabled: false,
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("kafka_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
