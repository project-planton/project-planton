package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/plantonhq/planton/pkg/downloads/modulemanifest"
	"sigs.k8s.io/yaml"
)

// repoRoot is the repository root, or "" in a sandbox that carries only the
// package's declared inputs (Bazel); the workflow mirror gates run from
// `go test` at the repository root, like the other release-lane gates.
func repoRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".github", "workflows", "release.yaml")); err != nil {
		t.Skipf("repository tree not present at %s (sandboxed run)", root)
	}
	return root
}

type workflow struct {
	Jobs map[string]struct {
		Strategy struct {
			Matrix struct {
				Provider []string `json:"provider"`
				Platform []struct {
					Goos   string `json:"goos"`
					Goarch string `json:"goarch"`
				} `json:"platform"`
			} `json:"matrix"`
		} `json:"strategy"`
	} `json:"jobs"`
}

func readWorkflow(t *testing.T, root, name string) workflow {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, ".github", "workflows", name))
	if err != nil {
		t.Fatal(err)
	}
	var w workflow
	if err := yaml.Unmarshal(data, &w); err != nil {
		t.Fatalf("%s: %v", name, err)
	}
	return w
}

// TestPulumiMatrixBuildsEveryManifestPlatform holds the release's Pulumi
// matrix to the platforms the module manifest requires: a platform the matrix
// dropped would leave every Pulumi module incomplete, and the release would
// publish no manifest at all.
func TestPulumiMatrixBuildsEveryManifestPlatform(t *testing.T) {
	root := repoRoot(t)
	var built []string
	for _, p := range readWorkflow(t, root, "release.pulumi-modules.yaml").Jobs["build"].Strategy.Matrix.Platform {
		built = append(built, p.Goos+"_"+p.Goarch)
	}
	sort.Strings(built)
	want := append([]string(nil), modulemanifest.PulumiPlatforms...)
	sort.Strings(want)
	if strings.Join(built, ",") != strings.Join(want, ",") {
		t.Fatalf("release.pulumi-modules.yaml builds %v; the module manifest requires %v", built, want)
	}
}

// TestModuleMatricesCoverEveryProvider holds each module lane's provider
// matrix to the providers that ship modules of its engine, so a new provider
// fails here, on its pull request, rather than at release time.
func TestModuleMatricesCoverEveryProvider(t *testing.T) {
	root := repoRoot(t)
	modules, err := ReleasedModules(root)
	if err != nil {
		t.Fatal(err)
	}
	for engine, file := range map[Engine]string{Pulumi: "release.pulumi-modules.yaml", Terraform: "release.terraform-modules.yaml"} {
		matrix := map[string]bool{}
		for _, p := range readWorkflow(t, root, file).Jobs["build"].Strategy.Matrix.Provider {
			matrix[p] = true
		}
		for _, m := range modules {
			if m.Engine == engine && !matrix[m.Provider] {
				t.Errorf("%s ships %s modules, but %s's matrix has no %s leg", m.Provider, engine, file, m.Provider)
			}
		}
	}
}
