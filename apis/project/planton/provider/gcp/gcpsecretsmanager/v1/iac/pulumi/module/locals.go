package module

import (
	gcpsecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsecretsmanager/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpSecretsManager *gcpsecretsmanagerv1.GcpSecretsManager
	GcpLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpsecretsmanagerv1.GcpSecretsManagerStackInput) *Locals {
	locals := &Locals{}

	locals.GcpSecretsManager = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceKind: "gcp_secrets_manager",
		gcplabelkeys.ResourceId:   locals.GcpSecretsManager.Metadata.Id,
	}

	if locals.GcpSecretsManager.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpSecretsManager.Metadata.Org
	}

	if locals.GcpSecretsManager.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpSecretsManager.Metadata.Env
	}

	return locals
}
