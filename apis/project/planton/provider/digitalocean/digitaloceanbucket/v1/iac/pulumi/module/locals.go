package module

import (
	"strconv"

	digitaloceanbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanbucket/v1"
	digitaloceanprovider "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanProviderConfig *digitaloceanprovider.DigitalOceanProviderConfig
	DigitalOceanBucket         *digitaloceanbucketv1.DigitalOceanBucket
	DigitalOceanLabels         map[string]string
}

// initializeLocals copies stack‑input fields into Locals and builds a label map.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceanbucketv1.DigitalOceanBucketStackInput) *Locals {
	var locals Locals

	locals.DigitalOceanBucket = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DigitalOceanLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanBucket.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanBucket.String(),
	}

	if locals.DigitalOceanBucket.Metadata.Org != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanBucket.Metadata.Org
	}
	if locals.DigitalOceanBucket.Metadata.Env != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanBucket.Metadata.Env
	}
	if locals.DigitalOceanBucket.Metadata.Id != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanBucket.Metadata.Id
	}

	locals.DigitalOceanProviderConfig = stackInput.ProviderConfig
	return &locals
}
