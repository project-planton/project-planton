package verify

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// The resource lanes are only as honest as the verifiers' reading of their
// scenarios: which behavior each lane proves, which claims and pools it
// expects. Each committed scenario is read here offline, so a scenario edit
// that changes what a lane asserts fails a test instead of quietly proving
// something else.

const agentSandboxResourceScenarios = "../../kubernetesagentsandbox/e2e/scenarios"

func TestAgentSandboxResourceVerifierReadsTheScenarios(t *testing.T) {
	cases := map[string]struct {
		suspended, suspendLane bool
		claims                 []string
	}{
		"minimal.yaml":               {},
		"full-surface.yaml":          {claims: []string{"workspace"}},
		"behavioral-suspend.yaml":    {suspendLane: true, claims: []string{"workspace"}},
		"behavioral-suspend.to.yaml": {suspended: true, suspendLane: true, claims: []string{"workspace"}},
	}
	if _, err := os.Stat(agentSandboxResourceScenarios); err != nil {
		t.Skip("catalog source tree not present (bazel sandbox); runs under go test")
	}
	files, err := filepath.Glob(filepath.Join(agentSandboxResourceScenarios, "*.yaml"))
	if err != nil || len(files) != len(cases) {
		t.Fatalf("every scenario needs a case here: found %d files, %d cases (%v)", len(files), len(cases), err)
	}
	for name, want := range cases {
		t.Run(name, func(t *testing.T) {
			v := newAgentSandboxVerifier(filepath.Join(agentSandboxResourceScenarios, name), "s", "ns")
			if v.Suspended != want.suspended || v.SuspendLane != want.suspendLane || !reflect.DeepEqual(v.ClaimTemplates, want.claims) {
				t.Errorf("suspended=%t suspendLane=%t claims=%v, want %t %t %v", v.Suspended, v.SuspendLane, v.ClaimTemplates, want.suspended, want.suspendLane, want.claims)
			}
		})
	}
}
