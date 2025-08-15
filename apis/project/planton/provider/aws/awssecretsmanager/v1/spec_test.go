package awssecretsmanagerv1

import (
	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsSecretsManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsSecretsManager Suite")
}

var _ = Describe("AwsSecretsManager Custom Validation Tests", func() {

	var input *AwsSecretsManager

	BeforeEach(func() {
		input = &AwsSecretsManager{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsSecretsManager",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-secret-manager",
			},
			Spec: &AwsSecretsManagerSpec{
				SecretNames: []string{"my-secret-1", "another-secret"},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("aws", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
