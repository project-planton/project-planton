package kuberneteskafkav1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesKafka(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesKafka Suite")
}

var _ = ginkgo.Describe("KubernetesKafka Custom Validation Tests", func() {
	var input *KubernetesKafka

	ginkgo.BeforeEach(func() {
		input = &KubernetesKafka{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesKafka",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-kafka",
			},
			Spec: &KubernetesKafkaSpec{
				KafkaTopics: []*KafkaTopic{
					{
						Name: "my-topic1",
					},
					{
						Name: "another_topic-1",
					},
				},
				BrokerContainer: &KubernetesKafkaBrokerContainer{
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
				ZookeeperContainer: &KubernetesKafkaZookeeperContainer{
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
				SchemaRegistryContainer: &KubernetesKafkaSchemaRegistryContainer{
					IsEnabled: false,
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("kafka_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
