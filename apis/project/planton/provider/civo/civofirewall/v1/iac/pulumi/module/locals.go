package module

import (
	civofirewallv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civofirewall/v1"
	civoprovider "github.com/project-planton/project-planton/apis/project/planton/provider/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles references we need across files.
type Locals struct {
	CivoProviderConfig *civoprovider.CivoProviderConfig
	CivoFirewall       *civofirewallv1.CivoFirewall
	CivoLabels         map[string]string
}

// initializeLocals gives us quick access to spec + metadata.
func initializeLocals(_ *pulumi.Context, stackInput *civofirewallv1.CivoFirewallStackInput) *Locals {
	locals := &Locals{}
	locals.CivoFirewall = stackInput.Target
	locals.CivoProviderConfig = stackInput.ProviderConfig

	// Basic labels — extend as needed for tagging conventions.
	locals.CivoLabels = map[string]string{
		"resource":      "true",
		"resource_name": locals.CivoFirewall.Metadata.Name,
		"resource_kind": "CivoFirewall",
	}
	if locals.CivoFirewall.Metadata.Org != "" {
		locals.CivoLabels["organization"] = locals.CivoFirewall.Metadata.Org
	}
	if locals.CivoFirewall.Metadata.Env != "" {
		locals.CivoLabels["environment"] = locals.CivoFirewall.Metadata.Env
	}
	if locals.CivoFirewall.Metadata.Id != "" {
		locals.CivoLabels["resource_id"] = locals.CivoFirewall.Metadata.Id
	}

	return locals
}
