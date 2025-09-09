package awsrdsinstancev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
)

func TestAwsRdsInstanceSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsRdsInstanceSpec Custom Validation Tests")
}

var _ = Describe("AwsRdsInstanceSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_rds_instance", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsRdsInstance{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsRdsInstance",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-rds-instance",
					},
					Spec: &AwsRdsInstanceSpec{
						SubnetIds: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-12345678"},
							},
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-87654321"},
							},
						},
						Engine:             "postgres",
						EngineVersion:      "14.10",
						InstanceClass:      "db.t3.micro",
						AllocatedStorageGb: 20,
						Username:           "dbmaster",
						Password:           "mypassword123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
