// gcpserviceaccount_custom_validation_test.go
package gcpserviceaccountv1

import (
	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"testing"
)

func TestGcpServiceAccount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpServiceAccount Custom Validation Suite")
}

var _ = Describe("GcpServiceAccount Custom Validation Tests", func() {

	var input *GcpServiceAccount

	BeforeEach(func() {
		input = &GcpServiceAccount{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpServiceAccount",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-service-account",
			},
			Spec: &GcpServiceAccountSpec{
				ServiceAccountId: "service1", // 8 characters -- within 6-30 limit
				ProjectId:        "my-gcp-project",
				CreateKey:        false,
			},
		}
	})

	Describe("When valid input is passed", func() {

		Context("GCP provider", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Context("ServiceAccountId Length Validations", func() {

		It("should fail when service_account_id is shorter than 6 characters", func() {
			input.Spec.ServiceAccountId = "svc1" // 4 characters
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should fail when service_account_id is longer than 30 characters", func() {
			input.Spec.ServiceAccountId = "this-is-a-very-long-service-account-id"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should succeed when service_account_id length is within 6-30 characters", func() {
			input.Spec.ServiceAccountId = "validsa" // 7 characters
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})
})
