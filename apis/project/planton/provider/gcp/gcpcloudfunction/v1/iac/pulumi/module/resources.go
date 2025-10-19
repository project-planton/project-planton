package module

import (
	"github.com/pkg/errors"
	gcpcloudfunctionv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudfunction/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpcloudfunctionv1.GcpCloudFunctionStackInput) error {
	//create gcp provider using the credentials from the input
	_, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}
	return nil
}
