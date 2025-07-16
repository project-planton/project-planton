package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceankubernetesnodepoolv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceankubernetesnodepool/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals aggregates handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec     *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanKubernetesNodePool *digitaloceankubernetesnodepoolv1.DigitalOceanKubernetesNodePool
	DoLabels                       map[string]string
}

// initializeLocals mirrors the pattern used in digital_ocean_vpc.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceankubernetesnodepoolv1.DigitalOceanKubernetesNodePoolStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanKubernetesNodePool = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanKubernetesNodePool.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanKubernetesNodePool.String(),
	}

	if locals.DigitalOceanKubernetesNodePool.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanKubernetesNodePool.Metadata.Org
	}

	if locals.DigitalOceanKubernetesNodePool.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanKubernetesNodePool.Metadata.Env
	}

	if locals.DigitalOceanKubernetesNodePool.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanKubernetesNodePool.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
