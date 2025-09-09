package awssecretsmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsSecretsManagerSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsSecretsManagerSpec Custom Validation Tests")
}

var _ = Describe("AwsSecretsManagerSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_secrets_manager", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsSecretsManager{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsSecretsManager",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &AwsSecretsManagerSpec{
						SecretNames: []string{"test-secret"},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
