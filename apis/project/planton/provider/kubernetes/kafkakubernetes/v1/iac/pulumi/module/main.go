package module

import (
	"github.com/pkg/errors"
	kafkakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kafkakubernetesv1.KafkaKubernetesStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-kowlConfigTemplateInput
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
					Labels: pulumi.ToStringMap(locals.Labels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//create kafka cluster custom resource
	createdKafkaCluster, err := kafkaCluster(ctx, locals, createdNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to create kafka-cluster resources")
	}

	//create kafka admin user
	if err := kafkaAdminUser(ctx, locals, createdNamespace, createdKafkaCluster); err != nil {
		return errors.Wrap(err, "failed to create kafka admin user")
	}

	//create kafka topics
	if err := kafkaTopics(ctx, locals, createdNamespace, createdKafkaCluster); err != nil {
		return errors.Wrap(err, "failed to create kafka topics")
	}

	//create schema-registry
	if locals.KafkaKubernetes.Spec.SchemaRegistryContainer != nil &&
		locals.KafkaKubernetes.Spec.SchemaRegistryContainer.IsEnabled {
		if err := schemaRegistry(ctx, locals, kubernetesProvider, createdNamespace, createdKafkaCluster); err != nil {
			return errors.Wrap(err, "failed to create schema registry deployment")
		}
	}

	//create kowl
	if locals.KafkaKubernetes.Spec.IsDeployKafkaUi {
		if err := kowl(ctx, locals, kubernetesProvider, createdNamespace, createdKafkaCluster); err != nil {
			return errors.Wrap(err, "failed to create kowl deployment")
		}
	}
	return nil
}
