package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/plantonhq/planton/pkg/downloads/modulemanifest"
)

// A part is one release leg's record of what it published: one line per
// artifact, "<module dir>\t<artifact name>\t<sha256>", where the checksum is
// the published object's (publish_once.sh reports it, including for an
// object a re-run leg kept). Parts are plain text so a leg's shell loop can
// append to one as it uploads.

var sha256Hex = regexp.MustCompile(`^[0-9a-f]{64}$`)

// ReadParts reads every part file in dir into module dir -> artifact -> sha256.
// The same artifact recorded twice must carry the same checksum.
func ReadParts(dir string) (map[string]map[string]string, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
	published := map[string]map[string]string{}
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			// actions/download-artifact puts each artifact in its own folder.
			nested, err := ReadParts(path)
			if err != nil {
				return nil, err
			}
			for module, artifacts := range nested {
				for name, sum := range artifacts {
					if err := record(published, module, name, sum, path); err != nil {
						return nil, err
					}
				}
			}
			continue
		}
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(f)
		line := 0
		for scanner.Scan() {
			line++
			text := strings.TrimSpace(scanner.Text())
			if text == "" {
				continue
			}
			fields := strings.Split(text, "\t")
			if len(fields) != 3 || !sha256Hex.MatchString(fields[2]) {
				f.Close()
				return nil, fmt.Errorf("%s:%d is not \"<module dir>\\t<artifact>\\t<sha256>\": %q", path, line, text)
			}
			if err := record(published, fields[0], fields[1], fields[2], path); err != nil {
				f.Close()
				return nil, err
			}
		}
		f.Close()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return published, nil
}

func record(published map[string]map[string]string, module, name, sum, source string) error {
	if published[module] == nil {
		published[module] = map[string]string{}
	}
	if prior, ok := published[module][name]; ok && prior != sum {
		return fmt.Errorf("%s %s is recorded with two checksums (%s, %s in %s)", module, name, prior, sum, source)
	}
	published[module][name] = sum
	return nil
}

// expectedArtifacts is what a module must have published for the release to
// describe it: every platform's binary for a Pulumi module, module.zip for an
// OpenTofu/Terraform one.
func expectedArtifacts(m Module) []string {
	if m.Engine == Pulumi {
		names := make([]string, 0, len(modulemanifest.PulumiPlatforms))
		for _, p := range modulemanifest.PulumiPlatforms {
			names = append(names, modulemanifest.PulumiArtifact(p))
		}
		return names
	}
	return []string{modulemanifest.TerraformArtifact}
}

// BuildManifest assembles the release's module manifest from what its legs
// published, and refuses an incomplete one: a module the release ships but no
// leg recorded, a module missing an artifact, or an artifact nobody expects.
// A manifest that described part of a release would let a platform trust a
// module it cannot verify, so the release publishes none until every leg has.
func BuildManifest(tag string, modules []Module, published map[string]map[string]string) ([]byte, error) {
	manifest := &modulemanifest.Manifest{
		FormatVersion: modulemanifest.FormatVersion,
		ReleaseTag:    tag,
		Modules:       map[string]modulemanifest.Module{},
	}
	var problems []string
	shipped := map[string]bool{}
	for _, m := range modules {
		shipped[m.Dir] = true
		got := published[m.Dir]
		artifacts := map[string]string{}
		for _, name := range expectedArtifacts(m) {
			sum, ok := got[name]
			if !ok {
				problems = append(problems, fmt.Sprintf("%s: %s was not published", m.Dir, name))
				continue
			}
			artifacts[name] = sum
		}
		for name := range got {
			if !contains(expectedArtifacts(m), name) {
				problems = append(problems, fmt.Sprintf("%s: %s is not an artifact this module ships", m.Dir, name))
			}
		}
		manifest.Modules[m.Dir] = modulemanifest.Module{Digest: modulemanifest.Digest(artifacts), Artifacts: artifacts}
	}
	for dir := range published {
		if !shipped[dir] {
			problems = append(problems, fmt.Sprintf("%s: published, but the release ships no such module", dir))
		}
	}
	if len(problems) > 0 {
		sort.Strings(problems)
		return nil, fmt.Errorf("the release's modules are incomplete, so no manifest is published:\n  %s", strings.Join(problems, "\n  "))
	}
	return modulemanifest.Encode(manifest)
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
