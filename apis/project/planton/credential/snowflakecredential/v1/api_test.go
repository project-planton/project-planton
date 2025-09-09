package snowflakecredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSnowflakeCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SnowflakeCredential Suite")
}

var _ = Describe("SnowflakeCredentialSpec Custom Validation Tests", func() {
	var input *SnowflakeCredential

	BeforeEach(func() {
		input = &SnowflakeCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "SnowflakeCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-snowflake-cred",
			},
			Spec: &SnowflakeCredentialSpec{
				Account:  "my_snowflake_account",
				Username: "my_user",
				Password: "my_password",
				Region:   "us-west",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("with correct api_version and kind", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("When invalid input is passed", func() {
		Context("with incorrect api_version", func() {
			It("should return a validation error", func() {
				input.ApiVersion = "invalid-value"
				err := protovalidate.Validate(input)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("with incorrect kind", func() {
			It("should return a validation error", func() {
				input.Kind = "NotSnowflakeCredential"
				err := protovalidate.Validate(input)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
