package locustkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestLocustKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LocustKubernetes Suite")
}

var _ = Describe("LocustKubernetes Custom Validation Tests", func() {
	var input *LocustKubernetes

	BeforeEach(func() {
		input = &LocustKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "LocustKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "sample-locust",
			},
			Spec: &LocustKubernetesSpec{
				MasterContainer: &LocustKubernetesContainer{
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
				WorkerContainer: &LocustKubernetesContainer{
					Replicas: 2,
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
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "locust.example.com",
				},
				LoadTest: &LocustKubernetesLoadTest{
					Name:          "example-loadtest",
					MainPyContent: "print('Hello, Locust')",
					LibFilesContent: map[string]string{
						"utils.py": "def helper(): pass",
					},
					PipPackages: []string{"requests", "locust"},
				},
				HelmValues: map[string]string{
					"someKey": "someValue",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("locust_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
