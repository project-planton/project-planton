package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceanvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanvpc/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanVpc            *digitaloceanvpcv1.DigitalOceanVpc
	DoLabels                   map[string]string
}

// initializeLocals copies stack‑input fields into the Locals struct and builds
// a reusable label map. Mirrors the style of gcp_vpc's initializeLocals().
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceanvpcv1.DigitalOceanVpcStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanVpc = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanVpc.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanVpc.String(),
	}

	if locals.DigitalOceanVpc.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanVpc.Metadata.Org
	}

	if locals.DigitalOceanVpc.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanVpc.Metadata.Env
	}

	if locals.DigitalOceanVpc.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanVpc.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
