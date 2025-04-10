package awsalbv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsAlbSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsAlbSpec Custom Validation Tests")
}

var _ = Describe("AwsAlbSpec Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_alb", func() {
			It("should not return a validation error", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-alb-resource",
					},
					Spec: &AwsAlbSpec{
						Subnets: []string{"subnet-abc123"},
					},
				}

				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
