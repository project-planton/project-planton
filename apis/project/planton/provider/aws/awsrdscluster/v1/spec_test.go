package awsrdsclusterv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    shared "github.com/project-planton/project-planton/apis/project/planton/shared"
    fk "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
)

func TestAwsRdsClusterSpec(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsRdsClusterSpec Validation Suite")
}

var _ = Describe("AwsRdsClusterSpec validations", func() {
    var res *AwsRdsCluster

    BeforeEach(func() {
        res = &AwsRdsCluster{
            ApiVersion: "aws.project-planton.org/v1",
            Kind:       "AwsRdsCluster",
            Metadata: &shared.ApiResourceMetadata{
                Name: "valid-name",
            },
            Spec: &AwsRdsClusterSpec{
                // satisfy either subnet_ids (>=2) or db_subnet_group_name
                SubnetIds: []*fk.StringValueOrRef{
                    {LiteralOrRef: &fk.StringValueOrRef_Value{Value: "subnet-aaaa"}},
                    {LiteralOrRef: &fk.StringValueOrRef_Value{Value: "subnet-bbbb"}},
                },
                Engine:        "aurora-mysql",
                EngineVersion: "8.0.mysql_aurora.3.05.2",
                ManageMasterUserPassword: true,
                // satisfy optional patterns and CEL
                PreferredMaintenanceWindow: "mon:00:00-tue:01:00",
                PreferredBackupWindow:      "00:00-01:00",
                SkipFinalSnapshot:          true,
            },
        }
    })

    It("accepts a minimal valid spec", func() {
        err := protovalidate.Validate(res)
        Expect(err).To(BeNil())
    })

    It("fails when engine_mode is not allowed", func() {
        res.Spec.EngineMode = "invalid-mode"
        err := protovalidate.Validate(res)
        Expect(err).NotTo(BeNil())
    })

    It("fails when logs exports do not match engine family (CEL)", func() {
        res.Spec.EnabledCloudwatchLogsExports = []string{"postgresql"}
        err := protovalidate.Validate(res)
        Expect(err).NotTo(BeNil())
    })

    It("fails when password is set while manage_master_user_password is true (CEL)", func() {
        res.Spec.Password = "secret"
        err := protovalidate.Validate(res)
        Expect(err).NotTo(BeNil())
    })

    It("fails when fewer than two subnet_ids are provided and no subnet group name (CEL)", func() {
        res.Spec.SubnetIds = []*fk.StringValueOrRef{
            {LiteralOrRef: &fk.StringValueOrRef_Value{Value: "subnet-only"}},
        }
        res.Spec.DbSubnetGroupName = nil
        err := protovalidate.Validate(res)
        Expect(err).NotTo(BeNil())
    })
})


