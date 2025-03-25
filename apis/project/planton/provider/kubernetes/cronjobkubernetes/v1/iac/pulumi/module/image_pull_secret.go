package module

import (
	"github.com/pkg/errors"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createImagePullSecret creates a Secret named "image-pull-secret" in the target namespace,
// using the Docker credentials from locals.ImagePullSecretData.
func createImagePullSecret(ctx *pulumi.Context, locals *Locals, createdNamespace *corev1.Namespace) error {
	_, err := corev1.NewSecret(ctx,
		"image-pull-secret",
		&corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("image-pull-secret"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Type:       pulumi.String("kubernetes.io/dockerconfigjson"),
			StringData: pulumi.ToStringMap(locals.ImagePullSecretData),
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create image pull secret")
	}
	return nil
}
