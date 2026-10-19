package actioninventory

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

// PrivilegedFileName is the authored list of privileged permissions inside
// this package's directory. Unlike the snapshots beside it, it is written
// by hand: no provider publishes which of its permissions let a principal
// act as another identity or grant access across an account, so the line
// is drawn here, once, and every reader takes it from here.
const PrivilegedFileName = "privileged.yaml"

// The privilege values, spelled as the permissions schema's
// PermissionPrivilege enum spells them.
const (
	PrivilegeOrdinary     = "ordinary"
	PrivilegeUsesIdentity = "uses_identity"
	PrivilegeGrantsAccess = "grants_access"
)

// Privileged is the authored privileged-permission list for every provider
// whose manifests carry a privilege.
type Privileged struct {
	Providers []PrivilegedProvider `yaml:"providers"`
}

// PrivilegedProvider is one provider's privileged permissions, in the
// provider's own exact spelling.
type PrivilegedProvider struct {
	// Provider is the catalog provider directory name (e.g. "gcp").
	Provider string `yaml:"provider"`
	// UsesIdentity lets a resource run as an identity that already exists.
	UsesIdentity []string `yaml:"uses_identity"`
	// GrantsAccess changes who holds power over the whole account.
	GrantsAccess []string `yaml:"grants_access"`
}

// LoadPrivileged reads and strictly parses the authored list and enforces
// its invariants: providers sorted and unique, each list sorted and
// unique, and no permission on both lists.
func LoadPrivileged(dir string) (*Privileged, error) {
	raw, err := os.ReadFile(path.Join(dir, PrivilegedFileName))
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(strings.NewReader(string(raw)))
	decoder.KnownFields(true)
	list := &Privileged{}
	if err := decoder.Decode(list); err != nil {
		return nil, fmt.Errorf("parse %s: %w", PrivilegedFileName, err)
	}
	for i, provider := range list.Providers {
		if provider.Provider == "" {
			return nil, fmt.Errorf("%s: provider %d has no name", PrivilegedFileName, i)
		}
		if i > 0 && list.Providers[i-1].Provider >= provider.Provider {
			return nil, fmt.Errorf("%s: providers not sorted or duplicated at %q", PrivilegedFileName, provider.Provider)
		}
		if err := checkActionList(PrivilegedFileName, provider.Provider, "uses_identity", provider.UsesIdentity); err != nil {
			return nil, err
		}
		if err := checkActionList(PrivilegedFileName, provider.Provider, "grants_access", provider.GrantsAccess); err != nil {
			return nil, err
		}
		for _, permission := range provider.UsesIdentity {
			if containsString(provider.GrantsAccess, permission) {
				return nil, fmt.Errorf("%s: %s permission %q is on both lists", PrivilegedFileName, provider.Provider, permission)
			}
		}
	}
	return list, nil
}

// For returns one provider's lists, or nil when the provider carries no
// privilege.
func (p *Privileged) For(provider string) *PrivilegedProvider {
	for i := range p.Providers {
		if p.Providers[i].Provider == provider {
			return &p.Providers[i]
		}
	}
	return nil
}

// Of returns the privilege of a set of permissions: the highest of theirs.
// AWS and Azure evaluate action names case-insensitively and a manifest
// action may be a wildcard pattern, so a pattern covering a listed
// permission carries its privilege; GCP permission names are exact.
func (p *PrivilegedProvider) Of(permissions []string) string {
	privilege := PrivilegeOrdinary
	for _, permission := range permissions {
		if p.matches(p.GrantsAccess, permission) {
			return PrivilegeGrantsAccess
		}
		if p.matches(p.UsesIdentity, permission) {
			privilege = PrivilegeUsesIdentity
		}
	}
	return privilege
}

func (p *PrivilegedProvider) matches(listed []string, permission string) bool {
	if p.Provider == "gcp" {
		return containsString(listed, permission)
	}
	pattern := strings.ToLower(permission)
	for _, name := range listed {
		if globMatch(pattern, strings.ToLower(name)) {
			return true
		}
	}
	return false
}
