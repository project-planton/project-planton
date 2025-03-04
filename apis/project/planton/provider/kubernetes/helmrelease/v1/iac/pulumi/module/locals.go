package module

import (
	helmreleasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	Labels      map[string]string
	Namespace   string
	HelmRelease *helmreleasev1.HelmRelease
}

func initializeLocals(ctx *pulumi.Context, stackInput *helmreleasev1.HelmReleaseStackInput) *Locals {
	locals := &Locals{}

	locals.HelmRelease = stackInput.Target

	helmRelease := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: helmRelease.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "helm_release",
	}

	if helmRelease.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = helmRelease.Metadata.Id
	}

	if helmRelease.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = helmRelease.Metadata.Org
	}

	if helmRelease.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = helmRelease.Metadata.Env
	}

	locals.Namespace = helmRelease.Metadata.Name

	if helmRelease.Metadata.Labels != nil &&
		helmRelease.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = helmRelease.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	return locals
}
