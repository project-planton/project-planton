package providerlocks

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// The expectations are what OpenTofu itself wrote into a lock file for each
// constraint; a projection that differs would be rewritten by every init.
func TestNormalizeConstraintsWritesWhatTofuWrites(t *testing.T) {
	cases := map[string]string{
		"3.9.1":         "3.9.1",
		"= 3.9.1":       "3.9.1",
		">= 3.0":        ">= 3.0.0",
		">= 2.20":       ">= 2.20.0",
		">= 0.5.0":      ">= 0.5.0",
		"~> 3":          "~> 3.0",
		"~> 3.6":        "~> 3.6",
		"~> 3.6.0":      "~> 3.6.0",
		">= 3.6, < 4.0": ">= 3.6.0, < 4.0.0",
	}
	for raw, want := range cases {
		got, err := NormalizeConstraints(raw)
		if err != nil {
			t.Fatalf("%q: %v", raw, err)
		}
		if got != want {
			t.Errorf("NormalizeConstraints(%q) = %q, want %q", raw, got, want)
		}
	}
	if _, err := NormalizeConstraints(">= 1.0.0-beta1"); err == nil {
		t.Error("a pre-release constraint must be refused")
	}
}

func TestRenderReproducesALockTofuWrote(t *testing.T) {
	written, err := os.ReadFile("testdata/tofu-written.terraform.lock.hcl")
	if err != nil {
		t.Fatal(err)
	}
	lock, err := ParseLock(written, "tofu-written")
	if err != nil {
		t.Fatal(err)
	}
	if got := string(Render(lock)); got != string(written) {
		t.Fatalf("Render differs from tofu's own file:\n%s", got)
	}
}

func TestSatisfiesAndNewest(t *testing.T) {
	versions := []string{"2.20.0", "2.35.0", "2.38.0", "3.0.0", "2.39.0-beta1"}
	got, err := Newest(versions, []string{"~> 2.35", ">= 2.20.0", ">= 2.0.0"})
	if err != nil || got != "2.38.0" {
		t.Fatalf("Newest = %q, %v; want 2.38.0 (stable, inside every constraint)", got, err)
	}
	if ok, _ := Satisfies("3.0.0", []string{"~> 2.35"}); ok {
		t.Error("~> 2.35 must exclude 3.0.0")
	}
	if ok, _ := Satisfies("2.11.0", []string{"2.11.0", ""}); !ok {
		t.Error("an exact constraint must allow its own version, and an empty one anything")
	}
}

// writeModule writes a catalog module requiring the given providers.
func writeModule(t *testing.T, root, dir string, providers map[string]string) {
	t.Helper()
	var body strings.Builder
	body.WriteString("terraform {\n  required_providers {\n")
	for name, line := range providers {
		body.WriteString("    " + name + " = { " + line + " }\n")
	}
	body.WriteString("  }\n}\n")
	path := filepath.Join(root, dir)
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "provider.tf"), []byte(body.String()), 0o644); err != nil {
		t.Fatal(err)
	}
}

func canonicalFixture() Lock {
	return Lock{
		"registry.opentofu.org/hashicorp/aws": {
			Address: "registry.opentofu.org/hashicorp/aws", Version: "6.60.0",
			Hashes: []string{"zh:bb", "h1:aa"},
		},
		"registry.opentofu.org/hashicorp/random": {
			Address: "registry.opentofu.org/hashicorp/random", Version: "3.9.1",
			Hashes: []string{"h1:rr"},
		},
	}
}

func TestProjectionHoldsOnlyTheModulesProvidersWithItsConstraints(t *testing.T) {
	root := t.TempDir()
	writeModule(t, root, "catalog/aws/awsvpc/iac/tf", map[string]string{
		"aws": `source = "hashicorp/aws", version = "~> 6.58"`,
	})
	modules, err := Modules(root)
	if err != nil {
		t.Fatal(err)
	}
	projected, err := Project(canonicalFixture(), modules[0])
	if err != nil {
		t.Fatal(err)
	}
	want := `# This file is maintained automatically by "tofu init".
# Manual edits may be lost in future updates.

provider "registry.opentofu.org/hashicorp/aws" {
  version     = "6.60.0"
  constraints = "~> 6.58"
  hashes = [
    "h1:aa",
    "zh:bb",
  ]
}
`
	if string(projected) != want {
		t.Fatalf("projection:\n%s\nwant:\n%s", projected, want)
	}
}

func TestCheckFindsEveryKindOfStaleLock(t *testing.T) {
	root := t.TempDir()
	writeModule(t, root, "catalog/aws/awsvpc/iac/tf", map[string]string{
		"aws": `source = "hashicorp/aws", version = "~> 6.58"`,
	})
	writeModule(t, root, "catalog/aws/awsnolock/iac/tf", map[string]string{
		"aws": `source = "hashicorp/aws", version = "~> 6.58"`,
	})
	writeModule(t, root, "catalog/_test/generic/iac/tf", map[string]string{})
	modules, err := Modules(root)
	if err != nil {
		t.Fatal(err)
	}
	canonical := canonicalFixture()
	if err := os.MkdirAll(filepath.Join(root, "catalog/_providers"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, CanonicalLockPath), Render(canonical), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := WriteProjections(root, canonical, modules); err != nil {
		t.Fatal(err)
	}
	if problems, err := Check(root); err != nil || len(problems) != 1 || !strings.Contains(problems[0], "random is locked but no module requires it") {
		t.Fatalf("a freshly written catalog must only report the unused provider, got %v, %v", problems, err)
	}

	delete(canonical, "registry.opentofu.org/hashicorp/random")
	if err := os.WriteFile(filepath.Join(root, CanonicalLockPath), Render(canonical), 0o644); err != nil {
		t.Fatal(err)
	}
	if problems, _ := Check(root); len(problems) != 0 {
		t.Fatalf("current locks must pass, got %v", problems)
	}

	// A hand-edited lock, a missing lock, a lock where none belongs, and a
	// constraint the locked version no longer meets are each reported.
	_ = os.WriteFile(filepath.Join(root, "catalog/aws/awsvpc/iac/tf", LockFileName), []byte("# edited\n"), 0o644)
	_ = os.Remove(filepath.Join(root, "catalog/aws/awsnolock/iac/tf", LockFileName))
	_ = os.WriteFile(filepath.Join(root, "catalog/_test/generic/iac/tf", LockFileName), []byte("x"), 0o644)
	writeModule(t, root, "catalog/aws/awsnew/iac/tf", map[string]string{
		"aws": `source = "hashicorp/aws", version = "~> 7.0"`,
	})
	problems, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(problems, "\n")
	for _, want := range []string{
		"awsvpc/iac/tf's lock file is not its projection",
		"awsnolock/iac/tf has no lock file",
		"generic/iac/tf requires no provider but has a lock file",
		"aws is locked at 6.60.0, which a module's constraint excludes",
	} {
		if !strings.Contains(joined, want) {
			t.Errorf("Check missed %q in:\n%s", want, joined)
		}
	}
}

type fakeRegistry map[string][]Release

func (f fakeRegistry) Releases(address string) ([]Release, error) { return f[address], nil }

func TestPlanMovesKeepsLockedVersionsUnlessExcludedOrUpgraded(t *testing.T) {
	all := []string{"darwin_amd64", "darwin_arm64", "linux_amd64", "linux_arm64", "windows_amd64"}
	registry := fakeRegistry{
		"registry.opentofu.org/hashicorp/aws": {
			{Version: "6.60.0", Platforms: all}, {Version: "6.70.0", Platforms: all}, {Version: "7.0.0", Platforms: all},
		},
		"registry.opentofu.org/stripe/stripe": {
			{Version: "0.3.0", Platforms: []string{"darwin_arm64", "linux_amd64", "freebsd_amd64"}},
		},
	}
	modules := []Module{{Dir: "catalog/aws/awsvpc/iac/tf", Requirements: map[string]Requirement{
		"registry.opentofu.org/hashicorp/aws": {Address: "registry.opentofu.org/hashicorp/aws", Constraints: "~> 6.58"},
		"registry.opentofu.org/stripe/stripe": {Address: "registry.opentofu.org/stripe/stripe", Constraints: "0.3.0"},
	}}}
	canonical := Lock{"registry.opentofu.org/hashicorp/aws": {Address: "registry.opentofu.org/hashicorp/aws", Version: "6.60.0"}}

	moves, err := PlanMoves(modules, canonical, nil, registry)
	if err != nil {
		t.Fatal(err)
	}
	if len(moves) != 1 || moves[0].Address != "registry.opentofu.org/stripe/stripe" || moves[0].Reason != "newly required" {
		t.Fatalf("only the unlocked provider moves; aws stays at its lock: %+v", moves)
	}
	if !reflect.DeepEqual(moves[0].Platforms, []string{"darwin_arm64", "linux_amd64"}) {
		t.Errorf("a provider is locked for the target platforms it publishes, got %v", moves[0].Platforms)
	}

	canonical["registry.opentofu.org/stripe/stripe"] = LockedProvider{Version: "0.3.0"}
	moves, err = PlanMoves(modules, canonical, []string{"aws"}, registry)
	if err != nil {
		t.Fatal(err)
	}
	if len(moves) != 1 || moves[0].Version != "6.70.0" || moves[0].From != "6.60.0" {
		t.Fatalf("an upgrade moves aws to its newest allowed release, never past its constraint: %+v", moves)
	}

	if _, err := PlanMoves(modules, canonical, []string{"helm"}, registry); err == nil {
		t.Error("an upgrade naming a provider no module requires must be refused")
	}
}
