package module

import (
	"fmt"
	openfgakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/openfgakubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/openfgakubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	OpenfgaKubernetes            *openfgakubernetesv1.OpenfgaKubernetes
	Namespace                    string
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	IngressInternalHostname      string
	IngressExternalHostname      string
	IngressHostnames             []string
	KubeServiceFqdn              string
	KubeServiceName              string
	KubePortForwardCommand       string
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *openfgakubernetesv1.OpenfgaKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the local variable to make it available across the project
	locals.OpenfgaKubernetes = stackInput.Target

	openfgaKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Environment:  stackInput.Target.Metadata.Env.Id,
		kuberneteslabelkeys.Organization: stackInput.Target.Metadata.Org,
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   stackInput.Target.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: "openfga_kubernetes",
	}

	//decide on the namespace
	locals.Namespace = openfgaKubernetes.Metadata.Id

	locals.KubeServiceName = openfgaKubernetes.Metadata.Name

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local",
		openfgaKubernetes.Metadata.Name, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, openfgaKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if openfgaKubernetes.Spec.Ingress == nil ||
		!openfgaKubernetes.Spec.Ingress.IsEnabled ||
		openfgaKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s",
		openfgaKubernetes.Metadata.Id, openfgaKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", openfgaKubernetes.Metadata.Id,
		openfgaKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	//export ingress hostnames
	//ctx.Export(outputs.IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	//ctx.Export(outputs.IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	//note: a ClusterIssuer resource should have already exist on the kubernetes-cluster.
	//this is typically taken care of by the kubernetes cluster administrator.
	//if the kubernetes-cluster is created using Planton Cloud, then the cluster-issuer name will be
	//same as the ingress-domain-name as long as the same ingress-domain-name is added to the list of
	//ingress-domain-names for the GkeCluster/EksCluster/AksCluster spec.
	locals.IngressCertClusterIssuerName = openfgaKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = openfgaKubernetes.Metadata.Id

	return locals
}
