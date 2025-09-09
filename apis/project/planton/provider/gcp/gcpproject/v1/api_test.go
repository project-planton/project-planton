package gcpprojectv1

import (
	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"testing"
)

func TestGcpProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpProject Suite")
}

var _ = Describe("GcpProject Custom Validation Tests", func() {

	var input *GcpProject

	BeforeEach(func() {
		input = &GcpProject{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpProject",
			Metadata: &shared.ApiResourceMetadata{
				Name: "a-test-name",
			},
			Spec: &GcpProjectSpec{
				ParentType:       GcpProjectParentType_organization,
				ParentId:         "123456789012",
				BillingAccountId: "0123AB-4567CD-89EFGH",
				Labels: map[string]string{
					"env": "dev",
				},
				DisableDefaultNetwork: true,
				EnabledApis: []string{
					"compute.googleapis.com",
					"storage.googleapis.com",
				},
				OwnerMember: "alice@example.com",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("GCP", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Custom Field Validations", func() {

		Context("billingAccountId pattern", func() {
			It("should accept a correctly formatted ID", func() {
				input.Spec.BillingAccountId = "ABCDEF-123456-7890AB"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject lowercase characters", func() {
				input.Spec.BillingAccountId = "abcdef-123456-7890ab"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject an ID missing dashes", func() {
				input.Spec.BillingAccountId = "ABCDEF1234567890AB"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("enabledApis pattern", func() {
			It("should accept valid service endpoints", func() {
				input.Spec.EnabledApis = []string{"compute.googleapis.com", "bigquery.googleapis.com"}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an entry without .googleapis.com suffix", func() {
				input.Spec.EnabledApis = []string{"compute.googleapis"}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject an entry with uppercase letters", func() {
				input.Spec.EnabledApis = []string{"Compute.googleapis.com"}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("ownerMember email", func() {
			It("should accept a valid email", func() {
				input.Spec.OwnerMember = "bob@example.com"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid email", func() {
				input.Spec.OwnerMember = "user:bob@example.com"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
