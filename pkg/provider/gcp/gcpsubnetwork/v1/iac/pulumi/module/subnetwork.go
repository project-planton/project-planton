package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// subnetworkProjectApis lists the minimal set of Google Cloud APIs that
// need to be enabled for subnet operations.  Add or remove entries here
// if future requirements change.
var subnetworkProjectApis = []string{
	"compute.googleapis.com",
}

// subnetwork sets up the required APIs (if necessary) and then provisions a
// custom‑mode subnet in an existing VPC.
//
// Inputs:
//   - ctx      Pulumi context
//   - locals   Helper bundle with spec + metadata + labels
//   - provider Pre‑configured GCP provider
//
// Returns:
//   - *compute.Subnetwork pointer for further composition / export
//   - error if something goes wrong
func subnetwork(ctx *pulumi.Context,
	locals *Locals,
	provider *gcp.Provider) (*compute.Subnetwork, error) {

	// --- (1) Enable APIs ----------------------------------------------------
	createdGoogleApiResources := make([]pulumi.Resource, 0)

	for _, api := range subnetworkProjectApis {
		createdProjectService, err := projects.NewService(ctx,
			"subnetwork-"+api,
			&projects.ServiceArgs{
				Project:                  pulumi.String(locals.GcpSubnetwork.Spec.ProjectId.GetValue()),
				DisableDependentServices: pulumi.BoolPtr(true),
				Service:                  pulumi.String(api),
			}, pulumi.Provider(provider))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to enable %s api", api)
		}
		createdGoogleApiResources = append(createdGoogleApiResources, createdProjectService)
	}

	// --- (2) Prepare secondary ranges input --------------------------------
	var secondaryRanges compute.SubnetworkSecondaryIpRangeArray
	for _, r := range locals.GcpSubnetwork.Spec.SecondaryIpRanges {
		secondaryRanges = append(secondaryRanges, &compute.SubnetworkSecondaryIpRangeArgs{
			RangeName:   pulumi.String(r.RangeName),
			IpCidrRange: pulumi.String(r.IpCidrRange),
		})
	}

	// --- (3) Create the subnet ---------------------------------------------
	createdSubnetwork, err := compute.NewSubnetwork(ctx,
		"subnetwork",
		&compute.SubnetworkArgs{
			Name:                  pulumi.String(locals.GcpSubnetwork.Metadata.Name),
			Project:               pulumi.StringPtr(locals.GcpSubnetwork.Spec.ProjectId.GetValue()),
			Region:                pulumi.String(locals.GcpSubnetwork.Spec.Region),
			Network:               pulumi.String(locals.GcpSubnetwork.Spec.VpcSelfLink.GetValue()),
			IpCidrRange:           pulumi.String(locals.GcpSubnetwork.Spec.IpCidrRange),
			PrivateIpGoogleAccess: pulumi.BoolPtr(locals.GcpSubnetwork.Spec.PrivateIpGoogleAccess),
			SecondaryIpRanges:     secondaryRanges,
		},
		pulumi.Provider(provider),
		pulumi.DependsOn(createdGoogleApiResources))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subnetwork")
	}

	// --- (4) Export outputs -------------------------------------------------
	ctx.Export(OpSubnetworkSelfLink, createdSubnetwork.SelfLink)
	ctx.Export(OpRegion, createdSubnetwork.Region)
	ctx.Export(OpIpCidrRange, createdSubnetwork.IpCidrRange)
	ctx.Export(OpSecondaryRanges, createdSubnetwork.SecondaryIpRanges)

	return createdSubnetwork, nil
}
