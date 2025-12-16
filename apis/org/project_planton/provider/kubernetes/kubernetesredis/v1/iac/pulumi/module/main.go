package module

import (
	"github.com/pkg/errors"
	kubernetesredisv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesredis/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kubernetesredisv1.KubernetesRedisStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	// Create or reference namespace based on create_namespace flag
	namespace, err := createOrGetNamespace(ctx, locals, stackInput.Target.Spec, kubernetesProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create or get namespace")
	}

	if err := adminPassword(ctx, locals, namespace); err != nil {
		return errors.Wrap(err, "failed to create admin secret")
	}
	//install the redis helm-chart
	if err := helmChart(ctx, locals, namespace); err != nil {
		return errors.Wrap(err, "failed to create helm-chart resources")
	}

	//if ingress is enabled, create load-balancer ingress resources
	if locals.KubernetesRedis.Spec.Ingress != nil &&
		locals.KubernetesRedis.Spec.Ingress.Enabled &&
		locals.KubernetesRedis.Spec.Ingress.Hostname != "" {
		if err := loadBalancerIngress(ctx, locals, namespace); err != nil {
			return errors.Wrap(err, "failed to create load-balancer ingress resources")
		}
	}

	return nil
}
