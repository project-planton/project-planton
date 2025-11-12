package module

import (
	"github.com/pkg/errors"
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *clickhousekubernetesv1.ClickHouseKubernetesStackInput) error {
	// Initialize local variables and exports
	locals := initializeLocals(ctx, stackInput)

	// Create Kubernetes provider from stack-input credentials
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	// Create dedicated namespace for ClickHouse resources
	createdNamespace, err := createNamespace(ctx, locals, kubernetesProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// Create password secret for ClickHouse authentication
	createdSecret, err := createPasswordSecret(ctx, locals, createdNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to create password secret")
	}

	// Create ClickHouseKeeperInstallation if auto-managed Keeper is requested
	if shouldCreateClickHouseKeeper(locals.ClickHouseKubernetes.Spec) {
		keeperConfig := getKeeperConfig(locals.ClickHouseKubernetes.Spec)
		if err := clickhouseKeeperInstallation(ctx, locals, createdNamespace, keeperConfig); err != nil {
			return errors.Wrap(err, "failed to create ClickHouseKeeperInstallation")
		}
	}

	// Create ClickHouseInstallation CRD using Altinity operator
	if err := clickhouseInstallation(ctx, locals, createdNamespace, createdSecret); err != nil {
		return errors.Wrap(err, "failed to create ClickHouseInstallation")
	}

	// Create ingress LoadBalancer service if enabled
	if err := createIngressLoadBalancer(ctx, locals, createdNamespace, kubernetesProvider); err != nil {
		return errors.Wrap(err, "failed to create ingress load balancer")
	}

	return nil
}
