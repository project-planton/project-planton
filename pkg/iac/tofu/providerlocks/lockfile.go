package providerlocks

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// LockFileName is the dependency lock file OpenTofu reads beside a module.
const LockFileName = ".terraform.lock.hcl"

// LockedProvider is one provider block of a dependency lock file.
type LockedProvider struct {
	// Address is the fully qualified provider address, the block's label.
	Address string
	// Version is the exact release every module runs.
	Version string
	// Constraints is the constraint the lock was written for. The canonical
	// lock leaves it empty: its versions answer every module at once, and each
	// module's projection carries that module's own constraint.
	Constraints string
	// Hashes are the package checksums OpenTofu verifies before it runs the
	// provider: "h1:" per platform package, "zh:" per published zip.
	Hashes []string
}

// Lock is a dependency lock file: its providers by address.
type Lock map[string]LockedProvider

// ReadLock parses a dependency lock file. A missing file is an empty lock.
func ReadLock(path string) (Lock, error) {
	src, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Lock{}, nil
	}
	if err != nil {
		return nil, err
	}
	return ParseLock(src, path)
}

// ParseLock parses dependency lock file source.
func ParseLock(src []byte, filename string) (Lock, error) {
	parsed, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, fmt.Errorf("parse %s: %s", filename, diags.Error())
	}
	lock := Lock{}
	for _, block := range parsed.Body.(*hclsyntax.Body).Blocks {
		if block.Type != "provider" || len(block.Labels) != 1 {
			return nil, fmt.Errorf("%s: unexpected block %q", filename, block.Type)
		}
		provider := LockedProvider{Address: block.Labels[0]}
		for name, attribute := range block.Body.Attributes {
			value, diags := attribute.Expr.Value(nil)
			if diags.HasErrors() {
				return nil, fmt.Errorf("%s: %s.%s: %s", filename, provider.Address, name, diags.Error())
			}
			switch name {
			case "version":
				provider.Version = value.AsString()
			case "constraints":
				provider.Constraints = value.AsString()
			case "hashes":
				for it := value.ElementIterator(); it.Next(); {
					_, hash := it.Element()
					if hash.Type() != cty.String {
						return nil, fmt.Errorf("%s: %s.hashes holds a non-string", filename, provider.Address)
					}
					provider.Hashes = append(provider.Hashes, hash.AsString())
				}
			default:
				return nil, fmt.Errorf("%s: %s has unexpected attribute %q", filename, provider.Address, name)
			}
		}
		lock[provider.Address] = provider
	}
	return lock, nil
}

// Render writes a lock in exactly OpenTofu's own format -- its header, blocks
// sorted by address, hclwrite's attribute alignment, hashes sorted -- so a
// lock written here is the file `tofu init` would write, byte for byte.
func Render(lock Lock) []byte {
	var out bytes.Buffer
	out.WriteString("# This file is maintained automatically by \"tofu init\".\n")
	out.WriteString("# Manual edits may be lost in future updates.\n")
	for _, address := range lock.addresses() {
		provider := lock[address]
		fmt.Fprintf(&out, "\nprovider %q {\n", provider.Address)
		if provider.Constraints == "" {
			fmt.Fprintf(&out, "  version = %q\n", provider.Version)
		} else {
			// hclwrite aligns the "=" of consecutive one-line attributes.
			fmt.Fprintf(&out, "  version     = %q\n", provider.Version)
			fmt.Fprintf(&out, "  constraints = %q\n", provider.Constraints)
		}
		out.WriteString("  hashes = [\n")
		hashes := append([]string(nil), provider.Hashes...)
		sort.Strings(hashes)
		for _, hash := range hashes {
			fmt.Fprintf(&out, "    %q,\n", hash)
		}
		out.WriteString("  ]\n}\n")
	}
	return out.Bytes()
}

func (lock Lock) addresses() []string {
	addresses := make([]string, 0, len(lock))
	for address := range lock {
		addresses = append(addresses, address)
	}
	sort.Strings(addresses)
	return addresses
}

// ShortName is the provider's type, the last segment of its address
// ("registry.opentofu.org/hashicorp/aws" is "aws"): the name people use.
func ShortName(address string) string {
	return address[strings.LastIndex(address, "/")+1:]
}
