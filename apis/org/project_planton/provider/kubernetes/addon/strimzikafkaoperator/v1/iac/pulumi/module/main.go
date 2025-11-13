package module

import (
	"github.com/pkg/errors"
	strimzikafkaoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/addon/strimzikafkaoperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry‑point called by the Project Planton engine.
func Resources(
	ctx *pulumi.Context,
	stackInput *strimzikafkaoperatorv1.KafkaOperatorKubernetesStackInput,
) error {
	// ------------------------------------------------------------------
	// Provider set‑up
	// ------------------------------------------------------------------
	k8sProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx,
		stackInput.ProviderConfig,
		"kubernetes",
	)
	if err != nil {
		return errors.Wrap(err, "failed to set up Kubernetes provider")
	}

	// ------------------------------------------------------------------
	// Helm install
	// ------------------------------------------------------------------
	if err := kafkaOperator(ctx, stackInput.Target, k8sProvider); err != nil {
		return errors.Wrap(err, "failed to install Kafka operator resources")
	}

	return nil
}
