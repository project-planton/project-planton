package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflareloadbalancerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareloadbalancer/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals stores quick references for metadata, spec & credentials.
type Locals struct {
	CloudflareCredentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareLoadBalancer   *cloudflareloadbalancerv1.CloudflareLoadBalancer
}

// initializeLocals copies relevant stack‑input fields into Locals.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflareloadbalancerv1.CloudflareLoadBalancerStackInput) *Locals {
	return &Locals{
		CloudflareCredentialSpec: stackInput.ProviderCredential,
		CloudflareLoadBalancer:   stackInput.Target,
	}
}
