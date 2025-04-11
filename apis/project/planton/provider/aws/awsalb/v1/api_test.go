package awsalbv1

import (
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
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
			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-alb-resource",
					},
					Spec: &AwsAlbSpec{
						Subnets: []string{
							"subnet-abc1",
							"subnet-abc2",
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should not return a validation error when ssl.enabled is true and certificate_arn is set", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-ssl-arn-set",
					},
					Spec: &AwsAlbSpec{
						Subnets: []string{"subnet-1", "subnet-2"},
						Ssl: &AwsAlbSsl{
							Enabled: true,
							CertificateArn: &foreignkeyv1.StringValueOrRef{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
									Value: "arn:aws:acm:us-east-1:123456789012:certificate/test-cert",
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should not return a validation error when ssl.enabled is true and certificate_arn is set via reference", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-ssl-arn-ref-set",
					},
					Spec: &AwsAlbSpec{
						Subnets: []string{"subnet-1", "subnet-2"},
						Ssl: &AwsAlbSsl{
							Enabled: true,
							CertificateArn: &foreignkeyv1.StringValueOrRef{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_ValueFrom{
									ValueFrom: &foreignkeyv1.ValueFromRef{
										Env:       "dev",
										Name:      "some-other-resource",
										FieldPath: "spec.certificateArn",
									},
								},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("When SSL is enabled but certificate ARN is missing", func() {
		Context("aws_alb", func() {
			It("should return a validation error", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-ssl-no-arn",
					},
					Spec: &AwsAlbSpec{
						Subnets: []string{"subnet-1", "subnet-2"},
						Ssl: &AwsAlbSsl{
							Enabled: true,
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("certificate_arn must be set if ssl.enabled is true"))
			})
		})
	})
})
