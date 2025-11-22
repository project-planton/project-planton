package module

import (
	kuberneteskeycloakv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kuberneteskeycloak/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	// Keycloak namespace
	Namespace string

	// Resource labels for all resources
	Labels map[string]string

	// Ingress configuration
	IngressEnabled bool
	Hostname       string

	// Keycloak service configuration
	ServiceName string
	ServicePort int

	// Stack outputs
	PortForwardCommand string
	KubeEndpoint       string
	ExternalHostname   string
	InternalHostname   string
}

func initializeLocals(ctx *pulumi.Context, stackInput *kuberneteskeycloakv1.KubernetesKeycloakStackInput) *Locals {
	locals := &Locals{}

	// Determine namespace
	locals.Namespace = "keycloak-" + stackInput.Target.Metadata.Name

	// Set up labels
	locals.Labels = map[string]string{
		"app":      "keycloak",
		"resource": stackInput.Target.Metadata.Name,
	}

	if stackInput.Target.Metadata.Env != "" {
		locals.Labels["env"] = stackInput.Target.Metadata.Env
	}

	if stackInput.Target.Metadata.Org != "" {
		locals.Labels["org"] = stackInput.Target.Metadata.Org
	}

	// Ingress configuration
	if stackInput.Target.Spec != nil && stackInput.Target.Spec.Ingress != nil {
		locals.IngressEnabled = stackInput.Target.Spec.Ingress.Enabled
		locals.Hostname = stackInput.Target.Spec.Ingress.Hostname
	}

	// Service configuration
	locals.ServiceName = "keycloak-" + stackInput.Target.Metadata.Name
	locals.ServicePort = 8080

	// Stack outputs
	locals.PortForwardCommand = "kubectl port-forward -n " + locals.Namespace + " svc/" + locals.ServiceName + " 8080:8080"
	locals.KubeEndpoint = locals.ServiceName + "." + locals.Namespace + ".svc.cluster.local:" + "8080"

	if locals.IngressEnabled && locals.Hostname != "" {
		locals.ExternalHostname = "https://" + locals.Hostname
		locals.InternalHostname = "https://internal-" + locals.Hostname
	}

	return locals
}
