package module

import (
	"encoding/base64"

	"github.com/pkg/errors"
	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *clickhousekubernetesv1.ClickhouseKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.ProviderCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: kubernetesmetav1.ObjectMetaPtrInput(
				&kubernetesmetav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//create new random secret to set as password
	createdRandomString, err := random.NewRandomPassword(ctx,
		"root-password",
		&random.RandomPasswordArgs{
			Length:     pulumi.Int(16),
			Special:    pulumi.Bool(true),
			Numeric:    pulumi.Bool(true),
			Upper:      pulumi.Bool(true),
			Lower:      pulumi.Bool(true),
			MinSpecial: pulumi.Int(3),
			MinNumeric: pulumi.Int(2),
			MinUpper:   pulumi.Int(2),
			MinLower:   pulumi.Int(2),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to generate random password value")
	}

	// Encode the password in Base64
	createdBase64Password := createdRandomString.Result.ApplyT(func(p string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(p)), nil
	}).(pulumi.StringOutput)

	// create kubernetes secret to store generated password
	_, err = kubernetescorev1.NewSecret(ctx,
		locals.ClickhouseKubernetes.Metadata.Name,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.ClickhouseKubernetes.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
			},
			Data: pulumi.StringMap{
				vars.ClickhousePasswordKey: createdBase64Password,
			},
		}, pulumi.Parent(createdNamespace))

	//create clickhouse using helm-chart
	if err := clickhouse(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create clickhouse helm-chart resources")
	}

	//create service of type load-balancer if ingress is enabled.
	if locals.ClickhouseKubernetes.Spec.Ingress != nil && locals.ClickhouseKubernetes.Spec.Ingress.Enabled {
		_, err := kubernetescorev1.NewService(ctx,
			"ingress-external-lb",
			&kubernetescorev1.ServiceArgs{
				Metadata: &kubernetesmetav1.ObjectMetaArgs{
					Name:      pulumi.String("ingress-external-lb"),
					Namespace: createdNamespace.Metadata.Name(),
					Labels:    createdNamespace.Metadata.Labels(),
					Annotations: pulumi.StringMap{
						"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressExternalHostname),
					},
				},
				Spec: &kubernetescorev1.ServiceSpecArgs{
					Type: pulumi.String("LoadBalancer"), // Service type is LoadBalancer
					Ports: kubernetescorev1.ServicePortArray{
						&kubernetescorev1.ServicePortArgs{
							Name:       pulumi.String("http"),
							Port:       pulumi.Int(vars.ClickhouseHttpPort),
							Protocol:   pulumi.String("TCP"),
							TargetPort: pulumi.String("http"),
						},
						&kubernetescorev1.ServicePortArgs{
							Name:       pulumi.String("tcp"),
							Port:       pulumi.Int(vars.ClickhouseNativePort),
							Protocol:   pulumi.String("TCP"),
							TargetPort: pulumi.String("tcp"),
						},
					},
					Selector: pulumi.ToStringMap(locals.ClickhousePodSelectorLabels),
				},
			}, pulumi.Provider(kubernetesProvider), pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrapf(err, "failed to create external load balancer service")
		}
	}

	return nil
}
