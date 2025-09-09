package gcpserviceaccountv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpServiceAccountSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpServiceAccountSpec Custom Validation Tests")
}

var _ = Describe("GcpServiceAccountSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_service_account", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpServiceAccount{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpServiceAccount",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-service-account",
					},
					Spec: &GcpServiceAccountSpec{
						ServiceAccountId: "test-sa-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
