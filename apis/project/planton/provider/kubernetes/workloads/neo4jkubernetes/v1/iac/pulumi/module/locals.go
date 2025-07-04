package module

import (
	"fmt"
	"strconv"

	neo4jkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workloads/neo4jkubernetes/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/kuberneteslabelkeys"
	"github.com/project-planton/project-planton/pkg/overridelabels"
)

// Locals struct mirrors the "locals" concept from Terraform,
// providing a consolidated place for derived values.
type Locals struct {
	// The top-level resource from the user’s manifest (apiVersion, kind, metadata, spec).
	Neo4jKubernetes *neo4jkubernetesv1.Neo4JKubernetes

	// Standard labels used for all created resources.
	Labels map[string]string

	// The namespace in which everything will be deployed.
	Namespace string

	// Pod selector labels for identifying the pods that run Neo4j.
	Neo4jPodSelectorLabels map[string]string

	// Calculated hostnames for external or internal Ingress usage.
	IngressExternalHostname string
	IngressInternalHostname string

	// The name of the Kubernetes Service for the Neo4j instance.
	KubeServiceName string

	// The fully qualified domain name (in-cluster) for the Neo4j Service.
	KubeServiceFqdn string

	// A convenient port-forwarding command for local development.
	KubePortForwardCommand string
}

// initializeLocals populates Locals from the stack input and
// exports some fields to the Pulumi stack outputs.
func initializeLocals(
	ctx *pulumi.Context,
	stackInput *neo4jkubernetesv1.Neo4JKubernetesStackInput,
) *Locals {
	// Initialize the Locals struct.
	locals := &Locals{
		Neo4jKubernetes: stackInput.Target,
		Labels:          map[string]string{},
	}

	target := stackInput.Target

	// Basic labels that identify this resource in the cluster.
	locals.Labels[kuberneteslabelkeys.Resource] = strconv.FormatBool(true)
	locals.Labels[kuberneteslabelkeys.ResourceName] = target.Metadata.Name
	locals.Labels[kuberneteslabelkeys.ResourceKind] = cloudresourcekind.CloudResourceKind_Neo4jKubernetes.String()

	if target.Metadata.Id != "" {
		locals.Labels[kuberneteslabelkeys.ResourceId] = target.Metadata.Id
	}
	if target.Metadata.Org != "" {
		locals.Labels[kuberneteslabelkeys.Organization] = target.Metadata.Org
	}
	if target.Metadata.Env != "" {
		locals.Labels[kuberneteslabelkeys.Environment] = target.Metadata.Env
	}

	// Determine the namespace, defaulting to target.Metadata.Name,
	// but allowing override if there's a label for it.
	locals.Namespace = target.Metadata.Name
	if target.Metadata.Labels != nil {
		if overrideNS, exists := target.Metadata.Labels[overridelabels.KubernetesNamespaceLabelKey]; exists {
			locals.Namespace = overrideNS
		}
	}

	// Export the namespace in which Neo4j will reside.
	ctx.Export(OpNamespace, pulumi.String(locals.Namespace))

	// Define pod selector labels that Helm or raw K8s resources might use.
	locals.Neo4jPodSelectorLabels = map[string]string{
		"app.kubernetes.io/name":      "neo4j",
		"app.kubernetes.io/instance":  target.Metadata.Name,
		"app.kubernetes.io/component": "primary",
	}

	// Construct a service name and FQDN.
	locals.KubeServiceName = fmt.Sprintf("%s-neo4j", target.Metadata.Name)
	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)
	ctx.Export(OpService, pulumi.String(locals.KubeServiceName))
	ctx.Export(OpBoltUriKubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	// Provide a default port-forward command for local debugging (web interface is 7474).
	// For Bolt (7687), you could adapt this as needed.
	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 7474:7474",
		locals.Namespace, locals.KubeServiceName)
	ctx.Export(OpPortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	// If ingress is enabled, set up external and internal hostnames.
	if target.Spec.Ingress != nil &&
		target.Spec.Ingress.Enabled &&
		target.Spec.Ingress.DnsDomain != "" {
		locals.IngressExternalHostname = fmt.Sprintf("%s.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)
		locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", locals.Namespace, target.Spec.Ingress.DnsDomain)

		ctx.Export(OpExternalHostname, pulumi.String(locals.IngressExternalHostname))
		ctx.Export(OpInternalHostname, pulumi.String(locals.IngressInternalHostname))
	}

	return locals
}
