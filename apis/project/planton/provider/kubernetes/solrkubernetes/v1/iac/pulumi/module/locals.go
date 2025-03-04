package module

import (
	"fmt"
	solrkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/solrkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/solrkubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	IngressExternalHostname      string
	IngressHostnames             []string
	IngressInternalHostname      string
	KubePortForwardCommand       string
	KubeServiceFqdn              string
	KubeServiceName              string
	Namespace                    string
	SolrKubernetes               *solrkubernetesv1.SolrKubernetes
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *solrkubernetesv1.SolrKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the locals variable to make it available across the project
	locals.SolrKubernetes = stackInput.Target

	solrKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: solrKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "solr_kubernetes",
	}

	if solrKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = solrKubernetes.Metadata.Id
	}

	if solrKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = solrKubernetes.Metadata.Org
	}

	if solrKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = solrKubernetes.Metadata.Env
	}

	locals.Namespace = solrKubernetes.Metadata.Name

	if solrKubernetes.Metadata.Labels != nil &&
		solrKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = solrKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.KubeServiceName = fmt.Sprintf("%s-solrcloud-common", solrKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf(
		"%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, solrKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if solrKubernetes.Spec.Ingress == nil ||
		!solrKubernetes.Spec.Ingress.IsEnabled ||
		solrKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		solrKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		solrKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	//export ingress hostnames
	ctx.Export(outputs.ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.InternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = solrKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = fmt.Sprintf("cert-%s", locals.Namespace)

	return locals
}
