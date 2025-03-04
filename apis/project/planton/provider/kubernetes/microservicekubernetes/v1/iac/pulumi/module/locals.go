package module

import (
	"fmt"
	microservicekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/microservicekubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/microservicekubernetes/v1/iac/pulumi/module/outputs"
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
	MicroserviceKubernetes       *microservicekubernetesv1.MicroserviceKubernetes
	ImagePullSecretData          map[string]string
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *microservicekubernetesv1.MicroserviceKubernetesStackInput) (*Locals, error) {
	locals := &Locals{}

	locals.MicroserviceKubernetes = stackInput.Target

	microserviceKubernetes := stackInput.Target

	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: microserviceKubernetes.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: "microservice_kubernetes",
	}

	if microserviceKubernetes.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = microserviceKubernetes.Metadata.Id
	}

	if microserviceKubernetes.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = microserviceKubernetes.Metadata.Org
	}

	if microserviceKubernetes.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = microserviceKubernetes.Metadata.Env
	}

	locals.Namespace = microserviceKubernetes.Metadata.Name

	if microserviceKubernetes.Metadata.Labels != nil &&
		microserviceKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey] != "" {
		locals.Namespace = microserviceKubernetes.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]
	}

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))

	if stackInput.DockerConfigJson != "" {
		locals.ImagePullSecretData = map[string]string{".dockerconfigjson": stackInput.DockerConfigJson}
	}

	locals.KubeServiceName = microserviceKubernetes.Spec.Version

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, locals.KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if microserviceKubernetes.Spec.Ingress == nil ||
		!microserviceKubernetes.Spec.Ingress.IsEnabled ||
		microserviceKubernetes.Spec.Ingress.DnsDomain == "" {
		return locals, nil
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", microserviceKubernetes.Metadata.Id,
		microserviceKubernetes.Spec.Ingress.DnsDomain)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", microserviceKubernetes.Metadata.Id,
		microserviceKubernetes.Spec.Ingress.DnsDomain)

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
	locals.IngressCertClusterIssuerName = microserviceKubernetes.Spec.Ingress.DnsDomain

	locals.IngressCertSecretName = microserviceKubernetes.Metadata.Id

	return locals, nil
}
