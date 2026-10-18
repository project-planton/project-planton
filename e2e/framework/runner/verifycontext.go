package runner

import (
	"context"

	"github.com/plantonhq/planton/e2e/framework/provider"
)

// verifyContext is the context every verify phase of a scenario runs under:
// the manifest being verified, and the outputs of the prerequisites deployed
// in front of it (provider.DependencyOutputsKey), so a verifier can prove the
// resource through what it composes with. Every place the runner moves to a
// new manifest (reference resolution, a setup script's expansion, a second
// act) builds its context here, so no step drops either value; a step that
// only swaps the manifest (an expected-failure upgrade) layers the new path
// on a context built here, keeping the outputs.
func verifyContext(ctx context.Context, manifestPath string, depOutputs DependencyOutputs) context.Context {
	ctx = context.WithValue(ctx, provider.ManifestPathKey{}, manifestPath)
	if len(depOutputs) == 0 {
		return ctx
	}
	byName := make(provider.DependencyOutputs, len(depOutputs))
	for kind, instances := range depOutputs {
		byName[kind.String()] = instances
	}
	return context.WithValue(ctx, provider.DependencyOutputsKey{}, byName)
}
