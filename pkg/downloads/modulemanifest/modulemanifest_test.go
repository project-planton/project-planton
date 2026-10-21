package modulemanifest

import (
	"strings"
	"testing"
)

func TestDigestIsOrderIndependentAndMovesWithAnyByte(t *testing.T) {
	a := map[string]string{"linux_amd64.gz": "aa", "darwin_arm64.gz": "bb"}
	b := map[string]string{"darwin_arm64.gz": "bb", "linux_amd64.gz": "aa"}
	if Digest(a) != Digest(b) {
		t.Fatal("the digest depends on map order")
	}
	moved := map[string]string{"linux_amd64.gz": "ab", "darwin_arm64.gz": "bb"}
	if Digest(a) == Digest(moved) {
		t.Fatal("an artifact's checksum changed but the digest did not")
	}
	renamed := map[string]string{"linux_arm64.gz": "aa", "darwin_arm64.gz": "bb"}
	if Digest(a) == Digest(renamed) {
		t.Fatal("an artifact's name changed but the digest did not")
	}
	// The digest is sha256 over "<name> <sha256>\n" lines in name order --
	// pinned here so every reader (the release, the platform, a runner)
	// computes the same value.
	if got, want := Digest(map[string]string{"module.zip": "00"}), "3747d0b99153eed607c17e0d17ee4de76b4e08f166e35951b202b9aaf6f6de2b"; got != want {
		t.Fatalf("Digest() = %s, want %s (sha256 of \"module.zip 00\\n\")", got, want)
	}
}

func TestParseProvesTheManifestWhole(t *testing.T) {
	artifacts := map[string]string{TerraformArtifact: "abc"}
	good := &Manifest{FormatVersion: FormatVersion, ReleaseTag: "v0.6.15", Engine: Terraform, Modules: map[string]Module{
		"catalog/aws/awsvpc/iac/tf": {Digest: Digest(artifacts), Artifacts: artifacts},
	}}
	data, err := Encode(good)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := Parse(data)
	if err != nil {
		t.Fatalf("a whole manifest was refused: %v", err)
	}
	if parsed.Modules["catalog/aws/awsvpc/iac/tf"].Digest != Digest(artifacts) {
		t.Fatal("the parsed digest is not the encoded one")
	}

	cases := []struct {
		name string
		edit func(m *Manifest)
		want string
	}{
		{"another format", func(m *Manifest) { m.FormatVersion = 3 }, "format 3"},
		{"no tag", func(m *Manifest) { m.ReleaseTag = "" }, "no release tag"},
		{"no engine", func(m *Manifest) { m.Engine = "" }, "neither pulumi nor terraform"},
		{"an unknown engine", func(m *Manifest) { m.Engine = "tf" }, "neither pulumi nor terraform"},
		{"another engine's module", func(m *Manifest) {
			pulumi := map[string]string{PulumiArtifact("linux_amd64"): "abc"}
			m.Modules["catalog/aws/awsvpc/iac/pulumi"] = Module{Digest: Digest(pulumi), Artifacts: pulumi}
		}, "not a terraform module directory"},
		{"a directory that is no module's", func(m *Manifest) {
			m.Modules["catalog/aws/awsvpc"] = good.Modules["catalog/aws/awsvpc/iac/tf"]
		}, "not a terraform module directory"},
		{"a digest that is not its artifacts'", func(m *Manifest) {
			m.Modules["catalog/aws/awsvpc/iac/tf"] = Module{Digest: "f00", Artifacts: artifacts}
		}, "digest to"},
		{"a module with no artifacts", func(m *Manifest) {
			m.Modules["catalog/aws/awsvpc/iac/tf"] = Module{Digest: Digest(nil)}
		}, "no artifacts"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Manifest{FormatVersion: good.FormatVersion, ReleaseTag: good.ReleaseTag, Engine: good.Engine, Modules: map[string]Module{}}
			for k, v := range good.Modules {
				m.Modules[k] = v
			}
			tc.edit(m)
			data, err := Encode(m)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := Parse(data); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Parse() error = %v, want one mentioning %q", err, tc.want)
			}
		})
	}
}

func TestEncodeIsDeterministic(t *testing.T) {
	m := &Manifest{FormatVersion: FormatVersion, ReleaseTag: "v1.0.0", Engine: Terraform, Modules: map[string]Module{}}
	for _, dir := range []string{"catalog/b/x/iac/tf", "catalog/a/y/iac/tf", "catalog/c/z/iac/tf"} {
		artifacts := map[string]string{TerraformArtifact: dir}
		m.Modules[dir] = Module{Digest: Digest(artifacts), Artifacts: artifacts}
	}
	first, _ := Encode(m)
	for i := 0; i < 20; i++ {
		again, _ := Encode(m)
		if string(again) != string(first) {
			t.Fatal("encoding the same manifest twice gave different bytes")
		}
	}
}

func TestEngineOfModuleDir(t *testing.T) {
	for dir, want := range map[string]Engine{
		"catalog/aws/awsvpc/iac/tf":     Terraform,
		"catalog/aws/awsvpc/iac/pulumi": Pulumi,
		"catalog/aws/awsvpc/iac":        "",
		"catalog/aws/awsvpc/tf":         "",
		"catalog/aws/awsvpc/iac/cdk":    "",
	} {
		got, ok := EngineOfModuleDir(dir)
		if got != want || ok != (want != "") {
			t.Errorf("EngineOfModuleDir(%q) = %q, %v; want %q", dir, got, ok, want)
		}
	}
	for _, e := range Engines {
		if got, _ := EngineOfModuleDir("catalog/aws/awsvpc/iac/" + e.ModuleDirName()); got != e {
			t.Errorf("%s's module directory name reads back as %q", e, got)
		}
	}
}

func TestPulumiArtifact(t *testing.T) {
	for platform, want := range map[string]string{
		"linux_amd64":   "linux_amd64.gz",
		"darwin_arm64":  "darwin_arm64.gz",
		"windows_amd64": "windows_amd64.exe.gz",
	} {
		if got := PulumiArtifact(platform); got != want {
			t.Errorf("PulumiArtifact(%q) = %q, want %q", platform, got, want)
		}
	}
}
