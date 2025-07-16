package pulumidigitaloceanprovider

import (
	"fmt"
	"reflect"

	digitaloceancredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/digitaloceancredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/pulumi/pulumioutput"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get builds a pulumi‑digitalocean Provider using the supplied credential.
// If the credential is nil or any individual field is blank, Pulumi’s provider
// will fall back to environment variables (DIGITALOCEAN_TOKEN, etc.),
// matching Terraform’s behavior.
func Get(
	ctx *pulumi.Context,
	doCredentialSpec *digitaloceancredentialv1.DigitalOceanCredentialSpec,
	nameSuffixes ...string,
) (*digitalocean.Provider, error) {

	providerArgs := &digitalocean.ProviderArgs{}

	// Map credential fields when present; leave them nil to defer to env‑vars.
	if doCredentialSpec != nil {
		if doCredentialSpec.ApiToken != "" {
			providerArgs.Token = pulumi.StringPtr(doCredentialSpec.ApiToken)
		}
		if doCredentialSpec.SpacesAccessId != "" {
			providerArgs.SpacesAccessId = pulumi.StringPtr(doCredentialSpec.SpacesAccessId)
		}
		if doCredentialSpec.SpacesSecretKey != "" {
			providerArgs.SpacesSecretKey = pulumi.StringPtr(doCredentialSpec.SpacesSecretKey)
		}
	}

	provider, err := digitalocean.NewProvider(
		ctx,
		ProviderResourceName(nameSuffixes),
		providerArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean provider")
	}

	return provider, nil
}

// ProviderResourceName builds a deterministic Pulumi resource name such as
// "digitalocean‑primary". Mirrors the google helper for naming consistency.
func ProviderResourceName(suffixes []string) string {
	name := "digitalocean"
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

// PulumiOutputName produces canonical output names (e.g. "do_vpc_id") to keep
// stack outputs predictable across modules.
func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	output := fmt.Sprintf("do_%s", pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		output = fmt.Sprintf("%s_%s", output, s)
	}
	return output
}
