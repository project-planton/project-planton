package module

import (
	"github.com/pkg/errors"
	gcpgkeclustercorev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgkeclustercore/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi program entry‑point invoked by the ProjectPlanton
func Resources(
	ctx *pulumi.Context,
	stackInput *gcpgkeclustercorev1.GcpGkeClusterCoreStackInput,
) error {
	// gather locals (Terraform‑style “locals”)
	locals := initializeLocals(stackInput)

	// configure a GCP provider from the given credential
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup google provider")
	}

	// Cluster.
	_, err = cluster(ctx, locals, gcpProvider)
	if err != nil {
		return errors.Wrap(err, "cluster creation failed")
	}

	return nil
}
