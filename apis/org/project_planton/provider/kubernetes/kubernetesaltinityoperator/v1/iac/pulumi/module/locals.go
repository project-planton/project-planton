package module

import (
	kubernetesaltinityoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesaltinityoperator/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// locals contains computed values and data transformations for the Altinity operator deployment
type locals struct {
	// Namespace is the resolved namespace where the operator will be installed
	Namespace string

	// HelmValues is the prepared configuration map for the Helm chart
	HelmValues pulumi.Map
}

// newLocals creates and initializes local values from the stack input
func newLocals(stackInput *kubernetesaltinityoperatorv1.KubernetesAltinityOperatorStackInput) *locals {
	l := &locals{}

	// Determine namespace - use from spec or default
	l.Namespace = stackInput.Target.Spec.Namespace.GetValue()
	if l.Namespace == "" {
		l.Namespace = vars.DefaultNamespace
	}

	// Prepare helm values with CRD installation enabled and resource limits from spec
	l.HelmValues = pulumi.Map{
		"operator": pulumi.Map{
			"createCRD": pulumi.Bool(true),
			"resources": pulumi.Map{
				"limits": pulumi.Map{
					"cpu":    pulumi.String(stackInput.Target.Spec.Container.Resources.Limits.Cpu),
					"memory": pulumi.String(stackInput.Target.Spec.Container.Resources.Limits.Memory),
				},
				"requests": pulumi.Map{
					"cpu":    pulumi.String(stackInput.Target.Spec.Container.Resources.Requests.Cpu),
					"memory": pulumi.String(stackInput.Target.Spec.Container.Resources.Requests.Memory),
				},
			},
		},
		// Configure operator to watch all namespaces cluster-wide
		// Use regex pattern ".*" to match all namespaces (empty array watches only installation namespace)
		"configs": pulumi.Map{
			"files": pulumi.Map{
				"config.yaml": pulumi.Map{
					"watch": pulumi.Map{
						"namespaces": pulumi.Array{
							pulumi.String(".*"),
						},
					},
				},
			},
		},
	}

	return l
}
