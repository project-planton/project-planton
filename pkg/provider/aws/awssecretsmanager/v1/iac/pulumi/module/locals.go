package module

import (
	"strconv"

	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/pkg/provider/aws/awssecretsmanager/v1"
	"github.com/project-planton/project-planton/pkg/shared/cloudresourcekind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsSecretsManager.String(),
		awstagkeys.ResourceId:   locals.AwsSecretsManager.Metadata.Id,
	}
	return locals
}
