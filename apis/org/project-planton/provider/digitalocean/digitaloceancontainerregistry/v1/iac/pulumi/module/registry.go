package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	digitaloceancontainerregistryv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/digitalocean/digitaloceancontainerregistry/v1"
)

// registry provisions the DigitalOcean Container Registry and exports outputs.
func registry(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.ContainerRegistry, error) {

	// 1. Translate the subscription tier enum to the expected slug.
	var tierSlug string
	switch locals.DigitalOceanContainerRegistry.Spec.SubscriptionTier {
	case digitaloceancontainerregistryv1.DigitalOceanContainerRegistryTier_STARTER:
		tierSlug = "starter"
	case digitaloceancontainerregistryv1.DigitalOceanContainerRegistryTier_BASIC:
		tierSlug = "basic"
	case digitaloceancontainerregistryv1.DigitalOceanContainerRegistryTier_PROFESSIONAL:
		tierSlug = "professional"
	default:
		tierSlug = "starter"
	}

	// 2. Build the resource arguments from proto fields.
	registryArgs := &digitalocean.ContainerRegistryArgs{
		Name:                 pulumi.String(locals.DigitalOceanContainerRegistry.Spec.Name),
		SubscriptionTierSlug: pulumi.String(tierSlug),
		Region:               pulumi.StringPtr(locals.DigitalOceanContainerRegistry.Spec.Region.String()),
	}

	// 3. Create the Container Registry.
	createdRegistry, err := digitalocean.NewContainerRegistry(
		ctx,
		"registry",
		registryArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean container registry")
	}

	// 4. Warn if garbage collection flag is set (provider limitation).
	if locals.DigitalOceanContainerRegistry.Spec.GarbageCollectionEnabled {
		ctx.Log.Warn(
			"garbage_collection_enabled is set but currently not supported by the DigitalOcean provider; ignoring.",
			&pulumi.LogArgs{},
		)
	}

	// 5. Export stack outputs.
	ctx.Export(OpRegistryName, createdRegistry.Name)
	ctx.Export(OpRegion, createdRegistry.Region)
	ctx.Export(OpServerUrl, pulumi.Sprintf("registry.digitalocean.com/%s", createdRegistry.Name))

	return createdRegistry, nil
}
