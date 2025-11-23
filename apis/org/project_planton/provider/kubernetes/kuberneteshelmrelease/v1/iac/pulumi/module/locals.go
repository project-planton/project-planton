package module

import (
	"strconv"

	kuberneteshelmreleasev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteshelmrelease/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	Labels                map[string]string
	Namespace             string
	KubernetesHelmRelease *kuberneteshelmreleasev1.KubernetesHelmRelease
}

func initializeLocals(ctx *pulumi.Context, stackInput *kuberneteshelmreleasev1.KubernetesHelmReleaseStackInput) *Locals {
	locals := &Locals{}

	locals.KubernetesHelmRelease = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesHelmRelease.String(),
	}

	if target.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// Priority order for namespace:
	// 1. spec.namespace (required field with StringValueOrRef)
	// 2. Override with stackInput.KubernetesNamespace if provided
	// 3. Fall back to metadata.name if spec.namespace is empty

	locals.Namespace = target.Spec.Namespace.GetValue()

	if locals.Namespace == "" {
		locals.Namespace = target.Metadata.Name
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	return locals
}
