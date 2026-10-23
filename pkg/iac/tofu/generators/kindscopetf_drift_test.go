package generators

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// TestKindScopeTFDrift holds every module's kindscope.tf to the table in
// pkg/kubernetes/kindscope, so the OpenTofu engine decides a document's scope
// exactly as the Pulumi engine does.
func TestKindScopeTFDrift(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	repoRoot := filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "..")
	matches, err := filepath.Glob(filepath.Join(repoRoot, "catalog", "*", "*", "iac", "tf", KindScopeTFFileName))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) == 0 {
		t.Fatalf("no module carries %s -- the KubernetesManifest module reads it", KindScopeTFFileName)
	}
	canonical := KindScopeTF()
	regenerate := os.Getenv("PLANTON_REGEN_KINDSCOPE_TF") == "1"
	for _, path := range matches {
		rel, _ := filepath.Rel(repoRoot, path)
		committed, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(committed) == canonical {
			continue
		}
		if regenerate {
			if err := os.WriteFile(path, []byte(canonical), 0o644); err != nil {
				t.Fatal(err)
			}
			t.Logf("regenerated %s", rel)
			continue
		}
		t.Errorf("%s drifted from pkg/kubernetes/kindscope: regenerate with PLANTON_REGEN_KINDSCOPE_TF=1 go test ./pkg/iac/tofu/generators/ -run TestKindScopeTFDrift", rel)
	}
}
