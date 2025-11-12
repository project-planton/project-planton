package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpproject/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources provisions a GCP Project, generating its project ID with a 3-char suffix.
func Resources(ctx *pulumi.Context, stackInput *gcpprojectv1.GcpProjectStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	// Create gcp provider using credentials from the input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	createdProject, err := project(ctx, locals, gcpProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create GCP project")
	}

	if err := apis(ctx, locals, createdProject, gcpProvider); err != nil {
		return errors.Wrap(err, "failed to enabled apis for GCP project")
	}

	if err := iam(ctx, locals, createdProject, gcpProvider); err != nil {
		return errors.Wrap(err, "failed to create IAM bindings for GCP project")
	}

	// Export outputs
	ctx.Export(OpProjectId, createdProject.ProjectId)
	ctx.Export(OpProjectNumber, createdProject.Number)
	ctx.Export(OpProjectName, createdProject.Name)

	return nil
}
