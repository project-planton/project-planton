# Spec Tests Authoring Guide

Purpose: create Go tests that exercise `spec.proto` validations (field options and CEL) for a resource kind.

## Scope
- File: `apis/project/planton/provider/<provider>/<kindfolder>/v1/spec_test.go`
- Package: `package <kindfolder>v1`
- Use Ginkgo v2 and Gomega with protovalidate-go.

## Imports
- `github.com/bufbuild/protovalidate-go`
- `. "github.com/onsi/ginkgo/v2"`
- `. "github.com/onsi/gomega"`
- Optionally: `github.com/project-planton/project-planton/apis/project/planton/shared/validateutil`

## Test Coverage
- Positive: minimal valid `<Kind>Spec` passes.
- Negative: key validations fail as expected:
  - Required strings (min_len)
  - Repeated min_items and uniqueness
  - Enum defined_only and non-zero enforcement
  - CEL conditionals (required-when, ordering constraints, mutually exclusive fields)

## Pattern
- Suite: `Test<Kind>Spec(t *testing.T)` with `RegisterFailHandler(Fail)` and `RunSpecs`.
- `Describe`/`Context`/`It` blocks for scenarios.
- Validate via `protovalidate.Validate(spec)`; assert nil/non-nil.
- For detailed assertions, consider `protovalidate.New().Validate()` and matching violations.

## Example Skeleton
```go
package awscloudfrontv1

import (
  "testing"
  "github.com/bufbuild/protovalidate-go"
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
)

func TestAwsCloudFrontSpec(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "AwsCloudFrontSpec Validation Suite")
}

var _ = Describe("AwsCloudFrontSpec validations", func() {
  It("accepts a valid spec", func() {
    // construct minimal valid spec then
    Expect(protovalidate.Validate(&AwsCloudFrontSpec{})).To(BeNil())
  })
})
```

## Notes
- Keep tests focused and robust; avoid brittle provider-format regexes unless necessary.
