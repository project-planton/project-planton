package mongodbkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestMongodbKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "MongodbKubernetes Suite")
}

var _ = ginkgo.Describe("MongodbKubernetes Custom Validation Tests", func() {
	var input *MongodbKubernetes

	ginkgo.BeforeEach(func() {
		input = &MongodbKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "MongodbKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-mongo",
			},
			Spec: &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
					Replicas:           1,
					PersistenceEnabled: true,
					DiskSize:           "10Gi", // valid format
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
				Ingress: &MongodbKubernetesIngress{
					Enabled: false,
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("mongodb_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
