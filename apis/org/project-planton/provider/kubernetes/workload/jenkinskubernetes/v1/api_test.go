package jenkinskubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared/kubernetes"
)

func TestJenkinsKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "JenkinsKubernetes Suite")
}

var _ = ginkgo.Describe("JenkinsKubernetes Custom Validation Tests", func() {
	var input *JenkinsKubernetes

	ginkgo.BeforeEach(func() {
		input = &JenkinsKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "JenkinsKubernetes",
			Metadata: &shared.CloudResourceMetadata{
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("jenkins_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
