package module

import (
	awsstaticwebsitev1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/aws/awsstaticwebsite/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsStaticWebsite *awsstaticwebsitev1.AwsStaticWebsite
	Spec             *awsstaticwebsitev1.AwsStaticWebsiteSpec
	Labels           map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsstaticwebsitev1.AwsStaticWebsiteStackInput) *Locals {
	locals := &Locals{}

	locals.AwsStaticWebsite = stackInput.Target
	if locals.AwsStaticWebsite != nil {
		locals.Spec = locals.AwsStaticWebsite.Spec
	}

	locals.Labels = map[string]string{}
	if locals.AwsStaticWebsite != nil {
		locals.Labels = map[string]string{
			"planton.org/resource":      "true",
			"planton.org/organization":  locals.AwsStaticWebsite.Metadata.Org,
			"planton.org/environment":   locals.AwsStaticWebsite.Metadata.Env,
			"planton.org/resource-kind": "AwsStaticWebsite",
			"planton.org/resource-id":   locals.AwsStaticWebsite.Metadata.Id,
		}
	}

	return locals
}
