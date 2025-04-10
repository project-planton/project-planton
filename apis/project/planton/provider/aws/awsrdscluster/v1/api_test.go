package awsrdsclusterv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsRdsCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsRdsCluster Suite")
}

var _ = Describe("AwsRdsCluster Custom Validation Tests", func() {

	var input *AwsRdsCluster

	BeforeEach(func() {
		input = &AwsRdsCluster{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsRdsCluster",
			Metadata: &shared.ApiResourceMetadata{
				Name: "valid-rds-cluster",
			},
			Spec: &AwsRdsClusterSpec{
				Engine:                        "aurora-mysql",
				EngineVersion:                 "5.7.mysql_aurora.2.03.2",
				ClusterFamily:                 "some-family",
				InstanceType:                  "db.t4g.medium",
				EnhancedMonitoringRoleEnabled: false,
				RdsMonitoringInterval:         0,
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("aws_rds", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			Context("engine enumerations", func() {
				It("should fail validation for unsupported engine value", func() {
					input.Spec.Engine = "invalid-engine"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})

				It("should pass validation for a supported engine value", func() {
					input.Spec.Engine = "postgres"
					err := protovalidate.Validate(input)
					Expect(err).To(BeNil())
				})
			})

			Context("enhanced monitoring interval rule", func() {
				It("should fail when enhanced monitoring is enabled but interval is 0", func() {
					input.Spec.EnhancedMonitoringRoleEnabled = true
					input.Spec.RdsMonitoringInterval = 0
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})

				It("should pass when enhanced monitoring is enabled and interval is non-zero", func() {
					input.Spec.EnhancedMonitoringRoleEnabled = true
					input.Spec.RdsMonitoringInterval = 5
					err := protovalidate.Validate(input)
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
