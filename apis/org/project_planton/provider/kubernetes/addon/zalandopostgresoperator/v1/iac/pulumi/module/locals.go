package module

import (
	"strconv"

	zalandopostgresoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/addon/zalandopostgresoperator/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals collect computed values that are reused across resources.
type Locals struct {
	PostgresOperatorKubernetes *zalandopostgresoperatorv1.PostgresOperatorKubernetes
	KubernetesLabels           map[string]string
}

// initializeLocals builds the Locals struct once and re‑uses it elsewhere.
func initializeLocals(_ *pulumi.Context, stackInput *zalandopostgresoperatorv1.PostgresOperatorKubernetesStackInput) *Locals {
	target := stackInput.Target

	kubeLabels := map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "PostgresOperatorKubernetes",
	}

	if target.Metadata.Id != "" {
		kubeLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}
	if target.Metadata.Org != "" {
		kubeLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}
	if target.Metadata.Env != "" {
		kubeLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	return &Locals{
		PostgresOperatorKubernetes: target,
		KubernetesLabels:           kubeLabels,
	}
}
