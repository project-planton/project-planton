package snowflakecredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestSnowflakeCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SnowflakeCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("SnowflakeCredentialSpec Validation Tests", func() {
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
		ginkgo.Context("with valid credentials", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("missing required fields", func() {

			ginkgo.It("should return error if account is missing", func() {
				input.Spec.Account = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if username is missing", func() {
				input.Spec.Username = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if password is missing", func() {
				input.Spec.Password = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
