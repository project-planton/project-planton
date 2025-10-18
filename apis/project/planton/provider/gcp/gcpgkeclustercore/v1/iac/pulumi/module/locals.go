package module

import (
	gcpgkeclustercorev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeclustercore/v1"
)

// Locals groups commonly accessed input data.
type Locals struct {
	GcpGkeClusterCore *gcpgkeclustercorev1.GcpGkeClusterCore
	ReleaseChannelStr string
}

// initializeLocals converts the raw stack input into a Locals struct.
func initializeLocals(stackInput *gcpgkeclustercorev1.GcpGkeClusterCoreStackInput) *Locals {
	l := &Locals{
		GcpGkeClusterCore: stackInput.Target,
	}

	// Map the proto enum to the literal string Google expects.
	switch l.GcpGkeClusterCore.Spec.GetReleaseChannel() {
	case gcpgkeclustercorev1.GkeReleaseChannel_RAPID:
		l.ReleaseChannelStr = "RAPID"
	case gcpgkeclustercorev1.GkeReleaseChannel_REGULAR:
		l.ReleaseChannelStr = "REGULAR"
	case gcpgkeclustercorev1.GkeReleaseChannel_STABLE:
		l.ReleaseChannelStr = "STABLE"
	case gcpgkeclustercorev1.GkeReleaseChannel_NONE:
		l.ReleaseChannelStr = "UNSPECIFIED"
	default:
		l.ReleaseChannelStr = "REGULAR" // sensible default
	}

	return l
}
