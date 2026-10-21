package manifestgraph

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/plantonhq/planton/pkg/refcheck"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ManifestFields returns what a valueFrom reference can read from a
// resource's own manifest: its spec and metadata, keyed by proto field name
// the way the control plane stores a resource, as nested maps. Flattened
// beside the resource's outputs they give the lookup's "spec.name" and
// "metadata.labels.team" keys, so a reference to a field the referenced
// manifest declares resolves the same way it does on the platform.
func ManifestFields(msg proto.Message) (map[string]interface{}, error) {
	raw, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(msg)
	if err != nil {
		return nil, errors.Wrap(err, "serializing the manifest to read its fields")
	}
	var document map[string]interface{}
	if err := json.Unmarshal(raw, &document); err != nil {
		return nil, errors.Wrap(err, "reading the manifest's fields")
	}
	fields := make(map[string]interface{}, 2)
	for _, section := range []string{"spec", "metadata"} {
		if value, ok := document[section]; ok {
			fields[section] = value
		}
	}
	return fields, nil
}

// LookupKey is the key a reference's target is looked up by: the canonical
// field path (proto field names, bare indexes), with "status.outputs." dropped
// so an output reads by its flattened key ("backend_pool_ids.web") while a
// manifest field keeps its full path ("spec.name").
func LookupKey(target Target) string {
	return strings.TrimPrefix(refcheck.CanonicalValueFromPath(target.Kind, target.FieldPath), "status.outputs.")
}

// isManifestField reports whether a lookup key names a field the referenced
// manifest declares rather than one of its outputs.
func isManifestField(key string) bool {
	return strings.HasPrefix(key, "spec.") || strings.HasPrefix(key, "metadata.")
}
