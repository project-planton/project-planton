package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func apis(ctx *pulumi.Context, locals *Locals, createdProject *organizations.Project) error {
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
	return nil
}
