package module

import (
	"fmt"

	"github.com/pkg/errors"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// nackController deploys the NACK (NATS Controllers for Kubernetes) Helm chart.
// NACK provides a controller that watches JetStream CRDs and reconciles them
// to actual JetStream resources in the NATS cluster.
//
// This must be deployed after:
// 1. NATS Helm chart (provides the NATS server to connect to)
// 2. NACK CRDs (so the controller can watch them)
//
// Reference: https://github.com/nats-io/nack
func nackController(ctx *pulumi.Context, locals *Locals, kubernetesProvider pulumi.ProviderResource,
	nackCrds pulumi.Resource) (pulumi.Resource, error) {

	// Skip if NACK controller is not enabled
	if locals.KubernetesNats.Spec.NackController == nil ||
		!locals.KubernetesNats.Spec.NackController.Enabled {
		return nil, nil
	}

	// Build Helm values for NACK controller
	values := pulumi.Map{
		"jetstream": pulumi.Map{
			// Enable the JetStream controller
			"enabled": pulumi.Bool(true),
			// NATS URL - points to the NATS service deployed by the NATS Helm chart
			"nats": pulumi.Map{
				"url": pulumi.String(locals.ClientURLInternal),
			},
		},
	}

	// Enable control-loop mode if requested
	// Control-loop mode is required for KeyValue, ObjectStore, and Account support
	// and provides more reliable state enforcement
	if locals.KubernetesNats.Spec.NackController.EnableControlLoop {
		values["jetstream"].(pulumi.Map)["additionalArgs"] = pulumi.ToStringArray([]string{"--control-loop"})
	}

	// Build dependencies - NACK controller must wait for CRDs to be registered
	var deps []pulumi.Resource
	if nackCrds != nil {
		deps = append(deps, nackCrds)
	}

	// Deploy NACK Helm chart
	releaseName := fmt.Sprintf("%s-nack", locals.KubernetesNats.Metadata.Name)
	nackRelease, err := helmv3.NewRelease(ctx, releaseName,
		&helmv3.ReleaseArgs{
			Chart:     pulumi.String(vars.NackHelmChartName),
			Version:   pulumi.String(locals.NackHelmChartVersion),
			Namespace: pulumi.String(locals.Namespace),
			RepositoryOpts: helmv3.RepositoryOptsArgs{
				Repo: pulumi.String(vars.NackHelmChartRepoUrl),
			},
			Values: values,
			// Skip CRDs - we install them separately for better control
			SkipCrds: pulumi.Bool(true),
			// Timeout for the release to be ready (5 minutes)
			Timeout: pulumi.Int(300),
		},
		pulumi.Provider(kubernetesProvider),
		pulumi.DependsOn(deps),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy NACK controller")
	}

	// Export NACK controller status
	ctx.Export(OpNackControllerEnabled, pulumi.Bool(true))
	ctx.Export(OpNackControllerVersion, pulumi.String(locals.NackHelmChartVersion))

	return nackRelease, nil
}
