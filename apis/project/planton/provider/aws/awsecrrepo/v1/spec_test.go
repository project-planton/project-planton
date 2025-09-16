package awsecrrepov1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
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
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsEcrRepoSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsEcrRepoSpec Custom Validation Tests", func() {
	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_ecr_repo", func() {
			var input *AwsEcrRepo

			ginkgo.BeforeEach(func() {
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

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.Context("repository_name length constraints", func() {
				ginkgo.It("should fail if repository_name is too short", func() {
					input.Spec.RepositoryName = "a"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})

				ginkgo.It("should fail if repository_name is too long", func() {
					input.Spec.RepositoryName = stringOfLength(257)
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})

				ginkgo.It("should succeed if repository_name length is within the valid range", func() {
					input.Spec.RepositoryName = "valid-repo-123"
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})
			})

			ginkgo.Context("encryption_type constraints", func() {
				ginkgo.It("should fail if encryption_type is invalid", func() {
					input.Spec.EncryptionType = "INVALID"
					err := protovalidate.Validate(input)
					gomega.Expect(err).NotTo(gomega.BeNil())
				})

				ginkgo.It("should succeed if encryption_type is AES256", func() {
					input.Spec.EncryptionType = "AES256"
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})

				ginkgo.It("should succeed if encryption_type is KMS", func() {
					input.Spec.EncryptionType = "KMS"
					err := protovalidate.Validate(input)
					gomega.Expect(err).To(gomega.BeNil())
				})
			})
		})
	})
})
