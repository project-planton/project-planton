package module

import (
	"fmt"
	"strconv"

	clickhousekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	IngressExternalHostname     string
	IngressInternalHostname     string
	KubePortForwardCommand      string
	KubeServiceFqdn             string
	KubeServiceName             string
	ClickHouseKubernetes        *clickhousekubernetesv1.ClickHouseKubernetes
	Namespace                   string
	ClickhousePodSelectorLabels map[string]string
	KubernetesLabels            map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *clickhousekubernetesv1.ClickHouseKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.ClickHouseKubernetes = stackInput.Target

	target := stackInput.Target

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_ClickHouseKubernetes.String(),
	}

	if target.Metadata.Id != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

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
	ctx.Export(OpUsername, pulumi.String(vars.DefaultUsername))
	ctx.Export(OpPasswordSecretName, pulumi.String(target.Metadata.Name))
	ctx.Export(OpPasswordSecretKey, pulumi.String(vars.ClickhousePasswordKey))

	locals.KubeServiceName = target.Metadata.Name

	// Determine cluster name - use spec.cluster_name if provided, otherwise use metadata.name
	clusterName := target.Spec.ClusterName
	if clusterName == "" {
		clusterName = target.Metadata.Name
	}

	// Altinity operator uses these labels for pod selection
	// These labels are automatically applied by the operator to ClickHouse pods
	locals.ClickhousePodSelectorLabels = map[string]string{
		"clickhouse.altinity.com/chi":     clusterName,
		"clickhouse.altinity.com/cluster": clusterName,
	}

	//export kubernetes service name
	ctx.Export(OpService, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local:%d", locals.KubeServiceName, locals.Namespace, vars.ClickhouseHttpPort)

	//export kubernetes endpoint
	ctx.Export(OpKubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, target.Metadata.Name, vars.ClickhouseHttpPort, vars.ClickhouseHttpPort)

	//export kube-port-forward command
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled ||
		target.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	//export ingress-external-hostname
	ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	//export ingress-internal-hostname
	ctx.Export(OpInternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
