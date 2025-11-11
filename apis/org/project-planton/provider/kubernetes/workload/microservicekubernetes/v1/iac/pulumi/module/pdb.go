package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	policyv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/policy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// podDisruptionBudget creates a PodDisruptionBudget resource if configured.
// PodDisruptionBudgets ensure minimum availability during voluntary disruptions like node maintenance.
func podDisruptionBudget(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	// Check if PDB is enabled
	pdbConfig := locals.MicroserviceKubernetes.Spec.Availability.PodDisruptionBudget
	if pdbConfig == nil || !pdbConfig.Enabled {
		return nil
	}

	pdbSpec := &policyv1.PodDisruptionBudgetSpecArgs{
		Selector: &metav1.LabelSelectorArgs{
			MatchLabels: pulumi.ToStringMap(locals.Labels),
		},
	}

	// Set minAvailable or maxUnavailable (they are mutually exclusive)
	if pdbConfig.MinAvailable != "" {
		pdbSpec.MinAvailable = parseIntOrString(pdbConfig.MinAvailable)
	} else if pdbConfig.MaxUnavailable != "" {
		pdbSpec.MaxUnavailable = parseIntOrString(pdbConfig.MaxUnavailable)
	} else {
		// Default to minAvailable: 1 if neither is specified
		pdbSpec.MinAvailable = pulumi.Int(1)
	}

	_, err := policyv1.NewPodDisruptionBudget(ctx,
		"pdb",
		&policyv1.PodDisruptionBudgetArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.MicroserviceKubernetes.Metadata.Name),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: pdbSpec,
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create pod disruption budget")
	}

	return nil
}
