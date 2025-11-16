package module

import (
	"fmt"
	"strconv"

	kubernetesargocdv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesargocd/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	KubernetesArgocd             *kubernetesargocdv1.KubernetesArgocd
	Namespace                    string
	ServiceName                  string
	KubeServiceFqdn              string
	KubePortForwardCommand       string
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	IngressInternalHostname      string
	IngressExternalHostname      string
	IngressHostnames             []string
	Labels                       map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kubernetesargocdv1.KubernetesArgocdStackInput) *Locals {
	locals := &Locals{}

	// Assign value for the local variable to make it available across the project
	locals.KubernetesArgocd = stackInput.Target

	target := stackInput.Target

	// Build labels
	locals.Labels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_KubernetesArgocd.String(),
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

	// Namespace determination
	// Priority order:
	// 1. Default: "argo-" + metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided

	resourceId := target.Metadata.Name
	if target.Metadata.Id != "" {
		resourceId = target.Metadata.Id
	}

	locals.Namespace = fmt.Sprintf("argo-%s", resourceId)

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.KubernetesNamespace != "" {
		locals.Namespace = stackInput.KubernetesNamespace
	}

	// Export namespace
	ctx.Export(Namespace, pulumi.String(locals.Namespace))

	// Service name follows Helm chart naming: <release-name>-argocd-server
	locals.ServiceName = fmt.Sprintf("%s-argocd-server", resourceId)

	// Export service name
	ctx.Export(Service, pulumi.String(locals.ServiceName))

	// Kubernetes service FQDN for internal cluster access
	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local",
		locals.ServiceName, locals.Namespace)

	// Export kubernetes endpoint
	ctx.Export(KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	// Port-forward command for local access
	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:80",
		locals.Namespace, locals.ServiceName)

	// Export port-forward command
	ctx.Export(PortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	// Handle ingress configuration
	if target.Spec.Ingress == nil ||
		!target.Spec.Ingress.Enabled ||
		target.Spec.Ingress.DnsDomain == "" {
		// Export empty hostnames if ingress is not enabled
		ctx.Export(ExternalHostname, pulumi.String(""))
		ctx.Export(InternalHostname, pulumi.String(""))
		return locals
	}

	// External hostname (public ingress)
	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	// Internal hostname (private ingress)
	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace,
		target.Spec.Ingress.DnsDomain)

	locals.IngressHostnames = []string{
		locals.IngressExternalHostname,
		locals.IngressInternalHostname,
	}

	// Export ingress hostnames
	ctx.Export(ExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(InternalHostname, pulumi.String(locals.IngressInternalHostname))

	// Certificate configuration
	// Note: a ClusterIssuer resource should already exist on the kubernetes-cluster
	// The cluster-issuer name will be same as the ingress-domain-name
	locals.IngressCertClusterIssuerName = target.Spec.Ingress.DnsDomain
	locals.IngressCertSecretName = resourceId

	return locals
}
