package providerlocks

import (
	"fmt"
	"regexp"
	"strings"

	version "github.com/hashicorp/go-version"
)

var constraintPart = regexp.MustCompile(`^\s*(~>|>=|<=|!=|>|<|=)?\s*v?(\d+(?:\.\d+){0,2})\s*$`)

// NormalizeConstraints writes a version constraint the way OpenTofu records it
// in a lock file's `constraints` attribute, so a projected lock is exactly
// what `tofu init` would write and an init never rewrites it:
//
//   - an exact version is written bare, in full: "3.9.1";
//   - "~>" keeps the precision it was given (it decides what may move), with at
//     least major.minor: "~> 3" is "~> 3.0", "~> 3.6.0" stays "~> 3.6.0";
//   - every other operator takes the full major.minor.patch: ">= 3.0" is
//     ">= 3.0.0".
//
// Pre-release constraints are refused: no catalog module pins one, and a lock
// built from one would hold a version no stable constraint allows.
func NormalizeConstraints(raw string) (string, error) {
	var parts []string
	for _, piece := range strings.Split(raw, ",") {
		match := constraintPart.FindStringSubmatch(piece)
		if match == nil {
			return "", fmt.Errorf("version constraint %q: %q is not <operator> <version>", raw, strings.TrimSpace(piece))
		}
		operator, segments := match[1], strings.Split(match[2], ".")
		switch operator {
		case "", "=":
			parts = append(parts, full(segments))
		case "~>":
			for len(segments) < 2 {
				segments = append(segments, "0")
			}
			parts = append(parts, "~> "+strings.Join(segments, "."))
		default:
			parts = append(parts, operator+" "+full(segments))
		}
	}
	return strings.Join(parts, ", "), nil
}

func full(segments []string) string {
	for len(segments) < 3 {
		segments = append(segments, "0")
	}
	return strings.Join(segments, ".")
}

// Satisfies reports whether a provider version meets every constraint in the
// list (each one a module's constraint for that provider). An empty constraint
// allows any version.
func Satisfies(providerVersion string, constraints []string) (bool, error) {
	candidate, err := version.NewVersion(providerVersion)
	if err != nil {
		return false, fmt.Errorf("provider version %q: %w", providerVersion, err)
	}
	for _, raw := range constraints {
		if raw == "" {
			continue
		}
		constraint, err := version.NewConstraint(raw)
		if err != nil {
			return false, fmt.Errorf("version constraint %q: %w", raw, err)
		}
		if !constraint.Check(candidate) {
			return false, nil
		}
	}
	return true, nil
}

// Newest returns the newest stable version among available that meets every
// constraint, or "" when none does. Pre-releases are never chosen.
func Newest(available []string, constraints []string) (string, error) {
	var best *version.Version
	for _, raw := range available {
		candidate, err := version.NewVersion(raw)
		if err != nil || candidate.Prerelease() != "" {
			continue
		}
		ok, err := Satisfies(candidate.String(), constraints)
		if err != nil {
			return "", err
		}
		if ok && (best == nil || candidate.GreaterThan(best)) {
			best = candidate
		}
	}
	if best == nil {
		return "", nil
	}
	return best.String(), nil
}
