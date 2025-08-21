package awsrdsinstancev1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAwsRdsInstanceSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsRdsInstanceSpec Validation Suite")
}

var _ = Describe("AwsRdsInstanceSpec validations", func() {
	var spec *AwsRdsInstanceSpec

	BeforeEach(func() {
		spec = &AwsRdsInstanceSpec{
			SubnetIds: []*foreignkeyv1.StringValueOrRef{
				{LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-aaaaaaaa"}},
				{LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-bbbbbbbb"}},
			},
			Engine:             "postgres",
			EngineVersion:      "14.10",
			InstanceClass:      "db.t3.micro",
			AllocatedStorageGb: 20,
			Username:           "admin",
			Password:           "secret",
			Port:               5432,
			PubliclyAccessible: false,
			MultiAz:            false,
		}
	})

	It("accepts a valid spec with subnets", func() {
		Expect(protovalidate.Validate(spec)).To(BeNil())
	})

	It("accepts a valid spec with db_subnet_group_name instead of subnets", func() {
		spec.SubnetIds = nil
		spec.DbSubnetGroupName = &foreignkeyv1.StringValueOrRef{
			LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "my-db-subnet-group"},
		}
		Expect(protovalidate.Validate(spec)).To(BeNil())
	})

	Context("field validations", func() {
		It("fails when engine is empty", func() {
			spec.Engine = ""
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when engine_version is empty", func() {
			spec.EngineVersion = ""
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when instance_class does not start with 'db.'", func() {
			spec.InstanceClass = "t3.micro"
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when allocated_storage_gb <= 0", func() {
			spec.AllocatedStorageGb = 0
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when username is empty", func() {
			spec.Username = ""
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when password is empty", func() {
			spec.Password = ""
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when port is negative", func() {
			spec.Port = -1
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when port is greater than 65535", func() {
			spec.Port = 70000
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
	})

	Context("CEL validations: subnets_or_group", func() {
		It("fails when only one subnet is provided and no group name", func() {
			spec.SubnetIds = []*foreignkeyv1.StringValueOrRef{
				{LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-aaaa1111"}},
			}
			spec.DbSubnetGroupName = nil
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
		It("fails when no subnets and no group name are provided", func() {
			spec.SubnetIds = nil
			spec.DbSubnetGroupName = nil
			Expect(protovalidate.Validate(spec)).NotTo(BeNil())
		})
	})
})
