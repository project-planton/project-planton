package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceandatabaseclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceandatabasecluster/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec  *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanDatabaseCluster *digitaloceandatabaseclusterv1.DigitalOceanDatabaseCluster
	DoLabels                    map[string]string
}

// initializeLocals copies stack‑input fields into the Locals struct and builds
// a reusable label map. Mirrors the pattern of the DigitalOcean VPC module.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceandatabaseclusterv1.DigitalOceanDatabaseClusterStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanDatabaseCluster = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanDatabaseCluster.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanDatabaseCluster.String(),
	}

	if locals.DigitalOceanDatabaseCluster.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanDatabaseCluster.Metadata.Org
	}

	if locals.DigitalOceanDatabaseCluster.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanDatabaseCluster.Metadata.Env
	}

	if locals.DigitalOceanDatabaseCluster.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanDatabaseCluster.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
