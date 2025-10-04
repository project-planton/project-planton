package module

import (
	"fmt"
	"strconv"

	elasticsearchkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/elasticsearchkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	ElasticsearchIngressExternalHostname string
	ElasticsearchIngressInternalHostname string
	ElasticsearchKubePortForwardCommand  string
	ElasticsearchKubeServiceFqdn         string
	ElasticsearchKubeServiceName         string
	Namespace                            string
	ElasticsearchKubernetes              *elasticsearchkubernetesv1.ElasticsearchKubernetes
	KibanaIngressExternalHostname        string
	KibanaIngressInternalHostname        string
	KibanaKubePortForwardCommand         string
	KibanaKubeServiceFqdn                string
	KibanaKubeServiceName                string
	IngressHostnames                     []string
	IngressCertClusterIssuerName         string
	IngressCertSecretName                string
	Labels                               map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *elasticsearchkubernetesv1.ElasticsearchKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.ElasticsearchKubernetes = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_ElasticsearchKubernetes.String(),
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

	ctx.Export(OpElasticsearchUsername, pulumi.String("elastic"))
	ctx.Export(OpElasticsearchPasswordSecretName, pulumi.Sprintf("%s-es-elastic-user", target.Metadata.Name))
	ctx.Export(OpElasticsearchPasswordSecretKey, pulumi.String("elastic"))

	locals.ElasticsearchKubeServiceName = fmt.Sprintf("%s-es-http", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(OpElasticsearchService, pulumi.String(locals.ElasticsearchKubeServiceName))

	locals.ElasticsearchKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.ElasticsearchKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(OpElasticsearchKubeEndpoint, pulumi.String(locals.ElasticsearchKubeServiceFqdn))

	locals.ElasticsearchKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.ElasticsearchKubeServiceName, vars.ElasticsearchPort, vars.ElasticsearchPort)

	//export kube-port-forward command
	ctx.Export(OpElasticsearchPortForwardCommand, pulumi.String(locals.ElasticsearchKubePortForwardCommand))

	locals.KibanaKubeServiceName = fmt.Sprintf("%s-kb-http", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(OpKibanaService, pulumi.String(locals.KibanaKubeServiceName))

	locals.KibanaKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KibanaKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(OpKibanaKubeEndpoint, pulumi.String(locals.KibanaKubeServiceFqdn))

	locals.KibanaKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.KibanaKubeServiceName, vars.KibanaPort, vars.KibanaPort)

	//export kube-port-forward command
	ctx.Export(OpKibanaPortForwardCommand, pulumi.String(locals.KibanaKubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled ||
		target.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.ElasticsearchIngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.ElasticsearchIngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.KibanaIngressExternalHostname = fmt.Sprintf("%s-kb.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.KibanaIngressInternalHostname = fmt.Sprintf("%s-kb-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.ElasticsearchIngressExternalHostname,
		locals.ElasticsearchIngressInternalHostname,
		locals.KibanaIngressExternalHostname,
		locals.KibanaIngressInternalHostname,
	}

	locals.IngressCertClusterIssuerName = target.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = locals.Namespace

	//export ingress hostnames
	ctx.Export(OpElasticsearchExternalHostname, pulumi.String(locals.ElasticsearchIngressExternalHostname))
	ctx.Export(OpElasticsearchInternalHostname, pulumi.String(locals.ElasticsearchIngressInternalHostname))
	ctx.Export(OpKibanaExternalHostname, pulumi.String(locals.KibanaIngressExternalHostname))
	ctx.Export(OpKibanaInternalHostname, pulumi.String(locals.KibanaIngressInternalHostname))

	return locals
}
