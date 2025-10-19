package module

import (
	"github.com/pkg/errors"
	gcpgkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *gcpgkeaddonbundlev1.GcpGkeAddonBundleStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create gcp-provider using the gcp-credential from input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup google provider")
	}

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.KubernetesProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup kubernetes provider")
	}

	if err := gatewayApis(ctx, kubernetesProvider); err != nil {
		return errors.Wrap(err, "failed to add gateway-api resources")
	}

	if locals.GcpGkeAddonBundle.Spec.InstallIngressNginx {
		if err := ingressNginx(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install ingress-nginx resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.Istio != nil && locals.GcpGkeAddonBundle.Spec.Istio.Enabled {
		if err := istio(ctx, locals, gcpProvider, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install istio resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallCertManager {
		if err := certManager(ctx,
			locals,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install cert-manager resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallExternalSecrets {
		if err := externalSecrets(ctx,
			locals,
			gcpProvider,
			kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-secrets resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallExternalDns {
		if err := externalDns(ctx, locals, gcpProvider, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install external-dns resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallPostgresOperator {
		if err := zalandoPostgresOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install zalando postgres-operator resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallSolrOperator {
		if err := solrOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install solr-operator resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallKafkaOperator {
		if err := strimziKafkaOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install strimzi-kafka-operator resources")
		}
	}

	if locals.GcpGkeAddonBundle.Spec.InstallElasticOperator {
		if err := elasticOperator(ctx, locals, kubernetesProvider); err != nil {
			return errors.Wrap(err, "failed to install elastic-operator resources")
		}
	}

	return nil
}
