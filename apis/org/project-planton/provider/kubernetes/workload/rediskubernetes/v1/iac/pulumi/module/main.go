package module

import (
	"github.com/pkg/errors"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/workload/rediskubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *rediskubernetesv1.RedisKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
				Name:   pulumi.String(locals.Namespace),
				Labels: pulumi.ToStringMap(locals.Labels),
			}),
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	if err := adminPassword(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create admin secret")
	}
	//install the redis helm-chart
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create helm-chart resources")
	}

	//if ingress is enabled, create load-balancer ingress resources
	if locals.RedisKubernetes.Spec.Ingress != nil &&
		locals.RedisKubernetes.Spec.Ingress.Enabled &&
		locals.RedisKubernetes.Spec.Ingress.Hostname != "" {
		if err := loadBalancerIngress(ctx, locals, createdNamespace); err != nil {
			return errors.Wrap(err, "failed to create load-balancer ingress resources")
		}
	}

	return nil
}
