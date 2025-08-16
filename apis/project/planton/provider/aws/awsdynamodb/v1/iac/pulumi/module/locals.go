package module

import (
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsDynamodb *awsdynamodbv1.AwsDynamodb
	Labels      map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) *Locals {
	locals := &Locals{}

	locals.AwsDynamodb = stackInput.Target

	locals.Labels = map[string]string{
		"planton.org/resource":      "true",
		"planton.org/organization":  locals.AwsDynamodb.Metadata.Org,
		"planton.org/environment":   locals.AwsDynamodb.Metadata.Env,
		"planton.org/resource-kind": "AwsDynamodb",
		"planton.org/resource-id":   locals.AwsDynamodb.Metadata.Id,
	}

	return locals
}
