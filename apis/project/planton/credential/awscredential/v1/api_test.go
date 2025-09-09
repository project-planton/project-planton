package awscredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsCredential Suite")
}

var _ = Describe("AwsCredentialSpec Custom Validation Tests", func() {
	var (
		input          *AwsCredential
		validSecretKey = "ABCDEFGHIJ12345+/abcdeFGHIJ12345+/abcdef" // 40 chars
	)

	BeforeEach(func() {
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

	Context("when valid input is passed", func() {
		It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Describe("account_id custom numeric check", func() {
		Context("when account_id is purely numeric", func() {
			It("should pass validation", func() {
				input.Spec.AccountId = "9876543210"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when account_id contains non-numeric characters", func() {
			It("should fail validation", func() {
				input.Spec.AccountId = "12345ABC6789"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("access_key_id format rules", func() {
		Context("when access_key_id is valid", func() {
			It("should pass validation", func() {
				input.Spec.AccessKeyId = "AKIAABCDEFGHIJKLMNOP"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when access_key_id does not start with AKIA", func() {
			It("should fail validation", func() {
				input.Spec.AccessKeyId = "BKIAABCDEFGHIJKLMNOP"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when access_key_id is not exactly 20 characters", func() {
			It("should fail if shorter than 20", func() {
				input.Spec.AccessKeyId = "AKIAABCDE"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should fail if longer than 20", func() {
				input.Spec.AccessKeyId = "AKIAABCDEFGHIJKLMNOPQRS"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when access_key_id has invalid characters in the last 16", func() {
			It("should fail validation", func() {
				input.Spec.AccessKeyId = "AKIAABC*EFGHIJKLMNOP"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("secret_access_key character set rules", func() {
		Context("when secret_access_key is valid", func() {
			It("should pass validation", func() {
				input.Spec.SecretAccessKey = validSecretKey
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when secret_access_key length is incorrect", func() {
			It("should fail if shorter than 40", func() {
				input.Spec.SecretAccessKey = "ABC123"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should fail if longer than 40", func() {
				input.Spec.SecretAccessKey = validSecretKey + "AB"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when secret_access_key has invalid characters", func() {
			It("should fail validation", func() {
				input.Spec.SecretAccessKey = "ABCDEFGHIJ12345@/abcdeFGHIJ12345+/abcd"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
