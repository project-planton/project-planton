//go:build !codegen
// +build !codegen

package refcheck

import (
	"regexp"
	"strings"

	"github.com/plantonhq/planton/pkg/catalogkindreflect"
	"github.com/plantonhq/planton/shared/catalogkind"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ResolveValueFromPath validates a manifest-level valueFrom reference (a ValueFromRef's
// kind + field_path) against the referenced kind's descriptor and returns a human-readable
// reason when it does not resolve, or "" when it does.
//
// This is the manifest-side sibling of the annotation check in Analyze: the annotation
// check guards `default_kind_field_path` values authored in protos, while this guards
// `fieldPath` values authored in YAML (infra-chart templates, presets). The resolution
// semantics deliberately mirror the control plane's reference validator so a reference
// accepted here is accepted there:
//   - the path is rooted at the kind's top-level API message (so it starts with
//     "status.", "spec.", or "metadata.");
//   - segments resolve by exact proto field name first, then by a camelCase->snake_case
//     fallback (chart authors write either form);
//   - bracketed or bare integer segments index into repeated fields: into a repeated
//     message to descend further, or into a list of strings to name one element, which
//     then ends the path (a zone's name servers referenced one per record as
//     "status.outputs.nameservers.0"); a list of strings without an index is refused,
//     since a reference feeds one string and the flattened outputs hold only elements;
//   - a map field is addressed by KEY: the segment after it is the entry key (any
//     non-empty string -- keys are runtime data, e.g. a load balancer's name-keyed
//     pool-id outputs referenced as "status.outputs.backend_pool_ids.web"), never a
//     field on the synthetic map-entry message. Keys containing dots cannot be
//     addressed through a dot path -- name sub-resources accordingly;
//   - every non-terminal segment must be a message, and the terminal value must be a
//     string (a valueFrom always feeds a string-valued reference).
func ResolveValueFromPath(kind catalogkind.CatalogKind, fieldPath string) string {
	_, reason := walkValueFromPath(kind, fieldPath)
	return reason
}

// CanonicalValueFromPath returns the path a valueFrom reference reads, in the
// form the resource's stored document carries it: every schema field segment
// by its proto field name (a camelCase segment becomes snake_case), every
// index bare ("[0]" becomes "0"), and map keys exactly as written. A path
// that does not resolve is returned unchanged; ResolveValueFromPath says why.
func CanonicalValueFromPath(kind catalogkind.CatalogKind, fieldPath string) string {
	canonical, reason := walkValueFromPath(kind, fieldPath)
	if reason != "" {
		return fieldPath
	}
	return strings.Join(canonical, ".")
}

// walkValueFromPath walks fieldPath down kind's descriptor, returning its
// canonical segments, or the reason it does not resolve.
func walkValueFromPath(kind catalogkind.CatalogKind, fieldPath string) ([]string, string) {
	inst, err := catalogkindreflect.NewInstance(kind)
	if err != nil {
		return nil, "kind " + kind.String() + " is not a registered/implemented kind"
	}
	if fieldPath == "" {
		return nil, "fieldPath is empty"
	}

	current := inst.ProtoReflect().Descriptor()
	var fd protoreflect.FieldDescriptor
	segments := strings.Split(fieldPath, ".")
	canonical := make([]string, 0, len(segments))

	for i := 0; i < len(segments); i++ {
		seg := segments[i]
		if isIndexSegment(seg) {
			canonical = append(canonical, strings.TrimSuffix(strings.TrimPrefix(seg, "["), "]"))
			if fd == nil || !fd.IsList() {
				return nil, "cannot index into '" + seg + "' -- preceding field is not a list"
			}
			if fd.Kind() == protoreflect.MessageKind {
				current = fd.Message()
				continue
			}
			// A list of plain values: the index names one element, and the element
			// is the value -- nothing lies below it.
			if i != len(segments)-1 {
				return nil, "cannot descend into element '" + seg + "' of the " + fd.Kind().String() + " list '" + string(fd.Name()) + "'"
			}
			if fd.Kind() != protoreflect.StringKind {
				return nil, "elements of '" + string(fd.Name()) + "' are " + fd.Kind().String() + ", expected string"
			}
			return canonical, ""
		}

		fd = current.Fields().ByName(protoreflect.Name(seg))
		if fd == nil {
			fd = current.Fields().ByName(protoreflect.Name(camelToSnake(seg)))
		}
		if fd == nil {
			return nil, "no field named '" + seg + "' on " + string(current.FullName())
		}
		canonical = append(canonical, string(fd.Name()))

		if fd.IsMap() {
			if i == len(segments)-1 {
				return nil, "field '" + seg + "' is a map -- address one entry by appending its key (e.g. '" + fieldPath + ".<key>')"
			}
			i++ // the next segment is the entry KEY -- runtime data, any non-empty string
			if segments[i] == "" {
				return nil, "empty map key after '" + seg + "'"
			}
			canonical = append(canonical, segments[i])
			switch fd.MapValue().Kind() {
			case protoreflect.StringKind:
				if i != len(segments)-1 {
					return nil, "cannot descend into the string map value of '" + seg + "'"
				}
				return canonical, ""
			case protoreflect.MessageKind:
				if i == len(segments)-1 {
					return nil, "terminal field '" + fieldPath + "' is a message map value, expected string"
				}
				current = fd.MapValue().Message()
				fd = nil
				continue
			default:
				return nil, "map value of '" + seg + "' is " + fd.MapValue().Kind().String() + ", expected string"
			}
		}

		if i < len(segments)-1 {
			if fd.IsList() && fd.Kind() != protoreflect.MessageKind {
				continue // the next segment must index one element; the index branch judges it
			}
			if fd.Kind() != protoreflect.MessageKind {
				return nil, "cannot descend into scalar field '" + seg + "'"
			}
			current = fd.Message()
		}
	}

	if fd != nil && fd.IsList() && fd.Kind() != protoreflect.MessageKind {
		return nil, "field '" + string(fd.Name()) + "' is a list -- address one element by appending its index (e.g. '" + fieldPath + ".0')"
	}
	if fd == nil || fd.Kind() != protoreflect.StringKind {
		kindName := "not a field"
		if fd != nil {
			kindName = fd.Kind().String()
		}
		return nil, "terminal field '" + fieldPath + "' is " + kindName + ", expected string"
	}
	return canonical, ""
}

var camelBoundary = regexp.MustCompile(`([a-z0-9])([A-Z])`)

// camelToSnake converts camelCase to snake_case, matching the control plane's fallback
// so chart authors may write either the protojson (camelCase) or proto (snake_case) form.
func camelToSnake(s string) string {
	return strings.ToLower(camelBoundary.ReplaceAllString(s, "${1}_${2}"))
}
