package providerlocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Release is one published version of a provider and the platforms it ships.
type Release struct {
	Version   string
	Platforms []string
}

// Registry lists a provider's published releases.
type Registry interface {
	Releases(address string) ([]Release, error)
}

// Move is a provider the canonical lock must (re)record: the version chosen
// and the platforms its checksums are recorded for.
type Move struct {
	Address   string
	From      string // empty when the provider is new to the catalog
	Version   string
	Platforms []string
	Reason    string
}

// PlanMoves decides which providers the canonical lock must re-record. A
// locked version is kept -- a provider moves only on purpose -- unless a
// module's constraint now excludes it, or upgrade names the provider (by its
// short name, "aws", or its address). A provider that moves goes to the newest
// stable release every module's constraint allows.
func PlanMoves(modules []Module, canonical Lock, upgrade []string, registry Registry) ([]Move, error) {
	upgraded := map[string]bool{}
	for _, name := range upgrade {
		upgraded[strings.ToLower(strings.TrimSpace(name))] = true
	}
	constraints := ConstraintsByAddress(modules)
	addresses := make([]string, 0, len(constraints))
	for address := range constraints {
		addresses = append(addresses, address)
	}
	sort.Strings(addresses)

	var moves []Move
	for _, address := range addresses {
		locked, isLocked := canonical[address]
		reason := ""
		switch {
		case !isLocked:
			reason = "newly required"
		case upgraded[address] || upgraded[ShortName(address)]:
			reason = "upgrade asked for"
		default:
			ok, err := Satisfies(locked.Version, constraints[address])
			if err != nil {
				return nil, err
			}
			if ok {
				continue
			}
			reason = "a constraint excludes " + locked.Version
		}
		releases, err := registry.Releases(address)
		if err != nil {
			return nil, err
		}
		versions := make([]string, 0, len(releases))
		for _, release := range releases {
			versions = append(versions, release.Version)
		}
		chosen, err := Newest(versions, constraints[address])
		if err != nil {
			return nil, err
		}
		if chosen == "" {
			return nil, fmt.Errorf("%s: no published release satisfies every module's constraint %q",
				address, strings.Join(constraints[address], " and "))
		}
		moves = append(moves, Move{
			Address:   address,
			From:      locked.Version,
			Version:   chosen,
			Platforms: targetPlatformsOf(releases, chosen),
			Reason:    reason,
		})
	}
	for _, name := range upgrade {
		if !matchesAny(name, addresses) {
			return nil, fmt.Errorf("upgrade=%s names no provider a module requires", name)
		}
	}
	return moves, nil
}

func matchesAny(name string, addresses []string) bool {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, address := range addresses {
		if address == name || ShortName(address) == name {
			return true
		}
	}
	return false
}

// targetPlatformsOf is the target platforms the release publishes a package
// for: a provider that ships no Windows build is locked without one, rather
// than failing the whole catalog.
func targetPlatformsOf(releases []Release, version string) []string {
	published := map[string]bool{}
	for _, release := range releases {
		if release.Version == version {
			for _, platform := range release.Platforms {
				published[platform] = true
			}
		}
	}
	var platforms []string
	for _, platform := range TargetPlatforms {
		if published[platform] {
			platforms = append(platforms, platform)
		}
	}
	return platforms
}

// RecordChecksums runs `tofu providers lock` for one move in a throwaway
// module under scratchDir, requiring exactly the chosen version, and returns
// the locked provider. The packages tofu downloads to hash live only in that
// directory, which is removed before returning.
func RecordChecksums(move Move, scratchDir string) (LockedProvider, error) {
	dir, err := os.MkdirTemp(scratchDir, "lock-"+ShortName(move.Address)+"-")
	if err != nil {
		return LockedProvider{}, err
	}
	defer os.RemoveAll(dir)

	source := strings.TrimPrefix(move.Address, DefaultRegistryHost+"/")
	config := fmt.Sprintf("terraform {\n  required_providers {\n    p = {\n      source  = %q\n      version = %q\n    }\n  }\n}\n",
		source, move.Version)
	if err := os.WriteFile(filepath.Join(dir, "main.tf"), []byte(config), 0o644); err != nil {
		return LockedProvider{}, err
	}
	args := []string{"providers", "lock"}
	for _, platform := range move.Platforms {
		args = append(args, "-platform="+platform)
	}
	cmd := exec.Command("tofu", args...)
	cmd.Dir = dir
	// A plugin cache would keep every platform's package after the run; the
	// lock needs only their checksums.
	cmd.Env = append(os.Environ(), "TF_PLUGIN_CACHE_DIR=")
	if out, err := cmd.CombinedOutput(); err != nil {
		return LockedProvider{}, fmt.Errorf("tofu providers lock %s %s: %w\n%s", move.Address, move.Version, err, out)
	}
	lock, err := ReadLock(filepath.Join(dir, LockFileName))
	if err != nil {
		return LockedProvider{}, err
	}
	locked, ok := lock[move.Address]
	if !ok {
		return LockedProvider{}, fmt.Errorf("tofu providers lock wrote no entry for %s", move.Address)
	}
	return LockedProvider{Address: move.Address, Version: locked.Version, Hashes: locked.Hashes}, nil
}

// HTTPRegistry reads releases from a provider registry's public protocol
// (GET https://<host>/v1/providers/<namespace>/<type>/versions).
type HTTPRegistry struct {
	Client *http.Client
}

// Releases implements Registry.
func (r HTTPRegistry) Releases(address string) ([]Release, error) {
	parts := strings.Split(address, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("provider address %q is not <host>/<namespace>/<type>", address)
	}
	client := r.Client
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	url := fmt.Sprintf("https://%s/v1/providers/%s/%s/versions", parts[0], parts[1], parts[2])
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("list releases of %s: %w", address, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list releases of %s: %s returned %s", address, url, resp.Status)
	}
	var body struct {
		Versions []struct {
			Version   string `json:"version"`
			Platforms []struct {
				OS   string `json:"os"`
				Arch string `json:"arch"`
			} `json:"platforms"`
		} `json:"versions"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("list releases of %s: %w", address, err)
	}
	releases := make([]Release, 0, len(body.Versions))
	for _, v := range body.Versions {
		release := Release{Version: v.Version}
		for _, p := range v.Platforms {
			release.Platforms = append(release.Platforms, p.OS+"_"+p.Arch)
		}
		releases = append(releases, release)
	}
	return releases, nil
}
