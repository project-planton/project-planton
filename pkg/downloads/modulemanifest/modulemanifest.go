// Package modulemanifest is the grammar of a release's module manifest: the
// one document that says which bytes every platform IaC module of a release
// is, and what each module's digest is.
//
// A release publishes the manifest at releases/<tag>/modules/manifest.json
// once every module artifact is uploaded. It is keyed by module directory,
// exactly as the catalog names it (catalog/aws/awsrdsinstance/iac/tf), and
// records, for every artifact a runner may execute, the sha256 of the bytes
// actually published under the tag.
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
	"sort"
	"strings"
)

// FileName is the manifest's file name under a release's modules/ prefix.
const FileName = "manifest.json"

// FormatVersion is the manifest shape this package writes and reads.
const FormatVersion = 1

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

// Manifest is a release's module manifest.
type Manifest struct {
	FormatVersion int `json:"formatVersion"`
	// ReleaseTag is the tag the manifest describes.
	ReleaseTag string `json:"releaseTag"`
	// Modules is keyed by module directory (catalog/<provider>/<kind>/iac/<engine>).
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

// Parse reads a manifest and proves it whole: its format, its tag, and that
// every module's digest is the digest of the artifacts it lists. A manifest
// that fails any of these is refused rather than trusted in part.
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
	for dir, module := range m.Modules {
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
