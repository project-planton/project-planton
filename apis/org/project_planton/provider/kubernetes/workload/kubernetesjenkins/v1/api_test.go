package kubernetesjenkinsv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/kubernetes"
)

func TestKubernetesJenkins(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesJenkins Suite")
}

var _ = ginkgo.Describe("KubernetesJenkins Custom Validation Tests", func() {
	var input *KubernetesJenkins

	ginkgo.BeforeEach(func() {
		input = &KubernetesJenkins{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesJenkins",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-jenkins",
			},
			Spec: &KubernetesJenkinsSpec{
				ContainerResources: &kubernetes.ContainerResources{
					Limits: &kubernetes.CpuMemory{
						Cpu:    "1000m",
						Memory: "1Gi",
					},
					Requests: &kubernetes.CpuMemory{
						Cpu:    "50m",
						Memory: "100Mi",
					},
				},
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "jenkins.example.com",
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("jenkins_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
