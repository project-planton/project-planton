package awsroute53zonev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestAwsRoute53ZoneSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsRoute53ZoneSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsRoute53ZoneSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_route53_zone", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AwsRoute53Zone{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsRoute53Zone",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-route53-zone",
					},
					Spec: &AwsRoute53ZoneSpec{
						// No required fields, records is optional
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
