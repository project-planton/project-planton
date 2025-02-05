package module

import (
	"github.com/pkg/errors"
	gkeclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources function is the pulumi program that deploys GKE cluster along with chosen addons.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
//
// Returns:
// - error: An error object if there is any issue during the resource creation.
//
// The function performs the following steps:
// 1. Initializes local variables and configuration from the input.
// 2. Sets up the GCP provider using the provided GCP credentials.
// 3. Creates a GCP folder for organizing the projects.
// 4. Creates the GKE cluster within the specified folder.
// 5. Creates the node pools for the GKE cluster.
// 6. Creates a service account and key for deploying workloads to the cluster.
func Resources(ctx *pulumi.Context, stackInput *gkeclusterv1.GkeClusterStackInput) error {
	locals := localz.Initialize(ctx, stackInput)

	//create gcp-provider using the gcp-credential from input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup google provider")
	}

	//create cluster
	createdCluster, err := cluster(ctx, locals, gcpProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create container cluster")
	}

	//create node-pools
	if err = clusterNodePools(ctx, locals, createdCluster); err != nil {
		return errors.Wrap(err, "failed to create cluster node-pools")
	}

	//create workload-deployer google service account resources
	if err := workloadDeployer(ctx, createdCluster); err != nil {
		return errors.Wrap(err, "failed to create workload-deployer resources")
	}

	return nil
}
