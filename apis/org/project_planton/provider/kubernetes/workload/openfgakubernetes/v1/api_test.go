package openfgakubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestOpenFgaKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "OpenFgaKubernetes Suite")
}

var _ = ginkgo.Describe("OpenFgaKubernetes Custom Validation Tests", func() {
	var input *OpenFgaKubernetes

	ginkgo.BeforeEach(func() {
		input = &OpenFgaKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "OpenFgaKubernetes",
			Metadata: &shared.CloudResourceMetadata{
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
				Ingress: &OpenFgaKubernetesIngress{
					Enabled:  true,
					Hostname: "test-openfga.example.com",
				},
				Datastore: &OpenFgaKubernetesDataStore{
					Engine: "postgres",
					Uri:    "postgres://user:pass@localhost:5432/testdb",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("openfga_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
