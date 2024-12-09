package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/addons"
	"github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1/iac/pulumi/module/localz"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// clusterAddons installs the specified Kubernetes addons for the created GKE cluster.
// It checks the configuration and installs addons such as Ingress Nginx, Istio, Cert Manager, External Secrets,
// External DNS, Postgres Operator, Solr Operator, and Kafka Operator.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdCluster: The GKE cluster where the addons will be installed.
// - gcpProvider: The GCP provider for Pulumi.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation of the addons.
//
// The function performs the following steps:
// 1. Checks if Ingress Nginx addon is to be installed and installs it if required.
// 2. Checks if Istio addon is to be installed and installs it if required.
// 3. Checks if Cert Manager addon is to be installed and installs it if required.
// 4. Checks if External Secrets addon is to be installed and installs it if required.
// 5. Checks if External DNS addon is to be installed and installs it if required.
// 6. Checks if Postgres Operator addon is to be installed and installs it if required.
// 7. Checks if Solr Operator addon is to be installed and installs it if required.
// 8. Checks if Kafka Operator addon is to be installed and installs it if required.
// 9. Returns any errors encountered during the installation process.
func clusterAddons(ctx *pulumi.Context, locals *localz.Locals,
	createdCluster *container.Cluster, gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {

	if err := addons.GatewayApis(ctx, kubernetesProvider); err != nil {
		return errors.Wrap(err, "failed to add gateway-api resources")
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallIngressNginx {
		if err := addons.IngressNginx(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install ingress-nginx resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallIstio {
		if err := addons.Istio(ctx, locals, createdCluster, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install istio resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallCertManager {
		if err := addons.CertManager(ctx,
			locals,
			createdCluster,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install cert-manager resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallExternalSecrets {
		if err := addons.ExternalSecrets(ctx,
			locals,
			createdCluster,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-secrets resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallExternalDns {
		if err := addons.ExternalDns(ctx, locals, createdCluster, gcpProvider, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-dns resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallPostgresOperator {
		if err := addons.ZalandoPostgresOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install zalando postgres-operator resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallSolrOperator {
		if err := addons.SolrOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install solr-operator resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallKafkaOperator {
		if err := addons.StrimziKafkaOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install strimzi-kafka-operator resources")
		}
	}

	if locals.GkeCluster.Spec.KubernetesAddons.IsInstallElasticOperator {
		if err := addons.ElasticOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install elastic-operator resources")
		}
	}

	return nil
}
