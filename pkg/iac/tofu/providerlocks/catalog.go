package providerlocks

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

// CanonicalLockPath is the catalog's one source of truth for provider
// versions, relative to the repository root. It lives in an underscore
// directory, beside the catalog's other non-provider content, so nothing that
// walks catalog/<provider>/<kind> mistakes it for a kind.
const CanonicalLockPath = "catalog/_providers/" + LockFileName

// TargetPlatforms are the platforms a module can run on -- the platforms the
// runner and the CLI ship for -- so every lock carries a package checksum for
// each. A provider that publishes fewer platforms is locked for the ones it
// has.
var TargetPlatforms = []string{"darwin_amd64", "darwin_arm64", "linux_amd64", "linux_arm64", "windows_amd64"}

var moduleDir = regexp.MustCompile(`^catalog/[^/]+/[^/]+/iac/tf$`)

// Module is one OpenTofu module and the providers it requires.
type Module struct {
	// Dir is the module directory relative to the repository root.
	Dir          string
	Requirements map[string]Requirement
}

// Modules reads every OpenTofu module under repoRoot's catalog
// (catalog/<provider>/<kind>/iac/tf), sorted by directory.
func Modules(repoRoot string) ([]Module, error) {
	dirs, err := filepath.Glob(filepath.Join(repoRoot, "catalog", "*", "*", "iac", "tf"))
	if err != nil {
		return nil, err
	}
	var modules []Module
	for _, dir := range dirs {
		rel, err := filepath.Rel(repoRoot, dir)
		if err != nil {
			return nil, err
		}
		rel = filepath.ToSlash(rel)
		if info, err := os.Stat(dir); err != nil || !info.IsDir() || !moduleDir.MatchString(rel) {
			continue
		}
		requirements, err := ModuleRequirements(dir)
		if err != nil {
			return nil, err
		}
		modules = append(modules, Module{Dir: rel, Requirements: requirements})
	}
	sort.Slice(modules, func(i, j int) bool { return modules[i].Dir < modules[j].Dir })
	return modules, nil
}

// ConstraintsByAddress collects every module's constraint for each provider.
func ConstraintsByAddress(modules []Module) map[string][]string {
	byAddress := map[string][]string{}
	for _, module := range modules {
		for address, requirement := range module.Requirements {
			byAddress[address] = append(byAddress[address], requirement.Constraints)
		}
	}
	return byAddress
}

// Project is a module's lock: exactly the providers it requires, at the
// canonical versions and checksums, with its own constraints. A module that
// requires no provider has no lock (nil).
func Project(canonical Lock, module Module) ([]byte, error) {
	if len(module.Requirements) == 0 {
		return nil, nil
	}
	projected := Lock{}
	for address, requirement := range module.Requirements {
		locked, ok := canonical[address]
		if !ok {
			return nil, fmt.Errorf("%s requires %s, which %s does not lock", module.Dir, address, CanonicalLockPath)
		}
		projected[address] = LockedProvider{
			Address:     address,
			Version:     locked.Version,
			Constraints: requirement.Constraints,
			Hashes:      locked.Hashes,
		}
	}
	return Render(projected), nil
}

// Check proves, offline, that the committed locks are current:
//   - the canonical lock holds every provider a module requires, at a version
//     every module's constraint allows, and nothing no module requires;
//   - every module's committed lock is exactly its projection, and a module
//     that requires no provider has none.
//
// It returns one sentence per problem, each saying how to fix it.
func Check(repoRoot string) ([]string, error) {
	modules, err := Modules(repoRoot)
	if err != nil {
		return nil, err
	}
	canonical, err := ReadLock(filepath.Join(repoRoot, CanonicalLockPath))
	if err != nil {
		return nil, err
	}
	var problems []string
	constraints := ConstraintsByAddress(modules)
	for address, all := range constraints {
		locked, ok := canonical[address]
		if !ok {
			problems = append(problems, fmt.Sprintf("%s is required by a module but not locked; run `make lock-providers`", address))
			continue
		}
		ok, err := Satisfies(locked.Version, all)
		if err != nil {
			return nil, err
		}
		if !ok {
			problems = append(problems, fmt.Sprintf("%s is locked at %s, which a module's constraint excludes; run `make lock-providers`", address, locked.Version))
		}
	}
	for address := range canonical {
		if _, ok := constraints[address]; !ok {
			problems = append(problems, fmt.Sprintf("%s is locked but no module requires it; run `make lock-providers`", address))
		}
	}
	for _, module := range modules {
		path := filepath.Join(repoRoot, module.Dir, LockFileName)
		committed, err := os.ReadFile(path)
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
		want, err := Project(canonical, module)
		if err != nil {
			continue // already reported as an unlocked provider
		}
		switch {
		case want == nil && committed != nil:
			problems = append(problems, fmt.Sprintf("%s requires no provider but has a lock file; run `make lock-providers`", module.Dir))
		case want != nil && committed == nil:
			problems = append(problems, fmt.Sprintf("%s has no lock file; run `make lock-providers`", module.Dir))
		case !bytes.Equal(want, committed):
			problems = append(problems, fmt.Sprintf("%s's lock file is not its projection of %s; run `make lock-providers`", module.Dir, CanonicalLockPath))
		}
	}
	sort.Strings(problems)
	return problems, nil
}

// WriteProjections writes every module's projection, and removes the lock of
// a module that requires no provider.
func WriteProjections(repoRoot string, canonical Lock, modules []Module) error {
	for _, module := range modules {
		path := filepath.Join(repoRoot, module.Dir, LockFileName)
		projected, err := Project(canonical, module)
		if err != nil {
			return err
		}
		if projected == nil {
			if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
				return err
			}
			continue
		}
		if err := os.WriteFile(path, projected, 0o644); err != nil {
			return err
		}
	}
	return nil
}
