package kuberneteslocustv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesLocust(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesLocust Suite")
}

var _ = ginkgo.Describe("KubernetesLocust Custom Validation Tests", func() {
	var input *KubernetesLocust

	ginkgo.BeforeEach(func() {
		input = &KubernetesLocust{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesLocust",
			Metadata: &shared.CloudResourceMetadata{
				Name: "sample-locust",
			},
			Spec: &KubernetesLocustSpec{
				MasterContainer: &KubernetesLocustContainer{
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
				WorkerContainer: &KubernetesLocustContainer{
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
				LoadTest: &KubernetesLocustLoadTest{
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("locust_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
