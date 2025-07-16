package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceanappplatformservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanappplatformservice/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals gathers convenient references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec     *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanAppPlatformService *digitaloceanappplatformservicev1.DigitalOceanAppPlatformService
	DoLabels                       map[string]string
}

// initializeLocals mirrors the pattern used in digital_ocean_vpc's initializeLocals().
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceanappplatformservicev1.DigitalOceanAppPlatformServiceStackInput) *Locals {
	locals := &Locals{}

	locals.DigitalOceanAppPlatformService = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DoLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanAppPlatformService.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanAppPlatformService.String(),
	}

	if locals.DigitalOceanAppPlatformService.Metadata.Org != "" {
		locals.DoLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanAppPlatformService.Metadata.Org
	}

	if locals.DigitalOceanAppPlatformService.Metadata.Env != "" {
		locals.DoLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanAppPlatformService.Metadata.Env
	}

	if locals.DigitalOceanAppPlatformService.Metadata.Id != "" {
		locals.DoLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanAppPlatformService.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential

	return locals
}
