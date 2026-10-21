package aa_e2e

import (
	"strings"
	"testing"
)

// A profile cluster that needs a newer kind than the machine has must refuse
// before `kind create cluster` runs, with all three parts and the version it
// needs; a new-enough kind passes.
func TestKindVersionRefusal(t *testing.T) {
	cases := map[string]struct {
		output  string
		refused bool
	}{
		"older kind":       {output: "kind v0.24.0 go1.22.6 darwin/arm64\n", refused: true},
		"older minor":      {output: "kind v0.27.0 go1.23.4 linux/amd64", refused: true},
		"exact minimum":    {output: "kind v0.33.0 go1.25.1 darwin/arm64"},
		"newer kind":       {output: "kind v0.34.2 go1.26.0 linux/amd64"},
		"not the kind CLI": {output: "command not found", refused: true},
		"no version":       {output: "kind", refused: true},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := kindVersionRefusal(c.output, podCertificatesMinKindVersion, "planton-e2e-podcert")
			if !c.refused {
				if err != nil {
					t.Fatalf("kind %q was refused: %v", c.output, err)
				}
				return
			}
			if err == nil {
				t.Fatalf("kind %q was accepted for a cluster that needs %s", c.output, podCertificatesMinKindVersion)
			}
			for _, part := range []string{"observed: ", "meaning: ", "next step: ", podCertificatesMinKindVersion, "planton-e2e-podcert"} {
				if !strings.Contains(err.Error(), part) {
					t.Errorf("refusal lacks %q:\n%s", part, err)
				}
			}
		})
	}
}

// Only the profile that needs a newer kind carries a minimum; the default
// cluster builds on whatever kind is installed.
func TestProfileHarnessesCarryTheirKindMinimum(t *testing.T) {
	if got := NewPodCertificatesHarness("planton-e2e").minKindVersion; got != podCertificatesMinKindVersion {
		t.Errorf("the pod-certificates harness needs kind %q, want %q", got, podCertificatesMinKindVersion)
	}
	if got := NewHarness("planton-e2e").minKindVersion; got != "" {
		t.Errorf("the default harness needs kind %q, want no minimum", got)
	}
}
