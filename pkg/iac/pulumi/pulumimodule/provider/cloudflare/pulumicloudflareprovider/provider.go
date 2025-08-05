package pulumicloudflareprovider

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Get(
	ctx *pulumi.Context,
	credentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec,
	nameSuffixes ...string,
) (*cloudflare.Provider, error) {

	return nil, nil
}
