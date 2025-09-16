package cloudflarer2bucketv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareR2BucketSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareR2BucketSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("CloudflareR2BucketSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cloudflare_r2_bucket", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
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
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
