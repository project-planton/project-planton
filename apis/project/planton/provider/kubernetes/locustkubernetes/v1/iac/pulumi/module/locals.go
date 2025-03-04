package module

import (
	"fmt"
	locustkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/locustkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/locustkubernetes/v1/iac/pulumi/module/outputs"
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
	LocustKubernetes             *locustkubernetesv1.LocustKubernetes
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *locustkubernetesv1.LocustKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the locals variable to make it available across the project
	locals.LocustKubernetes = stackInput.Target

	locustKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: locustKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "locust_kubernetes",
	}

	if locustKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = locustKubernetes.Metadata.Id
	}

	if locustKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = locustKubernetes.Metadata.Org
	}

	if locustKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = locustKubernetes.Metadata.Env
	}

	locals.Namespace = locustKubernetes.Metadata.Name

	if locustKubernetes.Metadata.Labels != nil &&
		locustKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = locustKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	locals.KubeServiceName = locustKubernetes.Metadata.Name

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locustKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if locustKubernetes.Spec.Ingress == nil ||
		!locustKubernetes.Spec.Ingress.IsEnabled ||
		locustKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locustKubernetes.Metadata.Id,
		locustKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locustKubernetes.Metadata.Id,
		locustKubernetes.Spec.Ingress.DnsDomain)

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
	locals.IngressCertClusterIssuerName = locustKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = locustKubernetes.Metadata.Id

	return locals
}
