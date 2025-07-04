package gitlabkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGitlabKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitlabKubernetes Suite")
}

var _ = Describe("GitlabKubernetes Custom Validation Tests", func() {
	var input *GitlabKubernetes

	BeforeEach(func() {
		input = &GitlabKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "GitlabKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-gitlab",
			},
			Spec: &GitlabKubernetesSpec{
				Container: &GitlabKubernetesSpecContainer{},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gitlab_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
