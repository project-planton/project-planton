package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// database provisions the Cloudflare D1 database and exports its outputs.
func database(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.D1Database, error) {

	// 1.  Build arguments directly from proto fields—no extra structs.
	d1Args := &cloudflare.D1DatabaseArgs{
		AccountId: pulumi.String(locals.CloudflareD1Database.Spec.AccountId),
		Name:      pulumi.String(locals.CloudflareD1Database.Spec.DatabaseName),
	}

	// 2.  Create the resource.
	createdD1Database, err := cloudflare.NewD1Database(
		ctx,
		"database",
		d1Args,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare d1 database")
	}

	// 3.  Export stack outputs.
	ctx.Export(OpDatabaseId, createdD1Database.ID())
	ctx.Export(OpDatabaseName, createdD1Database.Name)

	// NOTE: Pulumi’s Cloudflare provider (v6.4.1) does not yet expose a connection
	// string for D1. We export an empty value to satisfy the Project Planton schema.
	ctx.Export(OpConnectionString, pulumi.String(""))

	return createdD1Database, nil
}
