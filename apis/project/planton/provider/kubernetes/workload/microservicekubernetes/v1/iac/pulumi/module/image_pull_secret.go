package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// imagePullSecret creates a Kubernetes secret for pulling private container images.
//
// The docker config JSON can be provided in two ways (checked in priority order):
// 1. stackInput.DockerConfigJson - Takes precedence (used by Planton Cloud)
//   - If present, label is completely ignored
//
// 2. metadata.labels["kubernetes.planton.io/docker-config-json-file"] - File path (for open-source users)
//   - Only checked if stackInput.DockerConfigJson is empty
//   - Path can be relative, absolute, or use ~/ for home directory
//   - File must exist and contain valid JSON
//
// Returns nil if no docker config is configured (e.g., using GKE Workload Identity).
// Returns error if file path is specified in label but file cannot be read or is invalid.
func imagePullSecret(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) (*kubernetescorev1.Secret, error) {

	// If no image pull secret data is configured, return nil
	if locals.ImagePullSecretData == nil {
		return nil, nil
	}

	// Create image pull secret resource
	createdImagePullSecret, err := kubernetescorev1.NewSecret(ctx,
		"image-pull-secret",
		&kubernetescorev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("image-pull-secret"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Type:       pulumi.String("kubernetes.io/dockerconfigjson"),
			StringData: pulumi.ToStringMap(locals.ImagePullSecretData),
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create image pull secret")
	}

	return createdImagePullSecret, nil
}
