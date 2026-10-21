package manifestgraph

import (
	"fmt"

	"github.com/plantonhq/planton/shared/catalogkind"
	foreignkeyv1 "github.com/plantonhq/planton/shared/foreignkey/v1"
	"google.golang.org/protobuf/proto"
)

// OutputsLookup supplies what a reference can read from a deployed resource,
// by graph identity: its FLATTENED outputs (dotted string keys, the
// platform's flattening), and the fields its manifest declares under their
// full paths (ManifestFields, flattened: "spec.name"). Returning false
// means the identity's outputs are not available — the reference is left
// untouched and reported, and the caller's policy decides what that means
// (a preflighted deployment treats it as a bug; the E2E harness treats it as
// "not deployed by this chain, leave it for the engine's own wall").
type OutputsLookup func(id Identity) (map[string]string, bool)

// ResolveRefs rewrites the manifest's valueFrom references to literal values
// read from deployed resources' outputs. consumerEnv is the manifest's own
// env — the fallback for references that name none.
//
// Semantics, per reference:
//   - the effective kind and field path come from the strict rules (CheckRef);
//     a reference the rules reject is reported and left untouched;
//   - the target's outputs come from the lookup; a lookup miss leaves the
//     reference untouched and reports FindingUnresolvedRef;
//   - a lookup HIT that lacks the referenced field reports
//     FindingMissingOutput — the sharper class, because the resource exists
//     and the composition names a field it does not export or declare;
//   - on success the whole StringValueOrRef is replaced with the literal arm.
//
// Field paths are read canonically (LookupKey): proto field names, so a
// camelCase path reads the same value as its snake_case form. Outputs read by
// their flattened key, the "status.outputs." prefix stripped (map-typed
// composition keys address entries by suffix, e.g. "backend_pool_ids.web");
// a manifest field reads by its full path ("spec.name").
func ResolveRefs(msg proto.Message, consumerEnv string, lookup OutputsLookup) (int, []Finding) {
	resolved := 0
	var findings []Finding

	for _, use := range CollectRefUses(msg) {
		target, problems := CheckRef(use)
		if len(problems) > 0 {
			for _, p := range problems {
				findings = append(findings, Finding{
					Class: FindingRefRule, FieldPath: use.FieldPath, Message: p,
				})
			}
			continue
		}
		if target.Kind == catalogkind.CatalogKind_unspecified || target.Name == "" {
			continue
		}

		targetID := target.Identity(consumerEnv)
		outs, ok := lookup(targetID)
		if !ok {
			t := target
			findings = append(findings, Finding{
				Class: FindingUnresolvedRef, FieldPath: use.FieldPath, Target: &t,
				Message: fmt.Sprintf("%s: no outputs available for %s — the reference stays unresolved", use.FieldPath, targetID),
			})
			continue
		}

		key := LookupKey(target)
		value, ok := outs[key]
		if !ok {
			t := target
			message := fmt.Sprintf("%s: %s has no output %q — the resource deployed but does not export the referenced field", use.FieldPath, targetID, key)
			if isManifestField(key) {
				message = fmt.Sprintf("%s: %s does not set %q in its manifest — set it there, or reference a field it does set", use.FieldPath, targetID, key)
			}
			findings = append(findings, Finding{
				Class: FindingMissingOutput, FieldPath: use.FieldPath, Target: &t,
				Message: message,
			})
			continue
		}

		use.Replace(&foreignkeyv1.StringValueOrRef{
			LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: value},
		})
		resolved++
	}

	return resolved, findings
}
