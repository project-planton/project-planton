package argocdkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestArgocdKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ArgocdKubernetes Suite")
}

var _ = Describe("ArgocdKubernetes Custom Validation Tests", func() {
	var input *ArgocdKubernetes

	BeforeEach(func() {
		input = &ArgocdKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ArgocdKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-argocd",
			},
			Spec: &ArgocdKubernetesSpec{
				Container: &ArgocdKubernetesArgocdContainer{},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("argocd_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
