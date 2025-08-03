package pulumicivoprovider

import (
	"fmt"
	"reflect"

	civocredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/civocredential/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/pulumi/pulumioutput"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get builds a pulumi‑civo Provider using the supplied credential.
// Like DigitalOcean, leaving individual fields nil lets Pulumi fall
// back to environment variables (CIVO_TOKEN, CIVO_REGION, etc.).
func Get(
	ctx *pulumi.Context,
	civoCredentialSpec *civocredentialv1.CivoCredentialSpec,
	nameSuffixes ...string,
) (*civo.Provider, error) {

	providerArgs := &civo.ProviderArgs{}

	// Map credential fields when they are present; otherwise rely on env‑vars.
	if civoCredentialSpec != nil {
		if civoCredentialSpec.ApiToken != "" {
			providerArgs.Token = pulumi.StringPtr(civoCredentialSpec.ApiToken)
		}
		if civoCredentialSpec.DefaultRegion != 0 {
			providerArgs.Region = pulumi.StringPtr(civoCredentialSpec.DefaultRegion.String())
		}
	}

	provider, err := civo.NewProvider(
		ctx,
		ProviderResourceName(nameSuffixes),
		providerArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create civo provider")
	}

	return provider, nil
}

// ProviderResourceName builds deterministic Pulumi resource names such as
// "civo‑primary". Mirrors helper functions in other providers for consistency.
func ProviderResourceName(suffixes []string) string {
	name := "civo"
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

// PulumiOutputName yields canonical output names (e.g. "civo_vpc_id")
// so stack outputs stay predictable across modules.
func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	output := fmt.Sprintf("civo_%s", pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		output = fmt.Sprintf("%s_%s", output, s)
	}
	return output
}
