package provider

import "context"

// DependencyOutputsKey is the context key that carries, while a scenario is
// verified, the outputs of every prerequisite the runner deployed in front of
// it. A verifier whose proof runs through what the resource composes with --
// a worker pool proved by calling the control plane's API, at the address,
// server name and trust bundle that control plane published -- reads those
// values here instead of restating them, so the proof follows whatever the
// prerequisite actually deployed.
type DependencyOutputsKey struct{}

// DependencyOutputs holds each deployed prerequisite's outputs, keyed by the
// catalog kind's name (KubernetesAgentSubstrate) and then by the instance's
// manifest name. Kind names, not the runner's kind enum, keep this package
// free of the runner's types.
type DependencyOutputs map[string]map[string]map[string]interface{}

// DependencyOutputsFrom returns the prerequisite outputs the runner put on
// ctx, or nil when the scenario deployed no prerequisites.
func DependencyOutputsFrom(ctx context.Context) DependencyOutputs {
	outputs, _ := ctx.Value(DependencyOutputsKey{}).(DependencyOutputs)
	return outputs
}
