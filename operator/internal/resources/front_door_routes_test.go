package resources

import (
	"regexp"
	"slices"
	"strings"
	"testing"
)

// The handle-shaped first segments of the front door's routes, pinned. Organizations and people
// live at the root of the address (<host>/<organization>), so a name the front door routes to a
// backend before the console is a name no tenant may hold: the platform reserves each one
// (PlatformReservedHandles), and its guard (product/tools/ci/check_reserved_handles.py in
// planton-platform, OPERATOR_FRONT_DOOR_PREFIXES) pins this same list, because it cannot read this
// repository. A new prefix fails here until it is added there too.
var reservedFrontDoorSegments = []string{"idp", "rpc", "storage", "webhooks"}

// A segment a handle can never take (".well-known") needs no reservation.
var handleShape = regexp.MustCompile(`^[a-z0-9-]+$`)

func TestFrontDoorRoutesFirstSegmentsArePinnedAsReservedNames(t *testing.T) {
	var segments []string
	for _, route := range FrontDoorRoutes() {
		segment := strings.SplitN(strings.TrimPrefix(route.PathPrefix, "/"), "/", 2)[0]
		if segment == "" || !handleShape.MatchString(segment) || slices.Contains(segments, segment) {
			continue
		}
		segments = append(segments, segment)
	}
	slices.Sort(segments)
	if !slices.Equal(segments, reservedFrontDoorSegments) {
		t.Fatalf("the front door routes the first segments %v, but the pinned reserved names are %v: "+
			"update reservedFrontDoorSegments here, and the platform's reserved-handles guard "+
			"(OPERATOR_FRONT_DOOR_PREFIXES) and PlatformReservedHandles in planton-platform too",
			segments, reservedFrontDoorSegments)
	}
}
