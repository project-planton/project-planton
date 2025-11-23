package module

import (
	kubernetesprometheusv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesprometheus/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// locals contains computed values used throughout the Prometheus Kubernetes deployment
type locals struct {
	// resourceId is a stable identifier derived from metadata
	resourceId string
	// namespace is the Kubernetes namespace where Prometheus will be deployed
	namespace string
	// serviceName is the name of the Prometheus Kubernetes service
	serviceName string
	// labels are the common labels applied to all resources
	labels map[string]string
}

// newLocals creates and initializes the locals struct with computed values
func newLocals(stackInput *kubernetesprometheusv1.KubernetesPrometheusStackInput) *locals {
	metadata := stackInput.Target.Metadata

	// Derive resource ID from metadata
	resourceId := metadata.Id
	if resourceId == "" {
		resourceId = metadata.Name
	}

	// Build common labels
	labels := map[string]string{
		"resource":      "true",
		"resource_id":   resourceId,
		"resource_kind": "prometheus_kubernetes",
	}

	// Add organization label if present
	if metadata.Org != "" {
		labels["organization"] = metadata.Org
	}

	// Add environment label if present
	if metadata.Env != "" {
		labels["environment"] = metadata.Env
	}

	// Get namespace from spec
	namespace := stackInput.Target.Spec.Namespace.GetValue()

	return &locals{
		resourceId:  resourceId,
		namespace:   namespace,
		serviceName: metadata.Name + "-prometheus",
		labels:      labels,
	}
}

// exports creates the Pulumi stack outputs
func (l *locals) exports(ctx *pulumi.Context) error {
	// Export namespace
	ctx.Export(Namespace, pulumi.String(l.namespace))

	// Export service name
	ctx.Export(Service, pulumi.String(l.serviceName))

	// Export port forward command
	portForwardCmd := "kubectl port-forward -n " + l.namespace + " service/" + l.serviceName + " 9090:9090"
	ctx.Export(PortForwardCommand, pulumi.String(portForwardCmd))

	// Export Kubernetes endpoint (FQDN)
	kubeEndpoint := l.serviceName + "." + l.namespace + ".svc.cluster.local"
	ctx.Export(KubeEndpoint, pulumi.String(kubeEndpoint))

	// Export external hostname (if ingress is enabled)
	// This would be set based on ingress configuration
	ctx.Export(ExternalHostname, pulumi.String(""))

	// Export internal hostname (if ingress is enabled)
	// This would be set based on ingress configuration
	ctx.Export(InternalHostname, pulumi.String(""))

	return nil
}
