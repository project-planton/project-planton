package awscredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsCredential(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsCredential Suite")
}

var (
	_ = ginkgo.Describe("AwsCredentialSpec Custom Validation Tests", func() {
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

		ginkgo.Context("when valid input is passed", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Describe("account_id custom numeric check", func() {
			ginkgo.Context("when account_id is purely numeric", func() {
				ginkgo.It("should pass validation", func() {
					input.Spec.AccountId = "9876543210"
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})
			})

			ginkgo.Context("when account_id contains non-numeric characters", func() {
				ginkgo.It("should fail validation", func() {
					input.Spec.AccountId = "12345ABC6789"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})
		})

		ginkgo.Describe("access_key_id format rules", func() {
			ginkgo.Context("when access_key_id is valid", func() {
				ginkgo.It("should pass validation", func() {
					input.Spec.AccessKeyId = "AKIAABCDEFGHIJKLMNOP"
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})
			})

			ginkgo.Context("when access_key_id does not start with AKIA", func() {
				ginkgo.It("should fail validation", func() {
					input.Spec.AccessKeyId = "BKIAABCDEFGHIJKLMNOP"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})

			ginkgo.Context("when access_key_id is not exactly 20 characters", func() {
				ginkgo.It("should fail if shorter than 20", func() {
					input.Spec.AccessKeyId = "AKIAABCDE"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})

				ginkgo.It("should fail if longer than 20", func() {
					input.Spec.AccessKeyId = "AKIAABCDEFGHIJKLMNOPQRS"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})

			ginkgo.Context("when access_key_id has invalid characters in the last 16", func() {
				ginkgo.It("should fail validation", func() {
					input.Spec.AccessKeyId = "AKIAABC*EFGHIJKLMNOP"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})
		})

		ginkgo.Describe("secret_access_key character set rules", func() {
			ginkgo.Context("when secret_access_key is valid", func() {
				ginkgo.It("should pass validation", func() {
					input.Spec.SecretAccessKey = validSecretKey
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})
			})

			ginkgo.Context("when secret_access_key length is incorrect", func() {
				ginkgo.It("should fail if shorter than 40", func() {
					input.Spec.SecretAccessKey = "ABC123"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})

				ginkgo.It("should fail if longer than 40", func() {
					input.Spec.SecretAccessKey = validSecretKey + "AB"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})

			ginkgo.Context("when secret_access_key has invalid characters", func() {
				ginkgo.It("should fail validation", func() {
					input.Spec.SecretAccessKey = "ABCDEFGHIJ12345@/abcdeFGHIJ12345+/abcd"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})
			})
		})
	})
)
