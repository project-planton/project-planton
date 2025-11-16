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
	if stackInput.Metadata != nil && stackInput.Metadata.Name != "" {
		operatorName = stackInput.Metadata.Name
	}

	// Build common labels
	labels := pulumi.StringMap{
		"app.kubernetes.io/name":       pulumi.String("solr-operator"),
		"app.kubernetes.io/managed-by": pulumi.String("project-planton"),
		"planton.cloud/resource-kind":  pulumi.String("kubernetes-solr-operator"),
	}

	// Add metadata labels if provided
	if stackInput.Metadata != nil {
		if stackInput.Metadata.Name != "" {
			labels["planton.cloud/resource-id"] = pulumi.String(stackInput.Metadata.Name)
		}
		if stackInput.Metadata.Org != "" {
			labels["planton.cloud/organization"] = pulumi.String(stackInput.Metadata.Org)
		}
		if stackInput.Metadata.Env != "" {
			labels["planton.cloud/environment"] = pulumi.String(stackInput.Metadata.Env)
		}
	}

	return &locals{
		namespace:    vars.Namespace,
		labels:       labels,
		operatorName: operatorName,
		chartVersion: vars.DefaultStableVersion,
	}
}

