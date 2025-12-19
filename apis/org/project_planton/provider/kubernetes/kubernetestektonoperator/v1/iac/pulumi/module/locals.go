package module

import (
	"fmt"
	"strconv"

	kubernetestektonoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetestektonoperator/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	KubernetesTektonOperator *kubernetestektonoperatorv1.KubernetesTektonOperator
	KubeLabels               map[string]string
	OperatorNamespace        string
	ComponentsNamespace      string
	TektonConfigName         string

	// Component enablement
	EnablePipelines bool
	EnableTriggers  bool
	EnableDashboard bool
}

func initializeLocals(ctx *pulumi.Context, in *kubernetestektonoperatorv1.KubernetesTektonOperatorStackInput) *Locals {
	var l Locals
	l.KubernetesTektonOperator = in.Target

	l.KubeLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: in.Target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesTektonOperator.String(),
	}

	if id := in.Target.Metadata.Id; id != "" {
		l.KubeLabels[kuberneteslabelkeys.ResourceId] = id
	}
	if org := in.Target.Metadata.Org; org != "" {
		l.KubeLabels[kuberneteslabelkeys.Organization] = org
	}
	if env := in.Target.Metadata.Env; env != "" {
		l.KubeLabels[kuberneteslabelkeys.Environment] = env
	}

	// Use default namespaces for Tekton Operator
	l.OperatorNamespace = vars.OperatorNamespace
	l.ComponentsNamespace = vars.ComponentsNamespace
	l.TektonConfigName = vars.TektonConfigName

	// Get component enablement from spec
	if comp := in.Target.Spec.Components; comp != nil {
		l.EnablePipelines = comp.Pipelines
		l.EnableTriggers = comp.Triggers
		l.EnableDashboard = comp.Dashboard
	}

	// Export namespace as an output
	ctx.Export(OpNamespace, pulumi.String(l.ComponentsNamespace))
	ctx.Export(OpTektonConfigName, pulumi.String(l.TektonConfigName))

	// Export service names based on enabled components
	if l.EnablePipelines {
		ctx.Export(OpPipelinesControllerService, pulumi.String("tekton-pipelines-controller"))
	}
	if l.EnableTriggers {
		ctx.Export(OpTriggersControllerService, pulumi.String("tekton-triggers-controller"))
	}
	if l.EnableDashboard {
		ctx.Export(OpDashboardService, pulumi.String("tekton-dashboard"))
		ctx.Export(OpDashboardPortForwardCommand, pulumi.String(
			fmt.Sprintf("kubectl port-forward svc/tekton-dashboard -n %s 9097:9097", l.ComponentsNamespace)))
	}

	return &l
}
