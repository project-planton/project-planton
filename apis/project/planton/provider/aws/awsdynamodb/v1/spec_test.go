//go:build legacytests
// +build legacytests

package awsdynamodbv1_test

// The original api_test.go suite referenced an outdated version of the
// AwsDynamodbSpec proto definition that no longer exists.  It has been
// intentionally excluded from the default build so that the up-to-date
// validation tests in spec_test.go compile and run without conflict.
