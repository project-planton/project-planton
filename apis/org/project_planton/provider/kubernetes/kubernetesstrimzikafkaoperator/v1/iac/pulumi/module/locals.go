package module

import (
	kubernetesstrimzikafkaoperatorv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesstrimzikafkaoperator/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// locals computes derived configuration values from the stack input
type locals struct {
	namespace    string
	labels       pulumi.StringMap
	operatorName string
	chartVersion string
}

// newLocals creates computed values from stack input
func newLocals(stackInput *kubernetesstrimzikafkaoperatorv1.KubernetesStrimziKafkaOperatorStackInput) *locals {
	operatorName := "strimzi-kafka-operator"
	if stackInput.Metadata != nil && stackInput.Metadata.Name != "" {
		operatorName = stackInput.Metadata.Name
	}

	labels := pulumi.StringMap{
		"app.kubernetes.io/name":       pulumi.String("strimzi-kafka-operator"),
		"app.kubernetes.io/managed-by": pulumi.String("project-planton"),
		"planton.cloud/resource-kind":  pulumi.String("kubernetes-strimzi-kafka-operator"),
	}

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

