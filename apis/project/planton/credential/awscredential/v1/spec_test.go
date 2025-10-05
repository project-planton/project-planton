package awscredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("AwsCredentialSpec Validation Tests", func() {
	var (
		input          *AwsCredential
		validSecretKey = "ABCDEFGHIJ12345+/abcdeFGHIJ12345+/abcdef" // 40 chars
	)

	ginkgo.BeforeEach(func() {
		input = &AwsCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "AwsCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-aws-cred",
			},
			Spec: &AwsCredentialSpec{
				AccountId:       "123456789012",
				AccessKeyId:     "AKIAABCDEFGHIJKLMNOP",
				SecretAccessKey: "ABCDEFGHIJKLMNOPABCDEFGHIJKLMNOPABCDABCD",
				Region:          "us-west-2",
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
		ginkgo.Context("account_id validation", func() {

			ginkgo.It("should fail if account_id contains non-numeric characters", func() {
				input.Spec.AccountId = "12345ABC6789"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if account_id is missing", func() {
				input.Spec.AccountId = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("access_key_id validation", func() {

			ginkgo.It("should fail if access_key_id does not start with AKIA", func() {
				input.Spec.AccessKeyId = "BKIAABCDEFGHIJKLMNOP"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if access_key_id is shorter than 20", func() {
				input.Spec.AccessKeyId = "AKIAABCDE"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if access_key_id is longer than 20", func() {
				input.Spec.AccessKeyId = "AKIAABCDEFGHIJKLMNOPQRS"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if access_key_id has invalid characters", func() {
				input.Spec.AccessKeyId = "AKIAABC*EFGHIJKLMNOP"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if access_key_id is missing", func() {
				input.Spec.AccessKeyId = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("secret_access_key validation", func() {

			ginkgo.It("should fail if secret_access_key is shorter than 40", func() {
				input.Spec.SecretAccessKey = "ABC123"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if secret_access_key is longer than 40", func() {
				input.Spec.SecretAccessKey = validSecretKey + "AB"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if secret_access_key has invalid characters", func() {
				input.Spec.SecretAccessKey = "ABCDEFGHIJ12345@/abcdeFGHIJ12345+/abcd"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should fail if secret_access_key is missing", func() {
				input.Spec.SecretAccessKey = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
