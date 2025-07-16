package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceankubernetesclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceankubernetescluster/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec    *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanKubernetesCluster *digitaloceankubernetesclusterv1.DigitalOceanKubernetesCluster
	DoLabels                      map[string]string
}

// initializeLocals copies stack‑input fields into the Locals struct and builds
// a reusable label map—mirrors the style used in digital_ocean_vpc.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceankubernetesclusterv1.DigitalOceanKubernetesClusterStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanKubernetesCluster = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanKubernetesCluster.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanKubernetesCluster.String(),
	}

	if locals.DigitalOceanKubernetesCluster.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanKubernetesCluster.Metadata.Org
	}

	if locals.DigitalOceanKubernetesCluster.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanKubernetesCluster.Metadata.Env
	}

	if locals.DigitalOceanKubernetesCluster.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanKubernetesCluster.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
