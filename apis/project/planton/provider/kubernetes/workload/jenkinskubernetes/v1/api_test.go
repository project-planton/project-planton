package jenkinskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestJenkinsKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JenkinsKubernetes Suite")
}

var _ = Describe("JenkinsKubernetes Custom Validation Tests", func() {
	var input *JenkinsKubernetes

	BeforeEach(func() {
		input = &JenkinsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "JenkinsKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-jenkins",
			},
			Spec: &JenkinsKubernetesSpec{
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

	Describe("When valid input is passed", func() {
		Context("jenkins_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
