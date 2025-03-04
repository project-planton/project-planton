package module

import (
	"fmt"
	elasticsearchkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/elasticsearchkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/elasticsearchkubernetes/v1/iac/pulumi/module/outputs"
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
	//assign value for the local variable to make it available across the module.
	locals.ElasticsearchKubernetes = stackInput.Target

	elasticsearchKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: elasticsearchKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "elasticsearch_kubernetes",
	}

	if elasticsearchKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = elasticsearchKubernetes.Metadata.Id
	}

	if elasticsearchKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = elasticsearchKubernetes.Metadata.Org
	}

	if elasticsearchKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = elasticsearchKubernetes.Metadata.Env
	}

	locals.Namespace = elasticsearchKubernetes.Metadata.Name

	if elasticsearchKubernetes.Metadata.Labels != nil &&
		elasticsearchKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = elasticsearchKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	ctx.Export(outputs.ElasticsearchUsername, pulumi.String("elastic"))
	ctx.Export(outputs.ElasticsearchPasswordSecretName, pulumi.Sprintf("%s-es-elastic-user", elasticsearchKubernetes.Metadata.Name))
	ctx.Export(outputs.ElasticsearchPasswordSecretKey, pulumi.String("elastic"))

	locals.ElasticsearchKubeServiceName = fmt.Sprintf("%s-es-http", elasticsearchKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.ElasticsearchService, pulumi.String(locals.ElasticsearchKubeServiceName))

	locals.ElasticsearchKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.ElasticsearchKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.ElasticsearchKubeEndpoint, pulumi.String(locals.ElasticsearchKubeServiceFqdn))

	locals.ElasticsearchKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.ElasticsearchKubeServiceName, vars.ElasticsearchPort, vars.ElasticsearchPort)

	//export kube-port-forward command
	ctx.Export(outputs.ElasticsearchPortForwardCommand, pulumi.String(locals.ElasticsearchKubePortForwardCommand))

	locals.KibanaKubeServiceName = fmt.Sprintf("%s-kb-http", elasticsearchKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.KibanaService, pulumi.String(locals.KibanaKubeServiceName))

	locals.KibanaKubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KibanaKubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KibanaKubeEndpoint, pulumi.String(locals.KibanaKubeServiceFqdn))

	locals.KibanaKubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.KibanaKubeServiceName, vars.KibanaPort, vars.KibanaPort)

	//export kube-port-forward command
	ctx.Export(outputs.KibanaPortForwardCommand, pulumi.String(locals.KibanaKubePortForwardCommand))

	if elasticsearchKubernetes.Spec.Ingress == nil ||
		!elasticsearchKubernetes.Spec.Ingress.IsEnabled ||
		elasticsearchKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.ElasticsearchIngressExternalHostname = fmt.Sprintf("%s.%s", elasticsearchKubernetes.Metadata.Id,
		elasticsearchKubernetes.Spec.Ingress.DnsDomain)

	locals.ElasticsearchIngressInternalHostname = fmt.Sprintf("%s-internal.%s", elasticsearchKubernetes.Metadata.Id,
		elasticsearchKubernetes.Spec.Ingress.DnsDomain)

	locals.KibanaIngressExternalHostname = fmt.Sprintf("%s-kb.%s", elasticsearchKubernetes.Metadata.Id,
		elasticsearchKubernetes.Spec.Ingress.DnsDomain)

	locals.KibanaIngressInternalHostname = fmt.Sprintf("%s-kb-internal.%s", elasticsearchKubernetes.Metadata.Id,
		elasticsearchKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.ElasticsearchIngressExternalHostname,
		locals.ElasticsearchIngressInternalHostname,
		locals.KibanaIngressExternalHostname,
		locals.KibanaIngressInternalHostname,
	}

	locals.IngressCertClusterIssuerName = elasticsearchKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = elasticsearchKubernetes.Metadata.Id

	//export ingress hostnames
	ctx.Export(outputs.ElasticsearchExternalHostname, pulumi.String(locals.ElasticsearchIngressExternalHostname))
	ctx.Export(outputs.ElasticsearchInternalHostname, pulumi.String(locals.ElasticsearchIngressInternalHostname))
	ctx.Export(outputs.KibanaExternalHostname, pulumi.String(locals.KibanaIngressExternalHostname))
	ctx.Export(outputs.KibanaInternalHostname, pulumi.String(locals.KibanaIngressInternalHostname))

	return locals
}
