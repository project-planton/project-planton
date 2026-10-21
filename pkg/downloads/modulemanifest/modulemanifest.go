// Package modulemanifest is the grammar of a release's module manifests: the
// documents that say which bytes every platform IaC module of a release is,
// and what each module's digest is.
//
// A release publishes one manifest per engine, beside that engine's
// artifacts: releases/<tag>/modules/terraform/manifest.json and
// releases/<tag>/modules/pulumi/manifest.json. Each engine's modules are
// built by their own release lane, and each manifest is published the moment
// its lane has uploaded every module of that engine -- so the presence of an
// engine's manifest means that engine's modules are complete, and one
// engine's slow build never holds back a platform that runs the other.
// A manifest is keyed by module directory, exactly as the catalog names it
// (catalog/aws/awsrdsinstance/iac/tf), and records, for every artifact a
// runner may execute, the sha256 of the bytes actually published under the
// tag.
//
// A module's digest is computed from those checksums (see Digest), so it is
// exact by construction: two releases carry the same digest for a module only
// when they published byte-identical artifacts for it. A platform that records
// the digest a resource was last applied with can therefore tell, release
// after release, whether that resource's module really changed -- and a
// runner can refuse to run a download whose bytes are not the ones the
// release published.
package modulemanifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strings"
)

// FileName is the manifest's file name under a release's
// modules/<engine>/ prefix.
const FileName = "manifest.json"

// FormatVersion is the manifest shape this package writes and reads. Format 2
// is one manifest per engine, naming its engine.
const FormatVersion = 2

// Engine is the IaC engine a manifest describes. Its value is the engine's
// segment in the release's artifact keys (releases/<tag>/modules/<engine>/).
type Engine string

const (
	// Terraform modules are archived sources, run by OpenTofu or Terraform.
	Terraform Engine = "terraform"
	// Pulumi modules are compiled binaries, one per platform.
	Pulumi Engine = "pulumi"
)

// Engines are every engine a release publishes modules for.
var Engines = []Engine{Pulumi, Terraform}

// ParseEngine reads an engine by its name.
func ParseEngine(name string) (Engine, error) {
	for _, e := range Engines {
		if name == string(e) {
			return e, nil
		}
	}
	return "", fmt.Errorf("engine %q is neither pulumi nor terraform", name)
}

// ModuleDirName is the name of an engine's module directory under a kind's
// iac/ directory in the catalog (catalog/<provider>/<kind>/iac/<name>).
func (e Engine) ModuleDirName() string {
	if e == Terraform {
		return "tf"
	}
	return string(e)
}

// EngineOfModuleDir is the engine whose module lives at dir, a catalog
// module directory (catalog/aws/awsvpc/iac/tf); ok is false for any other
// path.
func EngineOfModuleDir(dir string) (engine Engine, ok bool) {
	if path.Base(path.Dir(dir)) != "iac" {
		return "", false
	}
	for _, e := range Engines {
		if path.Base(dir) == e.ModuleDirName() {
			return e, true
		}
	}
	return "", false
}

// TerraformArtifact is the one artifact an OpenTofu/Terraform module ships:
// its module directory, archived.
const TerraformArtifact = "module.zip"

// PulumiPlatforms are the platforms every Pulumi module ships a binary for,
// as "<goos>_<goarch>". The release's Pulumi matrix builds exactly these
// (a test holds the two lists equal).
var PulumiPlatforms = []string{
	"darwin_amd64",
	"darwin_arm64",
	"linux_amd64",
	"linux_arm64",
	"windows_amd64",
}

// PulumiArtifact names a Pulumi module's gzipped binary for a platform. The
// windows binary carries the executable suffix inside the gzip name.
func PulumiArtifact(platform string) string {
	if strings.HasPrefix(platform, "windows") {
		return platform + ".exe.gz"
	}
	return platform + ".gz"
}

// Manifest is a release's module manifest for one engine.
type Manifest struct {
	FormatVersion int `json:"formatVersion"`
	// ReleaseTag is the tag the manifest describes.
	ReleaseTag string `json:"releaseTag"`
	// Engine is the engine whose modules it lists. A manifest names it so a
	// copy read away from its URL (a platform vendors one) still says which
	// modules it vouches for.
	Engine Engine `json:"engine"`
	// Modules is keyed by module directory (catalog/<provider>/<kind>/iac/<engine dir>).
	Modules map[string]Module `json:"modules"`
}

// Module is one module's published identity.
type Module struct {
	// Digest is Digest(Artifacts).
	Digest string `json:"digest"`
	// Artifacts maps each executable artifact's name (module.zip, or one
	// gzipped binary per platform) to the sha256 of its published bytes.
	Artifacts map[string]string `json:"artifacts"`
}

// Digest is a module's digest: sha256 over its artifacts' "<name> <sha256>"
// lines in name order. Only what a runner executes is listed, so a module's
// digest moves exactly when a byte a runner could execute moves.
func Digest(artifacts map[string]string) string {
	names := make([]string, 0, len(artifacts))
	for name := range artifacts {
		names = append(names, name)
	}
	sort.Strings(names)
	h := sha256.New()
	for _, name := range names {
		fmt.Fprintf(h, "%s %s\n", name, artifacts[name])
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Parse reads a manifest and proves it whole: its format, its tag, its
// engine, that every module it lists is that engine's, and that every
// module's digest is the digest of the artifacts it lists. A manifest that
// fails any of these is refused rather than trusted in part.
func Parse(data []byte) (*Manifest, error) {
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("module manifest is not valid JSON: %w", err)
	}
	if m.FormatVersion != FormatVersion {
		return nil, fmt.Errorf("module manifest format %d is not the format %d this reader knows", m.FormatVersion, FormatVersion)
	}
	if m.ReleaseTag == "" {
		return nil, fmt.Errorf("module manifest names no release tag")
	}
	if _, err := ParseEngine(string(m.Engine)); err != nil {
		return nil, fmt.Errorf("module manifest's %w", err)
	}
	for dir, module := range m.Modules {
		if engine, ok := EngineOfModuleDir(dir); !ok || engine != m.Engine {
			return nil, fmt.Errorf("module manifest for %s lists %s, which is not a %s module directory", m.Engine, dir, m.Engine)
		}
		if len(module.Artifacts) == 0 {
			return nil, fmt.Errorf("module manifest lists no artifacts for %s", dir)
		}
		if want := Digest(module.Artifacts); module.Digest != want {
			return nil, fmt.Errorf("module manifest's digest for %s is %s, but its artifacts digest to %s", dir, module.Digest, want)
		}
	}
	return &m, nil
}

// Encode writes a manifest as the release publishes it: indented JSON with
// keys in a fixed order, so the same modules always publish the same bytes.
func Encode(m *Manifest) ([]byte, error) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}
