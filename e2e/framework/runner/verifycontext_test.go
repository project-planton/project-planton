package runner

import (
	"context"
	"testing"

	"github.com/plantonhq/planton/e2e/framework/provider"
	"github.com/plantonhq/planton/shared/catalogkind"
)

// The verify context carries the manifest under test and, keyed by kind name,
// every prerequisite's outputs: a verifier that proves a resource through its
// control plane reads the control plane's published endpoints from here.

func TestVerifyContext_CarriesManifestAndPrerequisiteOutputs(t *testing.T) {
	deps := DependencyOutputs{
		catalogkind.CatalogKind_KubernetesAgentSubstrate: {
			"e2e-substrate": {"api_address": "api.ate-system.svc:443"},
		},
	}
	ctx := verifyContext(context.Background(), "/tmp/pool.yaml", deps)

	if got, _ := ctx.Value(provider.ManifestPathKey{}).(string); got != "/tmp/pool.yaml" {
		t.Errorf("manifest path = %q, want /tmp/pool.yaml", got)
	}
	got := provider.DependencyOutputsFrom(ctx)["KubernetesAgentSubstrate"]["e2e-substrate"]["api_address"]
	if got != "api.ate-system.svc:443" {
		t.Errorf("KubernetesAgentSubstrate api_address = %v, want api.ate-system.svc:443", got)
	}
}

func TestVerifyContext_NoPrerequisitesMeansNoOutputs(t *testing.T) {
	ctx := verifyContext(context.Background(), "/tmp/ns.yaml", nil)
	if outputs := provider.DependencyOutputsFrom(ctx); outputs != nil {
		t.Errorf("outputs = %v, want nil when no prerequisite was deployed", outputs)
	}
}
