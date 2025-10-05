package pulumicloudflareprovider

import (
	"fmt"
	"reflect"

	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/pulumi/pulumioutput"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get builds a pulumi-cloudflare Provider using the supplied credential.
// It supports both API Token and Legacy API Key authentication schemes.
// If the credential is nil, Pulumi's provider will fall back to environment
// variables (CLOUDFLARE_API_TOKEN, CLOUDFLARE_API_KEY, etc.).
func Get(
	ctx *pulumi.Context,
	credentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec,
	nameSuffixes ...string,
) (*cloudflare.Provider, error) {

	providerArgs := &cloudflare.ProviderArgs{}

	// Map credential fields when present; leave them nil to defer to env-vars.
	if credentialSpec != nil {
		// Handle authentication based on scheme
		switch credentialSpec.AuthScheme {
		case cloudflarecredentialv1.CloudflareAuthScheme_api_token:
			if credentialSpec.ApiToken != "" {
				providerArgs.ApiToken = pulumi.StringPtr(credentialSpec.ApiToken)
			}
		case cloudflarecredentialv1.CloudflareAuthScheme_legacy_api_key:
			if credentialSpec.ApiKey != "" {
				providerArgs.ApiKey = pulumi.StringPtr(credentialSpec.ApiKey)
			}
			if credentialSpec.Email != "" {
				providerArgs.Email = pulumi.StringPtr(credentialSpec.Email)
			}
		}
	}

	provider, err := cloudflare.NewProvider(
		ctx,
		ProviderResourceName(nameSuffixes),
		providerArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare provider")
	}

	return provider, nil
}

// ProviderResourceName builds a deterministic Pulumi resource name such as
// "cloudflare-primary". Mirrors the google helper for naming consistency.
func ProviderResourceName(suffixes []string) string {
	name := "cloudflare"
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

// PulumiOutputName produces canonical output names (e.g. "cf_zone_id") to keep
// stack outputs predictable across modules.
func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	output := fmt.Sprintf("cf_%s", pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		output = fmt.Sprintf("%s_%s", output, s)
	}
	return output
}
