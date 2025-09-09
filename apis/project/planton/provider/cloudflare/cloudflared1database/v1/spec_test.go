package cloudflared1databasev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareD1DatabaseSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareD1DatabaseSpec Custom Validation Tests")
}

var _ = Describe("CloudflareD1DatabaseSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_d1_database", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareD1Database{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareD1Database",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-d1-database",
					},
					Spec: &CloudflareD1DatabaseSpec{
						DatabaseName: "test-database",
						AccountId:    "test-account-123",
						Region:       CloudflareD1Region_weur,
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
