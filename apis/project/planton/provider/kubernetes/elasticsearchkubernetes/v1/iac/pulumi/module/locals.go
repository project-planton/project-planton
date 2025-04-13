package module

import (
	"fmt"
	elasticsearchkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/elasticsearchkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/elasticsearchkubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
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

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	ctx.Export(outputs.ElasticsearchUsername, pulumi.String("elastic"))
	ctx.Export(outputs.ElasticsearchPasswordSecretName, pulumi.Sprintf("%s-es-elastic-user", target.Metadata.Name))
	ctx.Export(outputs.ElasticsearchPasswordSecretKey, pulumi.String("elastic"))

	locals.ElasticsearchKubeServiceName = fmt.Sprintf("%s-es-http", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.ElasticsearchService, pulumi.String(locals.ElasticsearchKubeServiceName))

	locals.ElasticsearchKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.ElasticsearchKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.ElasticsearchKubeEndpoint, pulumi.String(locals.ElasticsearchKubeServiceFqdn))

	locals.ElasticsearchKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.ElasticsearchKubeServiceName, vars.ElasticsearchPort, vars.ElasticsearchPort)

	//export kube-port-forward command
	ctx.Export(outputs.ElasticsearchPortForwardCommand, pulumi.String(locals.ElasticsearchKubePortForwardCommand))

	locals.KibanaKubeServiceName = fmt.Sprintf("%s-kb-http", target.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.KibanaService, pulumi.String(locals.KibanaKubeServiceName))

	locals.KibanaKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KibanaKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KibanaKubeEndpoint, pulumi.String(locals.KibanaKubeServiceFqdn))

	locals.KibanaKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.KibanaKubeServiceName, vars.KibanaPort, vars.KibanaPort)

	//export kube-port-forward command
	ctx.Export(outputs.KibanaPortForwardCommand, pulumi.String(locals.KibanaKubePortForwardCommand))

	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.IsEnabled ||
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
	ctx.Export(outputs.ElasticsearchExternalHostname, pulumi.String(locals.ElasticsearchIngressExternalHostname))
	ctx.Export(outputs.ElasticsearchInternalHostname, pulumi.String(locals.ElasticsearchIngressInternalHostname))
	ctx.Export(outputs.KibanaExternalHostname, pulumi.String(locals.KibanaIngressExternalHostname))
	ctx.Export(outputs.KibanaInternalHostname, pulumi.String(locals.KibanaIngressInternalHostname))

	return locals
}
