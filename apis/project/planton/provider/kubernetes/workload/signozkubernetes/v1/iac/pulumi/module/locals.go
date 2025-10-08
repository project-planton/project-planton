package module

import (
	"fmt"
	"strconv"

	signozkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	SignozKubernetes                  *signozkubernetesv1.SignozKubernetes
	Namespace                         string
	KubernetesLabels                  map[string]string
	SignozServiceName                 string
	OtelCollectorServiceName          string
	SignozKubeServiceFqdn             string
	OtelCollectorGrpcFqdn             string
	OtelCollectorHttpFqdn             string
	KubePortForwardCommand            string
	IngressExternalHostname           string
	IngressInternalHostname           string
	IngressHostnames                  []string
	IngressCertClusterIssuerName      string
	IngressCertSecretName             string
	OtelCollectorExternalGrpcHostname string
	OtelCollectorExternalHttpHostname string
	ClickhouseEndpoint                string
}

func initializeLocals(ctx *pulumi.Context, stackInput *signozkubernetesv1.SignozKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.SignozKubernetes = stackInput.Target
	target := stackInput.Target

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_SignozKubernetes.String(),
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

	// Priority order for namespace:
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

	//export namespace
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// Service names
	locals.SignozServiceName = fmt.Sprintf("%s-signoz", target.Metadata.Name)
	locals.OtelCollectorServiceName = fmt.Sprintf("%s-otel-collector", target.Metadata.Name)

	//export service names
	ctx.Export(OpSignozService, pulumi.String(locals.SignozServiceName))
	ctx.Export(OpOtelCollectorService, pulumi.String(locals.OtelCollectorServiceName))

	// Kubernetes FQDNs
	locals.SignozKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.SignozServiceName, locals.Namespace, vars.SignozUIPort)
	locals.OtelCollectorGrpcFqdn = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.OtelCollectorServiceName, locals.Namespace, vars.OtelGrpcPort)
	locals.OtelCollectorHttpFqdn = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.OtelCollectorServiceName, locals.Namespace, vars.OtelHttpPort)

	//export kubernetes endpoints
	ctx.Export(OpKubeEndpoint, pulumi.String(locals.SignozKubeServiceFqdn))
	ctx.Export(OpOtelCollectorGrpcEndpoint, pulumi.String(locals.OtelCollectorGrpcFqdn))
	ctx.Export(OpOtelCollectorHttpEndpoint, pulumi.String(locals.OtelCollectorHttpFqdn))

	// Port forward command
	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.SignozServiceName, vars.SignozUIPort, vars.SignozUIPort)

	//export kube-port-forward command
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	// ClickHouse outputs (only for self-managed)
	if target.Spec.Database != nil && !target.Spec.Database.IsExternal {
		locals.ClickhouseEndpoint = fmt.Sprintf("%s-clickhouse.%s.svc.cluster.local:8123",
			target.Metadata.Name, locals.Namespace)
		ctx.Export(OpClickhouseEndpoint, pulumi.String(locals.ClickhouseEndpoint))
		ctx.Export(OpClickhouseUsername, pulumi.String("admin"))
		ctx.Export(OpClickhousePasswordSecretName, pulumi.String(fmt.Sprintf("%s-clickhouse", target.Metadata.Name)))
		ctx.Export(OpClickhousePasswordSecretKey, pulumi.String("admin-password"))
	}

	// Ingress configuration for SigNoz UI
	if target.Spec.SignozIngress != nil &&
		target.Spec.SignozIngress.Enabled &&
		target.Spec.SignozIngress.DnsDomain != "" {
		locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
			target.Spec.SignozIngress.DnsDomain)
		ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))

		locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
			target.Spec.SignozIngress.DnsDomain)
		ctx.Export(OpInternalHostname, pulumi.String(locals.IngressInternalHostname))

		locals.IngressHostnames = []string{
			locals.IngressExternalHostname,
			locals.IngressInternalHostname,
		}

		// ClusterIssuer should already exist on the cluster
		// Typically managed by cluster administrator or created by Planton Cloud
		locals.IngressCertClusterIssuerName = target.Spec.SignozIngress.DnsDomain

		locals.IngressCertSecretName = fmt.Sprintf("cert-%s", locals.Namespace)
	}

	// Ingress configuration for OTel Collector
	if target.Spec.OtelCollectorIngress != nil &&
		target.Spec.OtelCollectorIngress.Enabled &&
		target.Spec.OtelCollectorIngress.DnsDomain != "" {
		locals.OtelCollectorExternalGrpcHostname = fmt.Sprintf("%s-ingest-grpc.%s", locals.Namespace,
			target.Spec.OtelCollectorIngress.DnsDomain)
		ctx.Export(OpOtelCollectorExternalGrpcHostname, pulumi.String(locals.OtelCollectorExternalGrpcHostname))

		locals.OtelCollectorExternalHttpHostname = fmt.Sprintf("%s-ingest-http.%s", locals.Namespace,
			target.Spec.OtelCollectorIngress.DnsDomain)
		ctx.Export(OpOtelCollectorExternalHttpHostname, pulumi.String(locals.OtelCollectorExternalHttpHostname))
	}

	return locals
}
