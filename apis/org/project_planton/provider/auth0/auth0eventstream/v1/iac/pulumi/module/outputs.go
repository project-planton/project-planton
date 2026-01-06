package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-auth0/sdk/v3/go/auth0"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// exportOutputs exports the stack outputs for the Auth0 event stream
func exportOutputs(ctx *pulumi.Context, eventStream *auth0.EventStream, locals *Locals) error {
	// Export core identifiers
	ctx.Export("id", eventStream.ID())
	ctx.Export("name", eventStream.Name)
	ctx.Export("status", eventStream.Status)
	ctx.Export("destination_type", eventStream.DestinationType)
	ctx.Export("subscriptions", eventStream.Subscriptions)
	ctx.Export("created_at", eventStream.CreatedAt)
	ctx.Export("updated_at", eventStream.UpdatedAt)

	// Export EventBridge-specific output
	ctx.Export("eventbridge_configuration", eventStream.EventbridgeConfiguration)

	// Export custom outputs for easy reference
	ctx.Export("metadata_name", pulumi.String(locals.StreamName))

	if err := ctx.RegisterResourceOutputs(nil, pulumi.Map{
		"id":               eventStream.ID(),
		"name":             eventStream.Name,
		"status":           eventStream.Status,
		"destination_type": eventStream.DestinationType,
	}); err != nil {
		return errors.Wrap(err, "failed to register resource outputs")
	}

	return nil
}
