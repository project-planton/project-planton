package locustkubernetesv1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bufbuild/protovalidate-go"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestLocustKubernetesSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LocustKubernetesSpec Suite")
}

var _ = Describe("LocustKubernetesSpec", func() {
	Context("when the spec is fully valid", func() {
		It("should pass validation without errors", func() {
			spec := &LocustKubernetesSpec{
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
					Replicas: 5,
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
					Name:          "my_load_test",
					MainPyContent: "from locust import HttpUser, task",
					LibFilesContent: map[string]string{
						"utils.py": "def helper(): pass",
					},
					PipPackages: []string{"requests", "locustio"},
				},
				HelmValues: map[string]string{
					"image.tag": "latest",
				},
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors")
		})
	})

	Context("when helm_values is not provided", func() {
		It("should pass validation if it’s optional", func() {
			spec := &LocustKubernetesSpec{
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
				LoadTest: &LocustKubernetesLoadTest{
					Name:          "my_test",
					MainPyContent: "from locust import HttpUser, task",
					LibFilesContent: map[string]string{
						"utils.py": "def helper(): pass",
					},
				},
				// No HelmValues provided, which should be valid if optional
			}

			err := protovalidate.Validate(spec)
			Expect(err).To(BeNil(), "expected no validation errors without helm_values")
		})
	})
})
