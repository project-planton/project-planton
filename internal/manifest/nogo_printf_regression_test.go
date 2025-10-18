package manifest

import (
	"fmt"
	"strings"
	"testing"
)

// TestNogoPrintfRegression ensures nogo catches non-constant format strings.
// This test exists as a regression guard for the nogo configuration.
// If this test compiles successfully, it indicates the nogo static analysis
// configuration is broken and no longer enforcing Go 1.24+ printf checks.
//
// Expected behavior: This test should FAIL during compilation with error:
// "non-constant format string in call to fmt.Sprintf"
//
// Background: Go 1.24 introduced stricter printf checking that requires
// format strings to be constant. The nogo framework must receive Go version
// metadata from rules_go to enable this check. This test verifies that
// version propagation is working correctly.
func TestNogoPrintfRegression(t *testing.T) {
	// This line should trigger a compilation error from nogo's printf analyzer
	// when the Go version metadata is correctly propagated (Go 1.24+).
	nonConstantFormat := "test format: %s"
	var msg strings.Builder
	msg.WriteString(fmt.Sprintf(nonConstantFormat)) // ERROR: non-constant format string
}

