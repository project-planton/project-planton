package module

import (
	"strconv"
	"strings"

	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/gcplabelkeys"

	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	gcpcloudrunv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudrun/v1"
)

// Locals holds handy references and derived values used across this module.
type Locals struct {
	GcpCredentialSpec *gcpcredentialv1.GcpCredentialSpec
	GcpCloudRun       *gcpcloudrunv1.GcpCloudRun
	GcpLabels         map[string]string
}

// initializeLocals fills the Locals struct from the incoming stack input.
func initializeLocals(stackInput *gcpcloudrunv1.GcpCloudRunStackInput) *Locals {
	locals := &Locals{}

	locals.GcpCloudRun = stackInput.Target

	target := stackInput.Target

	locals.GcpCredentialSpec = stackInput.ProviderCredential

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceName: target.Metadata.Name,
		gcplabelkeys.ResourceKind: strings.ToLower(cloudresourcekind.CloudResourceKind_GcpCloudRun.String()),
	}

	if target.Metadata.Id != "" {
		locals.GcpLabels[gcplabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.GcpLabels[gcplabelkeys.Environment] = target.Metadata.Env
	}

	return locals
}
