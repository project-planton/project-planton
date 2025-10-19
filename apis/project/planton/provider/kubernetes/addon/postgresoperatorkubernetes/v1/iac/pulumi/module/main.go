package module

import (
	"github.com/pkg/errors"
	postgresoperatorkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/postgresoperatorkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi entry‑point invoked by the Project‑Planton CLI.
func Resources(ctx *pulumi.Context, stackInput *postgresoperatorkubernetesv1.PostgresOperatorKubernetesStackInput) error {
	// Translate incoming protobuf‑generated types into helper data we
	//                need throughout the module (labels, metadata, etc.).
	locals := initializeLocals(ctx, stackInput)

	// Create a Kubernetes provider from the supplied cluster credential.
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(
		ctx,
		stackInput.ProviderConfig,
		"kubernetes",
	)
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	// Install / upgrade the Zalando Postgres‑Operator.
	if err := postgresOperator(ctx, locals, kubernetesProvider); err != nil {
		return errors.Wrap(err, "failed to install postgres-operator resources")
	}

	return nil
}
