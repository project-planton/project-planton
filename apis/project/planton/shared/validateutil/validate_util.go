package validateutil

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	. "github.com/onsi/gomega"
)

const (
	StringInConstraint = "string.in"
)

type ExpectedViolation struct {
	FieldPath    string
	ConstraintId string
	Message      string
}

// Match is a helper to compare an actual violation
// against an expected one. This way, your test code is clean, and
// all the pointer-taking logic is in one place.
func Match(actual *validate.Violation, expected *ExpectedViolation) {
	// Convert the expected FieldPath to a pointer for a BeEquivalentTo check
	Expect(actual.FieldPath).To(BeEquivalentTo(&expected.FieldPath))

	// Same pattern for constraint ID
	var wantConstraintId = expected.ConstraintId
	Expect(actual.ConstraintId).To(BeEquivalentTo(&wantConstraintId))

	// And again for the violation message
	var wantMessage = expected.Message
	Expect(actual.Message).To(BeEquivalentTo(&wantMessage))
}
