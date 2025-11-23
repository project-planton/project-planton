package module

import (
	kubernetessolroperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetessolroperator/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// locals computes derived configuration values from the stack input
type locals struct {
	// namespace is the Kubernetes namespace for operator deployment
	namespace string

	// labels are common labels applied to all resources
	labels pulumi.StringMap

	// operatorName is the name of the operator deployment
	operatorName string

	// chartVersion is the Helm chart version to install
	chartVersion string
}

// newLocals creates computed values from stack input
func newLocals(stackInput *kubernetessolroperatorv1.KubernetesSolrOperatorStackInput) *locals {
	// Use metadata.name or default to "solr-operator"
	operatorName := "solr-operator"
	if stackInput.Target != nil && stackInput.Target.Metadata != nil && stackInput.Target.Metadata.Name != "" {
		operatorName = stackInput.Target.Metadata.Name
	}

	// Build common labels
	labels := pulumi.StringMap{
		"app.kubernetes.io/name":       pulumi.String("solr-operator"),
		"app.kubernetes.io/managed-by": pulumi.String("project-planton"),
		"planton.cloud/resource-kind":  pulumi.String("kubernetes-solr-operator"),
	}

	// Add metadata labels if provided
	if stackInput.Target != nil && stackInput.Target.Metadata != nil {
		if stackInput.Target.Metadata.Name != "" {
			labels["planton.cloud/resource-id"] = pulumi.String(stackInput.Target.Metadata.Name)
		}
		if stackInput.Target.Metadata.Org != "" {
			labels["planton.cloud/organization"] = pulumi.String(stackInput.Target.Metadata.Org)
		}
		if stackInput.Target.Metadata.Env != "" {
			labels["planton.cloud/environment"] = pulumi.String(stackInput.Target.Metadata.Env)
		}
	}

	// get namespace from spec
	namespace := vars.Namespace
	if stackInput.Target != nil && stackInput.Target.Spec != nil && stackInput.Target.Spec.Namespace != nil {
		namespace = stackInput.Target.Spec.Namespace.GetValue()
	}

	return &locals{
		namespace:    namespace,
		labels:       labels,
		operatorName: operatorName,
		chartVersion: vars.DefaultStableVersion,
	}
}
