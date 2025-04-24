package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	gcpprojectv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpproject/v1"
)

// Resources provisions a GCP Project and related settings.
func Resources(ctx *pulumi.Context, stackInput *gcpprojectv1.GcpProjectStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	// Create GCP project
	createdProject, err := organizations.NewProject(ctx,
		locals.GcpProject.Metadata.Name,
		&organizations.ProjectArgs{
			Name:              pulumi.String(locals.GcpProject.Metadata.Name),
			ProjectId:         pulumi.String(locals.GcpProject.Spec.ProjectId),
			BillingAccount:    pulumi.String(locals.GcpProject.Spec.BillingAccountId),
			Labels:            pulumi.ToStringMap(locals.GcpLabels),
			FolderId:          pulumi.String(locals.GcpProject.Spec.FolderId),
			OrgId:             pulumi.String(locals.GcpProject.Spec.OrgId),
			AutoCreateNetwork: pulumi.Bool(!locals.GcpProject.Spec.DisableDefaultNetwork),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create GCP project")
	}

	// Enable specified APIs
	for _, api := range locals.GcpProject.Spec.EnabledApis {
		serviceName := fmt.Sprintf("%s-enable-%s", locals.GcpProject.Metadata.Name, api)

		_, srvErr := projects.NewService(ctx, serviceName, &projects.ServiceArgs{
			Project:                  createdProject.ProjectId,
			Service:                  pulumi.String(api),
			DisableDependentServices: pulumi.Bool(true),
			DisableOnDestroy:         pulumi.Bool(false),
		}, pulumi.DependsOn([]pulumi.Resource{createdProject}))
		if srvErr != nil {
			return errors.Wrapf(srvErr, "failed to enable API %s", api)
		}
	}

	// Optionally assign an owner IAM member
	if locals.GcpProject.Spec.OwnerMember != "" {
		_, iamErr := projects.NewIAMMember(ctx,
			fmt.Sprintf("%s-owner-binding", locals.GcpProject.Metadata.Name),
			&projects.IAMMemberArgs{
				Project: createdProject.ProjectId,
				Role:    pulumi.String("roles/owner"),
				Member:  pulumi.String(locals.GcpProject.Spec.OwnerMember),
			},
			pulumi.DependsOn([]pulumi.Resource{createdProject}),
		)
		if iamErr != nil {
			return errors.Wrap(iamErr, "failed to create owner IAM binding")
		}
	}

	// Export outputs
	ctx.Export(OpProjectId, createdProject.ProjectId)
	ctx.Export(OpProjectNumber, createdProject.Number)
	ctx.Export(OpProjectName, createdProject.Name)

	return nil
}
