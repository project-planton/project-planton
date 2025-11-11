package cloudflared1databasev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestCloudflareD1DatabaseSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CloudflareD1DatabaseSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("CloudflareD1DatabaseSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("cloudflare_d1_database", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareD1Database{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareD1Database",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-d1-database",
					},
					Spec: &CloudflareD1DatabaseSpec{
						DatabaseName: "test-database",
						AccountId:    "test-account-123",
						Region:       CloudflareD1Region_weur,
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
