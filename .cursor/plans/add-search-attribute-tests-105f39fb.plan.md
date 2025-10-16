<!-- 105f39fb-786c-42cf-a03e-be11cbe24307 47e47479-0551-4c27-b310-9458372aeb74 -->
# Add Test Cases for Temporal Kubernetes Search Attribute Validation

## Context

The `TemporalKubernetesSearchAttribute` message in `spec.proto` (lines 24-40) now has a custom CEL validation on the `type` field:

```32:39:apis/project/planton/provider/kubernetes/workload/temporalkubernetes/v1/spec.proto
string type = 2 [
  (buf.validate.field).required = true,
  (buf.validate.field).cel = {
    id: "type.valid_search_attribute_type"
    message: "type must be one of: Keyword, Text, Int, Double, Bool, Datetime, KeywordList"
    expression: "this in ['Keyword', 'Text', 'Int', 'Double', 'Bool', 'Datetime', 'KeywordList']"
  }
];
```

This validation ensures that the `type` field only accepts the exact strings: `Keyword`, `Text`, `Int`, `Double`, `Bool`, `Datetime`, `KeywordList`.

## Test Cases to Add

Add to `apis/project/planton/provider/kubernetes/workload/temporalkubernetes/v1/api_test.go`:

### 1. Valid Search Attribute Types Test

Test all seven valid type values to ensure they pass validation.

### 2. Invalid Search Attribute Type Test

Test invalid values (lowercase, wrong values, empty) to ensure they fail validation with the correct error message.

### 3. Missing Type Field Test

Test that omitting the required `type` field triggers a validation error.

### 4. Multiple Search Attributes Test

Test that multiple search attributes with various valid types work correctly together.

## Implementation Details

The new test cases will follow the existing Ginkgo/Gomega pattern:

- Add a new `Describe` block for "Search Attribute Validation Tests"
- Use `BeforeEach` to set up a base valid configuration with database
- Add `Context` blocks for valid and invalid scenarios
- Use `protovalidate.Validate()` to verify validation behavior
- Check error is `nil` for valid cases
- Check error is `NotTo(gomega.BeNil())` for invalid cases

### To-dos

- [ ] Add test cases for all valid search attribute types (Keyword, Text, Int, Double, Bool, Datetime, KeywordList)
- [ ] Add test cases for invalid search attribute types (lowercase, wrong values, empty string)
- [ ] Add test case for missing required type field
- [ ] Add test case for multiple search attributes with mixed valid types