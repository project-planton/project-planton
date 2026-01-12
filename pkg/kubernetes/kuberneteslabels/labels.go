package kuberneteslabels

const (
	// DockerConfigJsonFileLabelKey specifies the file path containing docker config JSON for image pull secret
	DockerConfigJsonFileLabelKey = "kubernetes.project-planton.org/docker-config-json-file"

	// KubeContextLabelKey specifies the kubectl context to use for Kubernetes deployments
	KubeContextLabelKey = "kubernetes.project-planton.org/context"
)
