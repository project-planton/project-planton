package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceancontainerregistryv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceancontainerregistry/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec    *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanContainerRegistry *digitaloceancontainerregistryv1.DigitalOceanContainerRegistry
	DoLabels                      map[string]string
}

// initializeLocals copies stack‑input fields into the Locals struct and builds
// a reusable label map.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceancontainerregistryv1.DigitalOceanContainerRegistryStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanContainerRegistry = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanContainerRegistry.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanContainerRegistry.String(),
	}

	if locals.DigitalOceanContainerRegistry.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanContainerRegistry.Metadata.Org
	}

	if locals.DigitalOceanContainerRegistry.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanContainerRegistry.Metadata.Env
	}

	if locals.DigitalOceanContainerRegistry.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanContainerRegistry.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
