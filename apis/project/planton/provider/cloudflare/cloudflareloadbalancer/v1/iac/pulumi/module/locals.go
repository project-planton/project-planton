package module

import (
	cloudflareloadbalancerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareloadbalancer/v1"
	cloudflareprovider "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals stores quick references for metadata, spec & credentials.
type Locals struct {
	CloudflareProviderConfig *cloudflareprovider.CloudflareProviderConfig
	CloudflareLoadBalancer   *cloudflareloadbalancerv1.CloudflareLoadBalancer
}

// initializeLocals copies relevant stack‑input fields into Locals.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflareloadbalancerv1.CloudflareLoadBalancerStackInput) *Locals {
	return &Locals{
		CloudflareProviderConfig: stackInput.ProviderConfig,
		CloudflareLoadBalancer:   stackInput.Target,
	}
}
