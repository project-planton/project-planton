package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceandnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceandnszone/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals holds quick references used by other files.
type Locals struct {
	DigitalOceanCredentialSpec *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanDnsZone        *digitaloceandnszonev1.DigitalOceanDnsZone
	DoLabels                   map[string]string
}

// initializeLocals mirrors the pattern from the VPC module.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceandnszonev1.DigitalOceanDnsZoneStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanDnsZone = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanDnsZone.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanDnsZone.String(),
	}

	if locals.DigitalOceanDnsZone.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanDnsZone.Metadata.Org
	}
	if locals.DigitalOceanDnsZone.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanDnsZone.Metadata.Env
	}
	if locals.DigitalOceanDnsZone.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanDnsZone.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
