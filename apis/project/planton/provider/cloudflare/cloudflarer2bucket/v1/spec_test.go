package cloudflarer2bucketv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareR2BucketSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareR2BucketSpec Custom Validation Tests")
}

var _ = Describe("CloudflareR2BucketSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_r2_bucket", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareR2Bucket{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareR2Bucket",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-r2-bucket",
					},
					Spec: &CloudflareR2BucketSpec{
						BucketName: "test-bucket",
						Location:   CloudflareR2Location_WEUR,
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
