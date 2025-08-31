package crkreflect

import (
	"github.com/pkg/errors"
	"strings"
)

// IdPrefixFromId extracts the id prefix from a cloud resource id
// For the new pattern: cr_<prefix>_<ulid>, it extracts <prefix>
// For legacy pattern: <prefix>_<rest>, it extracts <prefix>
func IdPrefixFromId(resourceId string) (string, error) {
	// Check if it follows the new cloud resource pattern
	if strings.HasPrefix(resourceId, "cr_") {
		// Pattern: cr_<prefix>_<ulid>
		parts := strings.SplitN(resourceId, "_", 3)
		if len(parts) < 3 || parts[1] == "" {
			return "", errors.Errorf("invalid cloud resource id format: %s", resourceId)
		}
		return parts[1], nil
	}

	// Legacy pattern: <prefix>_<rest>
	parts := strings.SplitN(resourceId, "_", 2)
	if len(parts) < 2 || parts[0] == "" {
		return "", errors.Errorf("failed to extract resource-id prefix from resource id: %s", resourceId)
	}
	return parts[0], nil
}
