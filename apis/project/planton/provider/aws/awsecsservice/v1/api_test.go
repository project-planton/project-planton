package awsecsservicev1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAwsEcsService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsEcsService Suite")
}

var _ = Describe("AwsEcsService Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_ecs_service", func() {
			var validInput *AwsEcsService

			BeforeEach(func() {
				validInput = &AwsEcsService{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEcsService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-service",
					},
					Spec: &AwsEcsServiceSpec{
						ClusterArn: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
								Value: "arn:aws:ecs:us-east-1:123456789012:cluster/my-cluster"},
						},
						Container: &AwsEcsServiceContainer{
							Image: &AwsEcsServiceContainerImage{
								Repo: "example-repo",
								Tag:  "latest",
							},
							Cpu:    512,
							Memory: 1024,
						},
						Network: &AwsEcsServiceNetwork{
							Subnets: []*foreignkeyv1.StringValueOrRef{
								{
									LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-abc123"},
								},
							},
						},
						Alb: &AwsEcsServiceAlb{
							Enabled: true,
							Arn: &foreignkeyv1.StringValueOrRef{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
									Value: "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/app/myAlb"},
							},
							RoutingType:      "path",
							ListenerPort:     80,
							ListenerPriority: 100,
						},
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(validInput)
				Expect(err).To(BeNil())
			})
		})
	})

	Context("api_version validations", func() {
		var input *AwsEcsService

		BeforeEach(func() {
			input = &AwsEcsService{
				ApiVersion: "aws.project-planton.org/v1",
				Kind:       "AwsEcsService",
				Metadata: &shared.ApiResourceMetadata{
					Name: "test-service",
				},
				Spec: &AwsEcsServiceSpec{
					ClusterArn: &foreignkeyv1.StringValueOrRef{
						LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
							Value: "arn:aws:ecs:us-east-1:123456789012:cluster/test-cluster"},
					},
					Container: &AwsEcsServiceContainer{
						Image: &AwsEcsServiceContainerImage{
							Repo: "example-repo",
							Tag:  "latest",
						},
						Cpu:    256,
						Memory: 512,
					},
					Network: &AwsEcsServiceNetwork{
						Subnets: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-abc123"},
							},
						},
					},
				},
			}
		})

		It("should fail if api_version does not match 'aws.project-planton.org/v1'", func() {
			input.ApiVersion = "invalid-version"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("kind validations", func() {
		var input *AwsEcsService

		BeforeEach(func() {
			input = &AwsEcsService{
				ApiVersion: "aws.project-planton.org/v1",
				Kind:       "AwsEcsService",
				Metadata: &shared.ApiResourceMetadata{
					Name: "test-service",
				},
				Spec: &AwsEcsServiceSpec{
					ClusterArn: &foreignkeyv1.StringValueOrRef{
						LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
							Value: "arn:aws:ecs:us-east-1:123456789012:cluster/test-cluster"},
					},
					Container: &AwsEcsServiceContainer{
						Image: &AwsEcsServiceContainerImage{
							Repo: "example-repo",
							Tag:  "latest",
						},
						Cpu:    256,
						Memory: 512,
					},
					Network: &AwsEcsServiceNetwork{
						Subnets: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-abc123"},
							},
						},
					},
				},
			}
		})

		It("should fail if kind does not match 'AwsEcsService'", func() {
			input.Kind = "SomeOtherKind"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("routing_type validations", func() {
		var input *AwsEcsService

		BeforeEach(func() {
			input = &AwsEcsService{
				ApiVersion: "aws.project-planton.org/v1",
				Kind:       "AwsEcsService",
				Metadata: &shared.ApiResourceMetadata{
					Name: "test-service",
				},
				Spec: &AwsEcsServiceSpec{
					ClusterArn: &foreignkeyv1.StringValueOrRef{
						LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
							Value: "arn:aws:ecs:us-east-1:123456789012:cluster/test-cluster"},
					},
					Container: &AwsEcsServiceContainer{
						Image: &AwsEcsServiceContainerImage{
							Repo: "example-repo",
							Tag:  "latest",
						},
						Cpu:    256,
						Memory: 512,
					},
					Network: &AwsEcsServiceNetwork{
						Subnets: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-abc123"},
							},
						},
					},
					Alb: &AwsEcsServiceAlb{
						Enabled: true,
						Arn: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
								Value: "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/app/myAlb"},
						},
						RoutingType:      "path",
						ListenerPort:     80,
						ListenerPriority: 100,
					},
				},
			}
		})

		It("should allow 'path'", func() {
			input.Spec.Alb.RoutingType = "path"
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should allow 'hostname'", func() {
			input.Spec.Alb.RoutingType = "hostname"
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})

		It("should fail for any other value", func() {
			input.Spec.Alb.RoutingType = "invalid"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})
})
