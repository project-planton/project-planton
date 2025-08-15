package awsstaticwebsitev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsStaticWebsite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsStaticWebsite Suite")
}

var _ = Describe("AwsStaticWebsite Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws", func() {
			var input *AwsStaticWebsite

			BeforeEach(func() {
				input = &AwsStaticWebsite{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsStaticWebsite",
					Metadata: &shared.ApiResourceMetadata{
						Name: "my-aws-website",
					},
					Spec: &AwsStaticWebsiteSpec{},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
