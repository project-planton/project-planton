package helmreleasev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestHelmRelease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HelmRelease Suite")
}

var _ = Describe("HelmRelease Custom Validation Tests", func() {
	var input *HelmRelease

	BeforeEach(func() {
		input = &HelmRelease{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "HelmRelease",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-helmrelease",
			},
			Spec: &HelmReleaseSpec{
				Repo:    "https://charts.helm.sh/stable",
				Name:    "nginx-ingress",
				Version: "1.41.3",
				Values: map[string]string{
					"someKey": "someValue",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("helmrelease", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
