package module

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	v1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
)

// Resources is the main entry point for setting up ECS cluster, task definition, and service.
func Resources(ctx *pulumi.Context, stackInput *v1.EcsServiceStackInput) error {
	return nil
}
