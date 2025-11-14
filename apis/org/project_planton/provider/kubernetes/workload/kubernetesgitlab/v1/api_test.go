package kubernetesgitlabv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestKubernetesGitlab(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesGitlab Suite")
}

var _ = ginkgo.Describe("KubernetesGitlab Custom Validation Tests", func() {
	var input *KubernetesGitlab

	ginkgo.BeforeEach(func() {
		input = &KubernetesGitlab{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesGitlab",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-gitlab",
			},
			Spec: &KubernetesGitlabSpec{
				Container: &KubernetesGitlabSpecContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gitlab_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
