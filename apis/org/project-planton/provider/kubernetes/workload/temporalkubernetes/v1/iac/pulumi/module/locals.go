package module

import (
	"fmt"
	"strconv"

	temporalkubernetesv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/kubernetes/workload/temporalkubernetes/v1"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps all frequently-used values in one place – similar to a
// Terraform “locals {}” block.
type Locals struct {
	Namespace                   string
	Labels                      map[string]string
	TemporalKubernetes          *temporalkubernetesv1.TemporalKubernetes
	FrontendServiceName         string
	UIServiceName               string
	FrontendEndpoint            string
	UIEndpoint                  string
	PortForwardFrontendCmd      string
	PortForwardUICmd            string
	IngressFrontendGrpcHostname string
	IngressFrontendHttpHostname string
	IngressUIHostname           string
}

// initializeLocals builds the Locals struct and immediately exports the
// values required by TemporalKubernetesStackOutputs.
func initializeLocals(ctx *pulumi.Context,
	stackInput *temporalkubernetesv1.TemporalKubernetesStackInput) *Locals {

	locals := &Locals{}
	locals.TemporalKubernetes = stackInput.Target
	target := stackInput.Target

	// -------------------------------- labels ---------------------------------
	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_TemporalKubernetes.String(),
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

	// ------------------------------- namespace -------------------------------
	// Priority order:
	// 1. Default: metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// ---------------------------- service names ------------------------------
	locals.FrontendServiceName = fmt.Sprintf("%s-frontend", target.Metadata.Name)
	locals.UIServiceName = fmt.Sprintf("%s-web", target.Metadata.Name)

	ctx.Export(OpFrontendService, pulumi.String(locals.FrontendServiceName))
	ctx.Export(OpUIService, pulumi.String(locals.UIServiceName))

	// --------------------------- cluster endpoints ---------------------------
	locals.FrontendEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.FrontendServiceName, locals.Namespace, vars.FrontendGrpcPort)
	locals.UIEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.UIServiceName, locals.Namespace, vars.UIPort)

	ctx.Export(OpFrontendEndpoint, pulumi.String(locals.FrontendEndpoint))
	ctx.Export(OpWebUiEndpoint, pulumi.String(locals.UIEndpoint))

	// --------------------------- port-forward cmds ---------------------------
	locals.PortForwardFrontendCmd = fmt.Sprintf(
		"kubectl port-forward -n %s service/%s 7233:7233",
		locals.Namespace, locals.FrontendServiceName)
	locals.PortForwardUICmd = fmt.Sprintf(
		"kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.UIServiceName)

	ctx.Export(OpPortForwardFrontendCommand, pulumi.String(locals.PortForwardFrontendCmd))
	ctx.Export(OpPortForwardUICommand, pulumi.String(locals.PortForwardUICmd))

	// ------------------------------- ingress ---------------------------------
	// Frontend ingress
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Frontend != nil &&
		target.Spec.Ingress.Frontend.Enabled {

		if target.Spec.Ingress.Frontend.GrpcHostname != "" {
			locals.IngressFrontendGrpcHostname = target.Spec.Ingress.Frontend.GrpcHostname
			ctx.Export(OpExternalFrontendHostname, pulumi.String(locals.IngressFrontendGrpcHostname))
		}

		if target.Spec.Ingress.Frontend.HttpHostname != "" {
			locals.IngressFrontendHttpHostname = target.Spec.Ingress.Frontend.HttpHostname
		}
	}

	// Web UI ingress
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.WebUi != nil &&
		target.Spec.Ingress.WebUi.Enabled &&
		target.Spec.Ingress.WebUi.Hostname != "" {

		locals.IngressUIHostname = target.Spec.Ingress.WebUi.Hostname
		ctx.Export(OpExternalUIHostname, pulumi.String(locals.IngressUIHostname))
	}

	return locals
}
