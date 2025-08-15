package awssecuritygroupv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
)

func TestAwsSecurityGroup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsSecurityGroup Suite")
}

var _ = Describe("AwsSecurityGroup Custom Validation Tests", func() {

	var input *AwsSecurityGroup

	BeforeEach(func() {
		input = &AwsSecurityGroup{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsSecurityGroup",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-sg",
			},
			Spec: &AwsSecurityGroupSpec{
				VpcId: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "vpc-12345678"},
				},
				Description: "Valid SG description",
				Ingress: []*SecurityGroupRule{
					{
						Protocol:    "tcp",
						FromPort:    80,
						ToPort:      80,
						Ipv4Cidrs:   []string{"0.0.0.0/0"},
						Description: "Allow HTTP inbound",
					},
				},
				Egress: []*SecurityGroupRule{
					{
						Protocol:    "tcp",
						FromPort:    443,
						ToPort:      443,
						Ipv4Cidrs:   []string{"0.0.0.0/0"},
						Description: "Allow HTTPS outbound",
					},
				},
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

	Context("Spec Description Validations", func() {
		It("should accept a description of exactly 255 characters", func() {
			longDesc := make([]byte, 255)
			for i := range longDesc {
				longDesc[i] = 'a'
			}
			input.Spec.Description = string(longDesc)
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should reject a description exceeding 255 characters", func() {
			tooLongDesc := make([]byte, 256)
			for i := range tooLongDesc {
				tooLongDesc[i] = 'a'
			}
			input.Spec.Description = string(tooLongDesc)
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Rule Description Validations", func() {
		It("should accept a rule description of exactly 255 characters", func() {
			longDesc := make([]byte, 255)
			for i := range longDesc {
				longDesc[i] = 'a'
			}
			input.Spec.Ingress[0].Description = string(longDesc)
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should reject a rule description exceeding 255 characters", func() {
			tooLongDesc := make([]byte, 256)
			for i := range tooLongDesc {
				tooLongDesc[i] = 'a'
			}
			input.Spec.Ingress[0].Description = string(tooLongDesc)
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})
})
