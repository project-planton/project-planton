package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/vars"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// workloadDeployer creates a service account for deploying workloads to the GKE cluster and assigns it necessary roles.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - createdCluster: The GKE cluster to which workloads will be deployed.
//
// Returns:
// - *serviceaccount.Key: A pointer to the created service account key.
// - error: An error object if there is any issue during the service account or key creation.
//
// The function performs the following steps:
// 1. Creates a service account with a description and display name for deploying workloads.
// 2. Exports the email of the created service account.
// 3. Creates a key for the service account and exports the private key.
// 4. Creates IAM bindings to grant the service account the roles of container admin and cluster admin.
// 5. Handles errors and returns the created service account key and any errors encountered.
func workloadDeployer(ctx *pulumi.Context, createdCluster *container.Cluster) (*serviceaccount.Key, error) {
	//create workload deployer service account
	createdWorkloadDeployerServiceAccount, err := serviceaccount.NewAccount(ctx,
		vars.WorkloadDeployServiceAccountName,
		&serviceaccount.AccountArgs{
			Project:     createdCluster.Project,
			Description: pulumi.String("service account to deploy workloads"),
			AccountId:   pulumi.String(vars.WorkloadDeployServiceAccountName),
			DisplayName: pulumi.String(vars.WorkloadDeployServiceAccountName),
		}, pulumi.Parent(createdCluster))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create workload deployer service account")
	}

	//export email of the created workload deployer service account
	ctx.Export(outputs.WorkloadDeployerGsaEmail, createdWorkloadDeployerServiceAccount.Email)

	//create key for workload deployer service account.
	createdWorkloadDeployerServiceAccountKey, err := serviceaccount.NewKey(ctx,
		vars.WorkloadDeployServiceAccountName,
		&serviceaccount.KeyArgs{
			ServiceAccountId: createdWorkloadDeployerServiceAccount.Name,
		}, pulumi.Parent(createdWorkloadDeployerServiceAccount))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create key for workload-deployer service account")
	}

	//export workload deployer google service account key
	ctx.Export(outputs.WorkloadDeployerGsaKeyBase64, createdWorkloadDeployerServiceAccountKey.PrivateKey)

	// create iam-binding for workload-deployer to manage the container cluster itself
	_, err = projects.NewIAMBinding(ctx,
		fmt.Sprintf("%s-container-admin", vars.WorkloadDeployServiceAccountName),
		&projects.IAMBindingArgs{
			Members: pulumi.StringArray{pulumi.Sprintf("serviceAccount:%s", createdWorkloadDeployerServiceAccount.Email)},
			Project: createdCluster.Project,
			Role:    pulumi.String("roles/container.admin"),
		}, pulumi.Parent(createdWorkloadDeployerServiceAccount))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create container-admin iam binding for workload deployer")
	}

	// create iam-binding for workload-deployer to manage resources inside container clusters
	_, err = projects.NewIAMBinding(ctx,
		fmt.Sprintf("%s-kube-cluster-admin", vars.WorkloadDeployServiceAccountName),
		&projects.IAMBindingArgs{
			Members: pulumi.StringArray{pulumi.Sprintf("serviceAccount:%s", createdWorkloadDeployerServiceAccount.Email)},
			Project: createdCluster.Project,
			Role:    pulumi.String("roles/container.clusterAdmin"),
		}, pulumi.Parent(createdWorkloadDeployerServiceAccount))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create cluster-admin iam binding for workload deployer")
	}

	return createdWorkloadDeployerServiceAccountKey, nil
}
