package module

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// appPlatformService provisions a single‑service App and exports its key outputs.
func appPlatformService(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.App, error) {

	serviceArgs := &digitalocean.AppSpecServiceArgs{
		InstanceCount:    pulumi.Int(int(locals.DigitalOceanAppPlatformService.Spec.InstanceCount)),
		InstanceSizeSlug: pulumi.String(strings.ReplaceAll(locals.DigitalOceanAppPlatformService.Spec.InstanceSizeSlug.String(), "_", "-")),
		Name:             pulumi.String(locals.DigitalOceanAppPlatformService.Spec.ServiceName),
	}

	// --------------------------------------------
	// 2. Source (git or image).
	// --------------------------------------------
	if locals.DigitalOceanAppPlatformService.Spec.GetGitSource() != nil {
		g := locals.DigitalOceanAppPlatformService.Spec.GetGitSource()
		serviceArgs.Git = &digitalocean.AppSpecServiceGitArgs{
			Branch: pulumi.String(g.Branch),
		}
	} else if locals.DigitalOceanAppPlatformService.Spec.GetImageSource() != nil {
		i := locals.DigitalOceanAppPlatformService.Spec.GetImageSource()
		serviceArgs.Image = &digitalocean.AppSpecServiceImageArgs{
			Registry:   pulumi.String(i.Registry.GetValue()), // handles foreign‑key / direct value
			Repository: pulumi.String(i.Repository),
			Tag:        pulumi.String(i.Tag),
		}
	}

	// --------------------------------------------
	// 4. Environment variables.
	// --------------------------------------------
	if len(locals.DigitalOceanAppPlatformService.Spec.Env) > 0 {
		var envArray digitalocean.AppSpecServiceEnvArray
		for k, v := range locals.DigitalOceanAppPlatformService.Spec.Env {
			envArray = append(envArray, digitalocean.AppSpecServiceEnvArgs{
				Key:   pulumi.String(k),
				Value: pulumi.String(v),
				Scope: pulumi.String("RUN_AND_BUILD_TIME"),
			})
		}
		serviceArgs.Envs = envArray
	}

	// --------------------------------------------
	// 6. Construct the full App spec.
	// --------------------------------------------
	appSpecArgs := &digitalocean.AppSpecArgs{
		Name:     pulumi.String(locals.DigitalOceanAppPlatformService.Spec.ServiceName),
		Region:   pulumi.String(locals.DigitalOceanAppPlatformService.Spec.Region.String()),
		Services: digitalocean.AppSpecServiceArray{serviceArgs},
	}

	// --------------------------------------------
	// 7. Create the App.
	// --------------------------------------------
	createdApp, err := digitalocean.NewApp(
		ctx,
		"app",
		&digitalocean.AppArgs{
			Spec: appSpecArgs,
		},
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean app")
	}

	// --------------------------------------------
	// 8. Export stack outputs.
	// --------------------------------------------
	ctx.Export(OpAppId, createdApp.ID())
	ctx.Export(OpLiveUrl, createdApp.LiveUrl)

	return createdApp, nil
}
