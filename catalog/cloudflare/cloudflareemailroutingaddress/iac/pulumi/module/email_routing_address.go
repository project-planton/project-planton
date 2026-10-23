package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// emailRoutingAddress creates an account-scoped Email Routing destination
// address. Verification (clicking the emailed link) happens out-of-band.
func emailRoutingAddress(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.EmailRoutingAddress, error) {
	spec := locals.CloudflareEmailRoutingAddress.Spec

	args := &cloudflare.EmailRoutingAddressArgs{
		AccountId: pulumi.String(spec.AccountId),
		Email:     pulumi.String(spec.Email),
	}
	// Explicit verification-state override; empty leaves the state to the
	// normal emailed-link flow.
	if spec.Status != "" {
		args.Status = pulumi.String(spec.Status)
	}
	created, err := cloudflare.NewEmailRoutingAddress(
		ctx,
		"email-routing-address",
		args,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare email routing address")
	}

	ctx.Export(OpAddressId, created.ID())
	ctx.Export(OpEmail, created.Email)
	ctx.Export(OpVerified, created.Verified)
	ctx.Export(OpCreated, created.Created)

	return created, nil
}
