package module

import (
	"github.com/pkg/errors"
	gcpvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpvpc/v1"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// vpc creates the VPC Network and enables the Compute API when necessary.
// It closely mirrors a Terraform-style module: enable provider‑level services first,
// then declare the core resource, and finally export stack outputs.
func vpc(ctx *pulumi.Context, locals *Locals, gcpProvider *gcp.Provider) (*compute.Network, error) {
	// 1. Enable the Compute Engine API for the target project.
	createdComputeService, err := projects.NewService(ctx,
		"compute-api",
		&projects.ServiceArgs{
			Project:                  pulumi.String(locals.GcpVpc.Spec.ProjectId.GetValue()),
			Service:                  pulumi.String("compute.googleapis.com"),
			DisableDependentServices: pulumi.BoolPtr(true),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to enable compute api")
	}

	// 2. Build the args for the network resource.
	networkArgs := &compute.NetworkArgs{
		AutoCreateSubnetworks: pulumi.BoolPtr(locals.GcpVpc.Spec.AutoCreateSubnetworks),
		Name:                  pulumi.String(locals.GcpVpc.Metadata.Name),
		Project:               pulumi.String(locals.GcpVpc.Spec.ProjectId.GetValue()),
	}

	// Map the routing mode enum to the expected GCP value if explicitly set.
	if locals.GcpVpc.Spec.GetRoutingMode() != gcpvpcv1.GcpVpcRoutingMode_REGIONAL {
		// GLOBAL is the only alternative at present.
		networkArgs.RoutingMode = pulumi.StringPtr("GLOBAL")
	}

	// 3. Create the VPC network.
	createdNetwork, err := compute.NewNetwork(ctx,
		"vpc",
		networkArgs,
		pulumi.Provider(gcpProvider),
		pulumi.DependsOn([]pulumi.Resource{createdComputeService}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vpc network")
	}

	// 4. Export stack outputs.
	ctx.Export(OpNetworkSelfLink, createdNetwork.SelfLink)

	return createdNetwork, nil
}
