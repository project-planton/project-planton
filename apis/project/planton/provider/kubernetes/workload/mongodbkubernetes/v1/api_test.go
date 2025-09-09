package mongodbkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestMongodbKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MongodbKubernetes Suite")
}

var _ = Describe("MongodbKubernetes Custom Validation Tests", func() {
	var input *MongodbKubernetes

	BeforeEach(func() {
		input = &MongodbKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "MongodbKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-mongo",
			},
			Spec: &MongodbKubernetesSpec{
				Container: &MongodbKubernetesContainer{
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
				Ingress: nil, // Omitted or set as needed; not part of custom validation tests
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("mongodb_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
