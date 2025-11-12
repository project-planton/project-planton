package module

import (
	"encoding/base64"

	"github.com/pkg/errors"
	gcpcertv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpcertmanagercert/v1"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for the gcp_cert_manager_cert Pulumi module.
// It prepares context, configures the GCP provider, and calls certManagerCert().
func Resources(ctx *pulumi.Context, stackInput *gcpcertv1.GcpCertManagerCertStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *gcp.Provider
	var err error
	gcpProviderConfig := stackInput.ProviderConfig

	if gcpProviderConfig == nil {
		provider, err = gcp.NewProvider(ctx, "classic-provider", &gcp.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default GCP provider")
		}
	} else {
		// Decode the base64 service account key
		decodedKey, decodeErr := base64.StdEncoding.DecodeString(gcpProviderConfig.ServiceAccountKeyBase64)
		if decodeErr != nil {
			return errors.Wrap(decodeErr, "failed to decode service account key")
		}

		provider, err = gcp.NewProvider(ctx, "classic-provider", &gcp.ProviderArgs{
			Credentials: pulumi.String(string(decodedKey)),
			Project:     pulumi.String(locals.GcpCertManagerCert.Spec.GcpProjectId),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create GCP provider with custom credentials")
		}
	}

	// Call the core logic for certificate creation and DNS validation setup.
	if err := certManagerCert(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create gcp cert manager cert resource")
	}

	return nil
}
