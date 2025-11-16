package module

import (
	"github.com/pkg/errors"
	kubernetesgrafanav1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesgrafana/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kubernetesgrafanav1.KubernetesGrafanaStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup kubernetes provider")
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

	//install the grafana helm-chart
	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create helm-chart resources")
	}

	//if ingress is enabled, create load-balancer ingress resources
	if locals.KubernetesGrafana.Spec.Ingress != nil &&
		locals.KubernetesGrafana.Spec.Ingress.Enabled {
		if err := ingress(ctx, locals, createdNamespace); err != nil {
			return errors.Wrap(err, "failed to create ingress")
		}
	}

	return nil
}
