package openfgakubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestOpenFgaKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenFgaKubernetes Suite")
}

var _ = Describe("OpenFgaKubernetes Custom Validation Tests", func() {
	var input *OpenFgaKubernetes

	BeforeEach(func() {
		input = &OpenFgaKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "OpenFgaKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-openfga",
			},
			Spec: &OpenFgaKubernetesSpec{
				Container: &OpenFgaKubernetesContainer{
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
				},
				Ingress: nil,
				Datastore: &OpenFgaKubernetesDataStore{
					Engine: "postgres",
					Uri:    "postgres://user:pass@localhost:5432/testdb",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("openfga_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
