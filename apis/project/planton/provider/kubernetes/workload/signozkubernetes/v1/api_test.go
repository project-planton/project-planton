package signozkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSignozKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SignozKubernetes Suite")
}

var _ = ginkgo.Describe("SignozKubernetes Custom Validation Tests", func() {
	var input *SignozKubernetes

	ginkgo.BeforeEach(func() {
		input = &SignozKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1", // standard field
			Kind:       "SignozKubernetes",                  // standard field
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-signoz",
			},
			Spec: &SignozKubernetesSpec{},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})
})
