package module

import (
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecretsmanager/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	AwsSecretsManager *awssecretsmanagerv1.AwsSecretsManager
	AwsTags           map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awssecretsmanagerv1.AwsSecretsManagerStackInput) *Locals {
	locals := &Locals{}
	locals.AwsSecretsManager = stackInput.Target

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsSecretsManager.Metadata.Org,
		awstagkeys.Environment:  locals.AwsSecretsManager.Metadata.Env,
		awstagkeys.ResourceKind: "aws-secrets-manager",
		awstagkeys.ResourceId:   locals.AwsSecretsManager.Metadata.Id,
	}
	return locals
}
