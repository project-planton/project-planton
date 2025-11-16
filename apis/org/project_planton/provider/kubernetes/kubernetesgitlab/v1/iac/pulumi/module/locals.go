package module

import (
	kubernetesgitlabv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesgitlab/v1"
)

// getNamespace returns the namespace for the GitLab deployment
// Uses the resource ID as the namespace name
func getNamespace(metadata *kubernetesgitlabv1.KubernetesGitlabMetadata) string {
	if metadata.Id != "" {
		return metadata.Id
	}
	return metadata.Name
}

// getLabels returns the standard labels for GitLab resources
func getLabels(metadata *kubernetesgitlabv1.KubernetesGitlabMetadata) map[string]string {
	labels := map[string]string{
		"resource":      "true",
		"resource_kind": "kubernetes_gitlab",
	}

	if metadata.Id != "" {
		labels["resource_id"] = metadata.Id
	} else {
		labels["resource_id"] = metadata.Name
	}

	if metadata.Org != "" {
		labels["organization"] = metadata.Org
	}

	if metadata.Env != "" {
		labels["environment"] = metadata.Env
	}

	return labels
}
