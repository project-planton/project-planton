package module

import (
	"fmt"
	"strconv"
	"strings"

	harborkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/harborkubernetes/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/kubernetes/kuberneteslabels"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	HarborKubernetes             *harborkubernetesv1.HarborKubernetes
	Namespace                    string
	KubernetesLabels             map[string]string
	CoreServiceName              string
	PortalServiceName            string
	RegistryServiceName          string
	JobserviceServiceName        string
	InternalCoreEndpoint         string
	InternalRegistryEndpoint     string
	KubePortForwardCommand       string
	IngressExternalHostname      string
	IngressHostnames             []string
	IngressCertClusterIssuerName string
	IngressCertSecretName        string
	RegistryExternalHostname     string
	NotaryExternalHostname       string
	DatabaseEndpoint             string
	RedisEndpoint                string
}

func initializeLocals(ctx *pulumi.Context, stackInput *harborkubernetesv1.HarborKubernetesStackInput) *Locals {
	locals := &Locals{}

	locals.HarborKubernetes = stackInput.GetTarget()
	target := stackInput.GetTarget()

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.ResourceName: target.Metadata.Name,
		kuberneteslabelkeys.ResourceKind: cloudresourcekind.CloudResourceKind_HarborKubernetes.String(),
	}

	if target.Metadata.Id != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}

	if target.Metadata.Org != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}

	if target.Metadata.Env != "" {
		locals.KubernetesLabels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// Priority order for namespace:
	// 1. Default: metadata.name
	// 2. Override with custom label if provided
	// 3. Override with stackInput if provided
	locals.Namespace = target.Metadata.Name

	if target.Metadata.Labels != nil &&
		target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey] != "" {
		locals.Namespace = target.Metadata.Labels[kuberneteslabels.NamespaceLabelKey]
	}

	if stackInput.GetKubernetesNamespace() != "" {
		locals.Namespace = stackInput.GetKubernetesNamespace()
	}

	//export namespace
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// Service names
	locals.CoreServiceName = fmt.Sprintf("%s-harbor-core", target.Metadata.Name)
	locals.PortalServiceName = fmt.Sprintf("%s-harbor-portal", target.Metadata.Name)
	locals.RegistryServiceName = fmt.Sprintf("%s-harbor-registry", target.Metadata.Name)
	locals.JobserviceServiceName = fmt.Sprintf("%s-harbor-jobservice", target.Metadata.Name)

	//export service names
	ctx.Export(OpCoreService, pulumi.String(locals.CoreServiceName))
	ctx.Export(OpPortalService, pulumi.String(locals.PortalServiceName))
	ctx.Export(OpRegistryService, pulumi.String(locals.RegistryServiceName))
	ctx.Export(OpJobserviceService, pulumi.String(locals.JobserviceServiceName))

	// Kubernetes FQDNs
	locals.InternalCoreEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.CoreServiceName, locals.Namespace, variables.HarborCorePort)
	locals.InternalRegistryEndpoint = fmt.Sprintf("%s.%s.svc.cluster.local:%d",
		locals.RegistryServiceName, locals.Namespace, variables.HarborRegistryPort)

	//export kubernetes endpoints
	ctx.Export(OpInternalCoreEndpoint, pulumi.String(locals.InternalCoreEndpoint))
	ctx.Export(OpInternalRegistryEndpoint, pulumi.String(locals.InternalRegistryEndpoint))

	// Port forward command
	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		locals.Namespace, locals.PortalServiceName, variables.HarborPortalPort, variables.HarborPortalPort)

	//export kube-port-forward command
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	// Admin credentials
	ctx.Export(OpAdminUsername, pulumi.String("admin"))
	ctx.Export(OpAdminPasswordSecretName, pulumi.String(fmt.Sprintf("%s-harbor-core", target.Metadata.Name)))
	ctx.Export(OpAdminPasswordSecretKey, pulumi.String("HARBOR_ADMIN_PASSWORD"))

	// Database outputs (only for self-managed)
	if target.Spec.Database != nil && !target.Spec.Database.IsExternal {
		locals.DatabaseEndpoint = fmt.Sprintf("%s-postgresql.%s.svc.cluster.local:%d",
			target.Metadata.Name, locals.Namespace, variables.PostgresPort)
		ctx.Export(OpDatabaseEndpoint, pulumi.String(locals.DatabaseEndpoint))
		ctx.Export(OpDatabaseUsername, pulumi.String("postgres"))
		ctx.Export(OpDatabasePasswordSecretName, pulumi.String(fmt.Sprintf("%s-postgresql", target.Metadata.Name)))
		ctx.Export(OpDatabasePasswordSecretKey, pulumi.String("postgres-password"))
	}

	// Redis outputs (only for self-managed)
	if target.Spec.Cache != nil && !target.Spec.Cache.IsExternal {
		locals.RedisEndpoint = fmt.Sprintf("%s-redis.%s.svc.cluster.local:%d",
			target.Metadata.Name, locals.Namespace, variables.RedisPort)
		ctx.Export(OpRedisEndpoint, pulumi.String(locals.RedisEndpoint))
		ctx.Export(OpRedisPasswordSecretName, pulumi.String(fmt.Sprintf("%s-redis", target.Metadata.Name)))
		ctx.Export(OpRedisPasswordSecretKey, pulumi.String("redis-password"))
	}

	// Ingress configuration for Core/Portal
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Core != nil &&
		target.Spec.Ingress.Core.Enabled &&
		target.Spec.Ingress.Core.Hostname != "" {
		locals.IngressExternalHostname = target.Spec.Ingress.Core.Hostname
		ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))

		// Registry external hostname is the same as core unless explicitly configured differently
		locals.RegistryExternalHostname = locals.IngressExternalHostname
		ctx.Export(OpRegistryExternalHostname, pulumi.String(locals.RegistryExternalHostname))

		locals.IngressHostnames = []string{
			locals.IngressExternalHostname,
		}

		// ClusterIssuer should already exist on the cluster
		// Extract domain from hostname for ClusterIssuer name
		hostnameParts := strings.Split(locals.IngressExternalHostname, ".")
		if len(hostnameParts) > 1 {
			locals.IngressCertClusterIssuerName = strings.Join(hostnameParts[1:], ".")
		}

		locals.IngressCertSecretName = fmt.Sprintf("cert-%s", locals.Namespace)
	}

	// Ingress configuration for Notary
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Notary != nil &&
		target.Spec.Ingress.Notary.Enabled &&
		target.Spec.Ingress.Notary.Hostname != "" {
		locals.NotaryExternalHostname = target.Spec.Ingress.Notary.Hostname
		ctx.Export(OpNotaryExternalHostname, pulumi.String(locals.NotaryExternalHostname))
	}

	return locals
}
