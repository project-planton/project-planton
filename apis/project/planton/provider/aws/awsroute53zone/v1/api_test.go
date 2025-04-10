package awsroute53zonev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/networking/enums/dnsrecordtype"
)

func TestAwsRoute53Zone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsRoute53Zone Suite")
}

var _ = Describe("AwsRoute53Zone Custom Validation Tests", func() {
	var input *AwsRoute53Zone

	BeforeEach(func() {
		input = &AwsRoute53Zone{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsRoute53Zone",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-zone",
			},
			Spec: &AwsRoute53ZoneSpec{
				Records: []*Route53DnsRecord{
					{
						RecordType: dnsrecordtype.DnsRecordType_A,
						Name:       "example.com",
						Values:     []string{"1.2.3.4"},
						TtlSeconds: 300,
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("aws-route53zone", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should accept a valid apex domain", func() {
				input.Spec.Records[0].Name = "example.com"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should accept a valid wildcard domain", func() {
				input.Spec.Records[0].Name = "*.example.com"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject a domain missing a TLD", func() {
				input.Spec.Records[0].Name = "example"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject multiple wildcard asterisks", func() {
				input.Spec.Records[0].Name = "**.example.com"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should reject a domain with invalid characters", func() {
				input.Spec.Records[0].Name = "exa@mple.com"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
