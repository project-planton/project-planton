package providerlocks

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// DefaultRegistryHost is the registry a provider source without a hostname
// resolves to. The runner and the catalog's tooling execute modules with
// OpenTofu, so "hashicorp/aws" is registry.opentofu.org/hashicorp/aws -- the
// address OpenTofu writes into a lock file.
const DefaultRegistryHost = "registry.opentofu.org"

// Requirement is one provider a module requires: its fully qualified address
// and the module's version constraint, normalized the way OpenTofu writes it
// into a lock file.
type Requirement struct {
	Address     string
	Constraints string
}

// ModuleRequirements reads the provider requirements a module's *.tf files
// declare in terraform { required_providers { ... } }, keyed by address. A
// module that declares no provider returns an empty map -- it needs no lock.
//
// A provider declared twice must be declared identically: two constraints for
// one provider in one module would make the module's lock ambiguous.
func ModuleRequirements(moduleDir string) (map[string]Requirement, error) {
	files, err := filepath.Glob(filepath.Join(moduleDir, "*.tf"))
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	requirements := map[string]Requirement{}
	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		parsed, diags := hclsyntax.ParseConfig(src, file, hcl.InitialPos)
		if diags.HasErrors() {
			return nil, fmt.Errorf("parse %s: %s", file, diags.Error())
		}
		body, ok := parsed.Body.(*hclsyntax.Body)
		if !ok {
			continue
		}
		for _, terraform := range body.Blocks {
			if terraform.Type != "terraform" {
				continue
			}
			for _, block := range terraform.Body.Blocks {
				if block.Type != "required_providers" {
					continue
				}
				if err := collectRequirements(file, block.Body, requirements); err != nil {
					return nil, err
				}
			}
		}
	}
	return requirements, nil
}

func collectRequirements(file string, body *hclsyntax.Body, into map[string]Requirement) error {
	for name, attribute := range body.Attributes {
		value, diags := attribute.Expr.Value(nil)
		if diags.HasErrors() {
			return fmt.Errorf("%s: required_providers.%s must be a literal: %s", file, name, diags.Error())
		}
		source, version := stringAttribute(value, "source"), stringAttribute(value, "version")
		if source == "" {
			// A provider without a source is hashicorp/<name>: the implicit
			// default every Terraform-language tool applies.
			source = "hashicorp/" + name
		}
		address, err := Address(source)
		if err != nil {
			return fmt.Errorf("%s: required_providers.%s: %w", file, name, err)
		}
		constraints := ""
		if version != "" {
			constraints, err = NormalizeConstraints(version)
			if err != nil {
				return fmt.Errorf("%s: required_providers.%s: %w", file, name, err)
			}
		}
		requirement := Requirement{Address: address, Constraints: constraints}
		if existing, ok := into[address]; ok && existing != requirement {
			return fmt.Errorf("%s: %s is required twice with different constraints (%q and %q)",
				file, address, existing.Constraints, constraints)
		}
		into[address] = requirement
	}
	return nil
}

func stringAttribute(value cty.Value, name string) string {
	if !value.Type().IsObjectType() || !value.Type().HasAttribute(name) {
		return ""
	}
	attribute := value.GetAttr(name)
	if attribute.IsNull() || !attribute.IsKnown() || attribute.Type() != cty.String {
		return ""
	}
	return attribute.AsString()
}

// Address turns a provider source ("hashicorp/aws", or a hostname-qualified
// "registry.example.com/acme/thing") into the lowercase, fully qualified
// address a lock file is keyed by.
func Address(source string) (string, error) {
	parts := strings.Split(strings.ToLower(strings.TrimSpace(source)), "/")
	switch len(parts) {
	case 2:
		return DefaultRegistryHost + "/" + parts[0] + "/" + parts[1], nil
	case 3:
		return strings.Join(parts, "/"), nil
	default:
		return "", fmt.Errorf("provider source %q is not <namespace>/<type> or <host>/<namespace>/<type>", source)
	}
}
