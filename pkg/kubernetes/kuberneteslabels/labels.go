package kuberneteslabels

const (
	// NamespaceLabelKey allows overriding the Kubernetes namespace for a resource
	NamespaceLabelKey = "kubernetes.project-planton.org/namespace"

	// DockerConfigJsonFileLabelKey specifies the file path containing docker config JSON for image pull secret
	DockerConfigJsonFileLabelKey = "kubernetes.project-planton.org/docker-config-json-file"
)
