package helmreleasev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestHelmRelease(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "HelmRelease Suite")
}

var _ = ginkgo.Describe("HelmRelease Custom Validation Tests", func() {
	var input *HelmRelease

	ginkgo.BeforeEach(func() {
		input = &HelmRelease{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "HelmRelease",
			Metadata: &shared.CloudResourceMetadata{
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("helmrelease", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
