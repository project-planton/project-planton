package certmanagerkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCertManagerKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CertManagerKubernetes Suite")
}

var _ = Describe("CertManagerKubernetes Custom Validation Tests", func() {
	var input *CertManagerKubernetes

	BeforeEach(func() {
		input = &CertManagerKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "CertManagerKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-cert-manager",
			},
			Spec: &CertManagerKubernetesSpec{},
		}
	})

	Describe("When valid input is passed", func() {
		Context("cert_manager_kubernetes", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
