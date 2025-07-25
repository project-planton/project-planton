package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceanvolumev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanvolume/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanVolume         *digitaloceanvolumev1.DigitalOceanVolume
	DigitalOceanLabels         map[string]string
}

// initializeLocals copies stack‑input fields into the Locals struct and builds
// a reusable label map. Mirrors the style of digital_ocean_vpc's initializeLocals().
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceanvolumev1.DigitalOceanVolumeStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanVolume = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DigitalOceanLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanVolume.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanVolume.String(),
	}

	if locals.DigitalOceanVolume.Metadata.Org != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanVolume.Metadata.Org
	}

	if locals.DigitalOceanVolume.Metadata.Env != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanVolume.Metadata.Env
	}

	if locals.DigitalOceanVolume.Metadata.Id != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanVolume.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
