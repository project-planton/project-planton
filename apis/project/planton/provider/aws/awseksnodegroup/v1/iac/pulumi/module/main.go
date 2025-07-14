package module

import (
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awseksnodegroup/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for the aws_eks_node_group Pulumi module.
// It prepares context, configures the AWS provider, and calls certManagerCert().
func Resources(ctx *pulumi.Context, stackInput *awseksnodegroupv1.AwsEksNodeGroupStackInput) error {
	return nil
}
