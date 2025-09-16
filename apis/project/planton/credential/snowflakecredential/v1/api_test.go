package snowflakecredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSnowflakeCredential(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SnowflakeCredential Suite")
}

var _ = ginkgo.Describe("SnowflakeCredentialSpec Custom Validation Tests", func() {
	var input *SnowflakeCredential

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with correct api_version and kind", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("with incorrect api_version", func() {
			ginkgo.It("should return a validation error", func() {
				input.ApiVersion = "invalid-value"
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		ginkgo.Context("with incorrect kind", func() {
			ginkgo.It("should return a validation error", func() {
				input.Kind = "NotSnowflakeCredential"
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
})
