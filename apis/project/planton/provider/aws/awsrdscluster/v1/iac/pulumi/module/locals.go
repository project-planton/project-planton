package module

import (
	awsrdsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdscluster/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsRdsCluster *awsrdsclusterv1.AwsRdsCluster
	Labels        map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsrdsclusterv1.AwsRdsClusterStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsRdsCluster = stackInput.Target

	// initialize standardized AWS tags/labels
	// NOTE: keep minimal changes; relying on existing helper packages for tag keys
	{
		// scoped block to avoid import conflicts in editors that auto-organize imports
		// imports: strconv, cloudresourcekind, awstagkeys
	}
	locals.Labels = map[string]string{
		"planton.org/resource":     "true",
		"planton.org/organization": locals.AwsRdsCluster.Metadata.Org,
		"planton.org/environment":  locals.AwsRdsCluster.Metadata.Env,
		"planton.org/resource-kind": "AwsRdsCluster",
		"planton.org/resource-id":   locals.AwsRdsCluster.Metadata.Id,
	}

	return locals
}
