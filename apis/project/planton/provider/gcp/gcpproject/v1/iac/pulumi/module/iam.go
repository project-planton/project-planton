package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func iam(ctx *pulumi.Context, locals *Locals, createdProject *organizations.Project) error {
	// Create IAM roles and bindings
	// Optionally assign an owner IAM member
	if locals.GcpProject.Spec.OwnerMember != "" {
		_, iamErr := projects.NewIAMMember(ctx,
			fmt.Sprintf("%s-owner-binding", locals.GcpProject.Metadata.Name),
			&projects.IAMMemberArgs{
				Project: createdProject.ProjectId,
				Role:    pulumi.String("roles/owner"),
				Member:  pulumi.String("user:" + locals.GcpProject.Spec.OwnerMember),
			},
			pulumi.DependsOn([]pulumi.Resource{createdProject}),
		)
		if iamErr != nil {
			return errors.Wrap(iamErr, "failed to create owner IAM binding")
		}
	}
	return nil
}
