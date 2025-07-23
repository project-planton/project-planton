package module

import (
	"strconv"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	digitaloceancertificatev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceancertificate/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/digitaloceanlabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	DigitalOceanCredentialSpec *digitaloceancredentialv1.DigitalOceanCredentialSpec
	DigitalOceanCertificate         *digitaloceancertificatev1.DigitalOceanCertificate
	DigitalOceanLabels         map[string]string
}

// initializeLocals copies stack‑input fields into Locals and builds a label map.
func initializeLocals(_ *pulumi.Context, stackInput *digitaloceancertificatev1.DigitalOceanCertificateStackInput) *Locals {
	var locals Locals

	locals.DigitalOceanCertificate = stackInput.Target

	// Standard Planton labels for DigitalOcean resources.
	locals.DigitalOceanLabels = map[string]string{
		digitaloceanlabelkeys.Resource:     strconv.FormatBool(true),
		digitaloceanlabelkeys.ResourceName: locals.DigitalOceanCertificate.Metadata.Name,
		digitaloceanlabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_DigitalOceanCertificate.String(),
	}

	if locals.DigitalOceanCertificate.Metadata.Org != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Organization] = locals.DigitalOceanCertificate.Metadata.Org
	}
	if locals.DigitalOceanCertificate.Metadata.Env != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.Environment] = locals.DigitalOceanCertificate.Metadata.Env
	}
	if locals.DigitalOceanCertificate.Metadata.Id != "" {
		locals.DigitalOceanLabels[digitaloceanlabelkeys.ResourceId] = locals.DigitalOceanCertificate.Metadata.Id
	}

	locals.DigitalOceanCredentialSpec = stackInput.ProviderCredential
	return &locals
}
