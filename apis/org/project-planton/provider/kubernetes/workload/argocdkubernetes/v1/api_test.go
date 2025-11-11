package argocdkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestArgocdKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ArgocdKubernetes Suite")
}

var _ = ginkgo.Describe("ArgocdKubernetes Custom Validation Tests", func() {
	var input *ArgocdKubernetes

	ginkgo.BeforeEach(func() {
		input = &ArgocdKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "ArgocdKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "test-argocd",
			},
			Spec: &ArgocdKubernetesSpec{
				Container: &ArgocdKubernetesArgocdContainer{},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("argocd_kubernetes", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
