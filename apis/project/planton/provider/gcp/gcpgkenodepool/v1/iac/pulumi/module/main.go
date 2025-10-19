package module

import (
	"github.com/pkg/errors"
	gcpgkenodepoolv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkenodepool/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi entry‑point invoked by the runtime.
func Resources(ctx *pulumi.Context, stackInput *gcpgkenodepoolv1.GcpGkeNodePoolStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	// Set up the Google provider from the supplied GCP credential.
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to configure google provider")
	}

	// Discover the parent GKE cluster so we can fetch its region/zone & project ID.
	clusterInfo, err := container.LookupCluster(ctx, &container.LookupClusterArgs{
		Name:    locals.ClusterName,
		Project: pulumi.StringRef(locals.GcpGkeNodePool.Spec.ClusterProjectId.GetValue()),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to lookup parent cluster %q", locals.ClusterName)
	}

	// Create the node pool.
	if err := nodePool(ctx, locals, clusterInfo, gcpProvider); err != nil {
		return errors.Wrap(err, "failed to create node pool")
	}

	return nil
}
