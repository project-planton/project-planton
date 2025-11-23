package module

import (
	"fmt"
	"strconv"
	"strings"

	kuberneteselasticsearchv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteselasticsearch/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	ElasticsearchIngressExternalHostname string
	ElasticsearchKubePortForwardCommand  string
	ElasticsearchKubeServiceFqdn         string
	ElasticsearchKubeServiceName         string
	Namespace                            string
	KubernetesElasticsearch              *kuberneteselasticsearchv1.KubernetesElasticsearch
	KibanaIngressExternalHostname        string
	KibanaKubePortForwardCommand         string
	KibanaKubeServiceFqdn                string
	KibanaKubeServiceName                string
	IngressHostnames                     []string
	IngressCertClusterIssuerName         string
	IngressCertSecretName                string
	Labels                               map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kuberneteselasticsearchv1.KubernetesElasticsearchStackInput) *Locals {
	locals := &Locals{}

	locals.KubernetesElasticsearch = stackInput.Target

	target := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesElasticsearch.String(),
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
	// 3. Override with spec.namespace if provided
	// 4. Override with stackInput if provided

	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if target.Spec.Namespace != nil && target.Spec.Namespace.GetValue() != "" {
		locals.Namespace = target.Spec.Namespace.GetValue()
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

	// Elasticsearch ingress
	if target.Spec.Elasticsearch.Ingress != nil &&
		target.Spec.Elasticsearch.Ingress.Enabled &&
		target.Spec.Elasticsearch.Ingress.Hostname != "" {

		locals.ElasticsearchIngressExternalHostname = target.Spec.Elasticsearch.Ingress.Hostname
		ctx.Export(OpElasticsearchExternalHostname, pulumi.String(locals.ElasticsearchIngressExternalHostname))
		locals.IngressHostnames = append(locals.IngressHostnames, locals.ElasticsearchIngressExternalHostname)
	}

	// Kibana ingress
	if target.Spec.Kibana != nil && target.Spec.Kibana.Enabled &&
		target.Spec.Kibana.Ingress != nil &&
		target.Spec.Kibana.Ingress.Enabled &&
		target.Spec.Kibana.Ingress.Hostname != "" {

		locals.KibanaIngressExternalHostname = target.Spec.Kibana.Ingress.Hostname
		ctx.Export(OpKibanaExternalHostname, pulumi.String(locals.KibanaIngressExternalHostname))
		locals.IngressHostnames = append(locals.IngressHostnames, locals.KibanaIngressExternalHostname)
	}

	// Set certificate issuer and secret name if any ingress is enabled
	if len(locals.IngressHostnames) > 0 {
		// Use first available hostname to extract domain for cert issuer
		firstHostname := locals.ElasticsearchIngressExternalHostname
		if firstHostname == "" {
			firstHostname = locals.KibanaIngressExternalHostname
		}
		// Extract domain from hostname for cert issuer
		parts := strings.Split(firstHostname, ".")
		if len(parts) > 1 {
			locals.IngressCertClusterIssuerName = strings.Join(parts[1:], ".")
		}
		locals.IngressCertSecretName = locals.Namespace
	}

	return locals
}
