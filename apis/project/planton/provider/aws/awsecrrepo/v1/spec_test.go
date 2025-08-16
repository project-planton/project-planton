package awsecrrepov1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

// stringOfLength is a helper for test scenarios needing an overly long string.
func stringOfLength(n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = 'a'
	}
	return string(s)
}

func TestAwsEcrRepoSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsEcrRepoSpec Custom Validation Tests")
}

var _ = Describe("AwsEcrRepoSpec Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_ecr_repo", func() {
			var input *AwsEcrRepo

			BeforeEach(func() {
				input = &AwsEcrRepo{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEcrRepo",
					Metadata: &shared.ApiResourceMetadata{
						Name: "valid-repo-metadata",
					},
					Spec: &AwsEcrRepoSpec{
						RepositoryName: "my-valid-repo",
						ImageImmutable: true,
						EncryptionType: "AES256",
						ForceDelete:    false,
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			Context("repository_name length constraints", func() {
				It("should fail if repository_name is too short", func() {
					input.Spec.RepositoryName = "a"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})

				It("should fail if repository_name is too long", func() {
					input.Spec.RepositoryName = stringOfLength(257)
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})

				It("should succeed if repository_name length is within the valid range", func() {
					input.Spec.RepositoryName = "valid-repo-123"
					err := protovalidate.Validate(input)
					Expect(err).To(BeNil())
				})
			})

			Context("encryption_type constraints", func() {
				It("should fail if encryption_type is invalid", func() {
					input.Spec.EncryptionType = "INVALID"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})

				It("should succeed if encryption_type is AES256", func() {
					input.Spec.EncryptionType = "AES256"
					err := protovalidate.Validate(input)
					Expect(err).To(BeNil())
				})

				It("should succeed if encryption_type is KMS", func() {
					input.Spec.EncryptionType = "KMS"
					err := protovalidate.Validate(input)
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
