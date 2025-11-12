package module

import (
	"github.com/pkg/errors"
	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/workload/temporalkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the single entry-point consumed by ProjectPlanton runtime.
// It wires together all noun-style helpers in a Terraform-like, top-down
// order so the flow is easy for DevOps engineers to follow.
func Resources(ctx *pulumi.Context,
	stackInput *temporalkubernetesv1.TemporalKubernetesStackInput) error {

	locals := initializeLocals(ctx, stackInput)

	// external_database is required when the backend is not cassandra
	if locals.TemporalKubernetes.Spec.Database.Backend != temporalkubernetesv1.TemporalKubernetesDatabaseBackend_cassandra &&
		locals.TemporalKubernetes.Spec.Database.ExternalDatabase == nil {

		return errors.New("external_database must be provided when backend is not cassandra")
	}

	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesProviderConfig(ctx,
		stackInput.ProviderConfig, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	createdNamespace, err := namespace(ctx, locals, kubernetesProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	if err := dbPasswordSecret(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create database password secret")
	}

	if err := helmChart(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to install Temporal Helm chart")
	}

	if err := frontendIngress(ctx, locals, createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create frontend gRPC ingress")
	}

	if err := frontendHttpIngress(
		ctx,
		locals,
		kubernetesProvider,
		createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create frontend HTTP ingress")
	}

	if err := webUiIngress(
		ctx,
		locals,
		kubernetesProvider,
		createdNamespace); err != nil {
		return errors.Wrap(err, "failed to create web UI ingress")
	}

	return nil
}
