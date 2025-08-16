package module

import (
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsRdsInstance *awsrdsinstancev1.AwsRdsInstance
	Labels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsrdsinstancev1.AwsRdsInstanceStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsRdsInstance = stackInput.Target

	// initialize standardized labels
	locals.Labels = map[string]string{
		"planton.org/resource":      "true",
		"planton.org/organization":  locals.AwsRdsInstance.Metadata.Org,
		"planton.org/environment":   locals.AwsRdsInstance.Metadata.Env,
		"planton.org/resource-kind": "AwsRdsInstance",
		"planton.org/resource-id":   locals.AwsRdsInstance.Metadata.Id,
	}

	return locals
}
