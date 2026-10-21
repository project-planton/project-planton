package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/plantonhq/planton/pkg/downloads/modulemanifest"
)

// fixtureRepo builds a tiny git repository shaped like the catalog: one
// released provider, one internal provider, tracked and untracked files.
func fixtureRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	write := func(path, content string) {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	git := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@t", "GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@t")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
	write("LICENSE", "license\n")
	write("NOTICE", "notice\n")
	write("catalog/aws/awsvpc/iac/tf/main.tf", "# vpc\n")
	write("catalog/aws/awsvpc/iac/tf/README.md", "# docs\n")
	write("catalog/aws/awsvpc/iac/tf/hack/manifest.yaml", "kind: x\n")
	write("catalog/aws/awsvpc/iac/pulumi/main.go", "package main\n")
	write("catalog/aws/awsvpc/iac/pulumi/README.md", "# docs\n")
	write("catalog/_test/testkind/iac/tf/main.tf", "# internal\n")
	git("init", "-q")
	git("add", ".")
	git("commit", "-q", "-m", "fixture")
	// Left behind by a local plan: untracked, so never part of a module.
	write("catalog/aws/awsvpc/iac/tf/terraform.tfplan", "plan")
	write("catalog/aws/awsvpc/iac/tf/.terraform.lock.hcl", "stray")
	return root
}

func keys(m map[string][]byte) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func TestReleasedModulesSkipsInternalProviders(t *testing.T) {
	root := fixtureRepo(t)
	modules, err := ReleasedModules(root)
	if err != nil {
		t.Fatal(err)
	}
	var dirs []string
	for _, m := range modules {
		dirs = append(dirs, m.Dir+" "+string(m.Engine))
	}
	want := []string{"catalog/aws/awsvpc/iac/pulumi pulumi", "catalog/aws/awsvpc/iac/tf terraform"}
	if strings.Join(dirs, ",") != strings.Join(want, ",") {
		t.Fatalf("released modules %v, want %v", dirs, want)
	}
}

func TestExecutedFilesAreTrackedAndUndocumented(t *testing.T) {
	root := fixtureRepo(t)
	entries, err := executedFiles(root, "catalog/aws/awsvpc/iac/tf")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.Join(keys(entries), ","), "LICENSE,NOTICE,hack/manifest.yaml,main.tf"; got != want {
		t.Fatalf("module.zip ships %s, want %s (no README, nothing untracked)", got, want)
	}
}

func TestSourceFilesKeepDocumentation(t *testing.T) {
	root := fixtureRepo(t)
	entries, err := sourceFiles(root, "catalog/aws/awsvpc/iac/pulumi")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.Join(keys(entries), ","), "LICENSE,NOTICE,README.md,main.go"; got != want {
		t.Fatalf("source.zip ships %s, want %s", got, want)
	}
}

func TestEditingAReadmeLeavesModuleZipUnchanged(t *testing.T) {
	root := fixtureRepo(t)
	dir := t.TempDir()
	first := filepath.Join(dir, "first.zip")
	if err := PackModule(root, "catalog/aws/awsvpc/iac/tf", first); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "catalog/aws/awsvpc/iac/tf/README.md"), []byte("# rewritten\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	second := filepath.Join(dir, "second.zip")
	if err := PackModule(root, "catalog/aws/awsvpc/iac/tf", second); err != nil {
		t.Fatal(err)
	}
	a, _ := os.ReadFile(first)
	b, _ := os.ReadFile(second)
	if string(a) != string(b) {
		t.Fatal("a README edit changed the module archive a runner executes")
	}
}

func TestBuildEnvFixesEveryBuildVariable(t *testing.T) {
	env := buildEnv([]string{"GOFLAGS=-race", "GOAMD64=v3", "GOWORK=/x/go.work", "HOME=/h", "CGO_ENABLED=1"}, "go1.26.6", "windows", "amd64")
	got := map[string]string{}
	for _, kv := range env {
		name, value, _ := strings.Cut(kv, "=")
		if _, dup := got[name]; dup {
			t.Fatalf("%s is set twice", name)
		}
		got[name] = value
	}
	want := map[string]string{
		"GOFLAGS": "-mod=readonly", "GOAMD64": "v1", "GOARM64": "v8.0", "GOWORK": "off", "CGO_ENABLED": "0",
		"GOOS": "windows", "GOARCH": "amd64", "GOTOOLCHAIN": "go1.26.6", "GOEXPERIMENT": "", "HOME": "/h",
	}
	for name, value := range want {
		if got[name] != value {
			t.Errorf("%s = %q, want %q", name, got[name], value)
		}
	}
}

func TestGoModToolchain(t *testing.T) {
	cases := map[string]string{
		"module x\n\ngo 1.26.6\n":                       "go1.26.6",
		"module x\n\ngo 1.26.0\n\ntoolchain go1.26.6\n": "go1.26.6",
	}
	for gomod, want := range cases {
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte(gomod), 0o644); err != nil {
			t.Fatal(err)
		}
		got, err := goModToolchain(root)
		if err != nil || got != want {
			t.Errorf("goModToolchain(%q) = %q, %v; want %q", gomod, got, err, want)
		}
	}
}

func TestBuildManifestRefusesAnIncompleteEngine(t *testing.T) {
	sum := strings.Repeat("a", 64)
	tf := Module{Dir: "catalog/aws/awsvpc/iac/tf", Engine: modulemanifest.Terraform}
	pu := Module{Dir: "catalog/aws/awsvpc/iac/pulumi", Engine: modulemanifest.Pulumi}
	released := []Module{tf, pu}
	// What each engine's legs publish, and nothing of the other engine's:
	// each manifest job downloads only its own lane's parts.
	complete := map[modulemanifest.Engine]func() map[string]map[string]string{
		modulemanifest.Terraform: func() map[string]map[string]string {
			return map[string]map[string]string{tf.Dir: {modulemanifest.TerraformArtifact: sum}}
		},
		modulemanifest.Pulumi: func() map[string]map[string]string {
			p := map[string]map[string]string{pu.Dir: {}}
			for _, platform := range modulemanifest.PulumiPlatforms {
				p[pu.Dir][modulemanifest.PulumiArtifact(platform)] = sum
			}
			return p
		},
	}

	for engine, own := range map[modulemanifest.Engine]Module{modulemanifest.Terraform: tf, modulemanifest.Pulumi: pu} {
		data, err := BuildManifest("v1.2.3", engine, released, complete[engine]())
		if err != nil {
			t.Fatalf("a complete %s lane was refused while the other engine published nothing: %v", engine, err)
		}
		m, err := modulemanifest.Parse(data)
		if err != nil {
			t.Fatalf("the %s manifest written does not parse: %v", engine, err)
		}
		if m.Engine != engine || m.ReleaseTag != "v1.2.3" || len(m.Modules) != 1 || m.Modules[own.Dir].Digest == "" {
			t.Fatalf("the %s manifest %+v does not describe exactly that engine's modules", engine, m)
		}
	}

	cases := []struct {
		name   string
		engine modulemanifest.Engine
		edit   func(p map[string]map[string]string)
		want   string
	}{
		{"a platform missing", modulemanifest.Pulumi, func(p map[string]map[string]string) { delete(p[pu.Dir], "windows_amd64.exe.gz") }, "windows_amd64.exe.gz was not published"},
		{"a module no leg recorded", modulemanifest.Terraform, func(p map[string]map[string]string) { delete(p, tf.Dir) }, "module.zip was not published"},
		{"an unexpected artifact", modulemanifest.Terraform, func(p map[string]map[string]string) { p[tf.Dir]["source.zip"] = sum }, "not an artifact this module ships"},
		{"a module the release does not ship", modulemanifest.Terraform, func(p map[string]map[string]string) {
			p["catalog/aws/gone/iac/tf"] = map[string]string{modulemanifest.TerraformArtifact: sum}
		}, "ships no such terraform module"},
		{"the other engine's part", modulemanifest.Terraform, func(p map[string]map[string]string) {
			p[pu.Dir] = map[string]string{modulemanifest.PulumiArtifact("linux_amd64"): sum}
		}, "ships no such terraform module"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := complete[tc.engine]()
			tc.edit(p)
			if _, err := BuildManifest("v1.2.3", tc.engine, released, p); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("BuildManifest() error = %v, want one mentioning %q", err, tc.want)
			}
		})
	}
}

func TestReadPartsMergesLegsAndRefusesTwoChecksums(t *testing.T) {
	dir := t.TempDir()
	a, b := strings.Repeat("a", 64), strings.Repeat("b", 64)
	// download-artifact nests each leg's part in its own folder.
	if err := os.MkdirAll(filepath.Join(dir, "leg-1"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "leg-1", "part.tsv"), []byte("catalog/aws/x/iac/tf\tmodule.zip\t"+a+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "leg-2.tsv"), []byte("catalog/aws/x/iac/pulumi\tlinux_amd64.gz\t"+b+"\n\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := ReadParts(dir)
	if err != nil {
		t.Fatal(err)
	}
	if got["catalog/aws/x/iac/tf"]["module.zip"] != a || got["catalog/aws/x/iac/pulumi"]["linux_amd64.gz"] != b {
		t.Fatalf("parts read as %v", got)
	}

	if err := os.WriteFile(filepath.Join(dir, "leg-3.tsv"), []byte("catalog/aws/x/iac/tf\tmodule.zip\t"+b+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadParts(dir); err == nil || !strings.Contains(err.Error(), "two checksums") {
		t.Fatalf("ReadParts() error = %v, want the two-checksums refusal", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "leg-3.tsv"), []byte("not a part line\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadParts(dir); err == nil {
		t.Fatal("a malformed part line was accepted")
	}
}
