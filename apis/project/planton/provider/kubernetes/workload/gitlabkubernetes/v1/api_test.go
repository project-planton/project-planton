package gitlabkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGitlabKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GitlabKubernetes Suite")
}

var _ = ginkgo.Describe("GitlabKubernetes Custom Validation Tests", func() {
	var input *GitlabKubernetes

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gitlab_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
