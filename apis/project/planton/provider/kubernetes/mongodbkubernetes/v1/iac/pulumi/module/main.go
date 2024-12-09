package module

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	mongodbkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *mongodbkubernetesv1.MongodbKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
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
		},
		pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "5s", Update: "5s", Delete: "5s"}),
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//create new random secret to set as password
	createdRandomString, err := random.NewRandomPassword(ctx,
		"root-password",
		&random.RandomPasswordArgs{
			Length:     pulumi.Int(12),
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
		return fmt.Errorf("failed to generate random password value: %w", err)
	}

	// Encode the password in Base64
	createdBase64Password := createdRandomString.Result.ApplyT(func(p string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(p)), nil
	}).(pulumi.StringOutput)

	// create kubernetes secret to store generated password
	_, err = kubernetescorev1.NewSecret(ctx,
		locals.MongodbKubernetes.Metadata.Name,
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.MongodbKubernetes.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
			},
			Data: pulumi.StringMap{
				vars.MongodbRootPasswordKey: createdBase64Password,
			},
		}, pulumi.Parent(createdNamespace))

	//create mongodb using helm-chart
	if err := mongodb(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create mongodb helm-chart resources")
	}

	//create service of type load-balancer if ingress is enabled.
	if locals.MongodbKubernetes.Spec.Ingress.IsEnabled {
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
							Name:       pulumi.String("tcp-mongodb"),
							Port:       pulumi.Int(vars.MongoDbPort),
							Protocol:   pulumi.String("TCP"),
							TargetPort: pulumi.String("mongodb"),
						},
					},
					Selector: pulumi.ToStringMap(locals.MongodbPodSelectorLabels),
				},
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrapf(err, "failed to create external load balancer service")
		}
	}

	return nil
}
