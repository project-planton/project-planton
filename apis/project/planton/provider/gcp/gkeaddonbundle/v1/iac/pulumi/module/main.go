package module

import (
	"github.com/pkg/errors"
	gkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkeaddonbundle/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gkeaddonbundlev1.GkeAddonBundleStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create gcp-provider using the gcp-credential from input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup google provider")
	}

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup kubernetes provider")
	}

	if err := gatewayApis(ctx, kubernetesProvider); err != nil {
		return errors.Wrap(err, "failed to add gateway-api resources")
	}

	if locals.GkeAddonBundle.Spec.InstallIngressNginx {
		if err := ingressNginx(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install ingress-nginx resources")
		}
	}

	if locals.GkeAddonBundle.Spec.Istio != nil && locals.GkeAddonBundle.Spec.Istio.Enabled {
		if err := istio(ctx, locals, gcpProvider, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install istio resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallCertManager {
		if err := certManager(ctx,
			locals,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install cert-manager resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallExternalSecrets {
		if err := externalSecrets(ctx,
			locals,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-secrets resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallExternalDns {
		if err := externalDns(ctx, locals, gcpProvider, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-dns resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallPostgresOperator {
		if err := zalandoPostgresOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install zalando postgres-operator resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallSolrOperator {
		if err := solrOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install solr-operator resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallKafkaOperator {
		if err := strimziKafkaOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install strimzi-kafka-operator resources")
		}
	}

	if locals.GkeAddonBundle.Spec.InstallElasticOperator {
		if err := elasticOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install elastic-operator resources")
		}
	}

	return nil
}
