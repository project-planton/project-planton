package awsecsservicev1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
)

func TestAwsEcsService(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsEcsService Suite")
}

var _ = ginkgo.Describe("AwsEcsService Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_ecs_service", func() {
			var validInput *AwsEcsService

			ginkgo.BeforeEach(func() {
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

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(validInput)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Context("api_version validations", func() {
		var input *AwsEcsService

		ginkgo.BeforeEach(func() {
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

		ginkgo.It("should fail if api_version does not match 'aws.project-planton.org/v1'", func() {
			input.ApiVersion = "invalid-version"
			err := protovalidate.Validate(input)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Context("kind validations", func() {
		var input *AwsEcsService

		ginkgo.BeforeEach(func() {
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

		ginkgo.It("should fail if kind does not match 'AwsEcsService'", func() {
			input.Kind = "SomeOtherKind"
			err := protovalidate.Validate(input)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Context("routing_type validations", func() {
		var input *AwsEcsService

		ginkgo.BeforeEach(func() {
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

		ginkgo.It("should allow 'path'", func() {
			input.Spec.Alb.RoutingType = "path"
			err := protovalidate.Validate(input)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("should allow 'hostname'", func() {
			input.Spec.Alb.RoutingType = "hostname"
			err := protovalidate.Validate(input)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("should fail for any other value", func() {
			input.Spec.Alb.RoutingType = "invalid"
			err := protovalidate.Validate(input)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})
})
