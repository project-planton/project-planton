package pulumigoogleprovider

import (
	"encoding/base64"
	"fmt"
	"reflect"

	gcpprovider "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/pulumi/pulumioutput"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Get(ctx *pulumi.Context, gcpProviderConfig *gcpprovider.GcpProviderConfig,
	nameSuffixes ...string) (*gcp.Provider, error) {
	gcpProviderArgs := &gcp.ProviderArgs{}

	//if stack-input contains gcp-credentials, provider will be created with those credentials
	if gcpProviderConfig != nil {
		serviceAccountKey, err := base64.StdEncoding.DecodeString(gcpProviderConfig.ServiceAccountKeyBase64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode base64 encoded"+
				" google service account credential")
		}
		gcpProviderArgs = &gcp.ProviderArgs{Credentials: pulumi.String(serviceAccountKey)}
	}

	googleProvider, err := gcp.NewProvider(ctx, ProviderResourceName(nameSuffixes), gcpProviderArgs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create google provider")
	}

	return googleProvider, nil
}

func ProviderResourceName(suffixes []string) string {
	name := "google"
	for _, s := range suffixes {
		name = fmt.Sprintf("%s-%s", name, s)
	}
	return name
}

func PulumiOutputName(r interface{}, name string, suffixes ...string) string {
	outputName := fmt.Sprintf("gcp_%s", pulumioutput.Name(reflect.TypeOf(r), name))
	for _, s := range suffixes {
		outputName = fmt.Sprintf("%s_%s", outputName, s)
	}
	return outputName
}
