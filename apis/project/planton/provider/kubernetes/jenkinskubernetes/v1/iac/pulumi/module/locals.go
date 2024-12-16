package module

import (
	"fmt"
	jenkinskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/jenkinskubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/jenkinskubernetes/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	JenkinsKubernetes            *jenkinskubernetesv1.JenkinsKubernetes
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

func initializeLocals(ctx *pulumi.Context, stackInput *jenkinskubernetesv1.JenkinsKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the local variable to make it available across the project
	locals.JenkinsKubernetes = stackInput.Target

	jenkinsKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Environment:  stackInput.Target.Metadata.Env.Id,
		kuberneteslabelkeys.Organization: stackInput.Target.Metadata.Org,
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceId:   stackInput.Target.Metadata.Id,
		kuberneteslabelkeys.ResourceKind: "jenkins_kubernetes",
	}

	//decide on the namespace
	locals.Namespace = jenkinsKubernetes.Metadata.Id

	locals.KubeServiceName = jenkinsKubernetes.Metadata.Name

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local",
		jenkinsKubernetes.Metadata.Name, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, jenkinsKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if jenkinsKubernetes.Spec.Ingress == nil ||
		!jenkinsKubernetes.Spec.Ingress.IsEnabled ||
		jenkinsKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", jenkinsKubernetes.Metadata.Id,
		jenkinsKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", jenkinsKubernetes.Metadata.Id,
		jenkinsKubernetes.Spec.Ingress.DnsDomain)

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
	locals.IngressCertClusterIssuerName = jenkinsKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = jenkinsKubernetes.Metadata.Id

	return locals
}
