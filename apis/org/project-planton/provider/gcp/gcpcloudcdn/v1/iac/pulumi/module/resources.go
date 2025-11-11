package module

import (
	"github.com/pkg/errors"
	gcpcloudcdnv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/gcp/gcpcloudcdn/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	GcpLabels map[string]string
}

func Resources(ctx *pulumi.Context, stackInput *gcpcloudcdnv1.GcpCloudCdnStackInput) error {
	//create gcp provider using the credentials from the input
	_, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}
	return nil
}
