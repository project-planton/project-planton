package awsstaticwebsitev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsStaticWebsiteSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsStaticWebsiteSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsStaticWebsiteSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_static_website", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AwsStaticWebsite{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsStaticWebsite",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-static-website",
					},
					Spec: &AwsStaticWebsiteSpec{
						// No required fields, just use defaults
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
