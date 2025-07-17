/*
Package awsdynamodb centralises all custom error types that can be returned by the
Pulumi AWS-DynamoDB component resource.  They are kept in an isolated file so
that they can be reused by the various helpers produced during code-generation
(attribute-conversion helpers, validation helpers, SDK calls, …) while keeping
a single import path for callers.

The file purposefully avoids introducing any dependency on the Pulumi SDK so
that it can be imported from every internal package (including test packages)
without creating an import cycle.
*/

package awsdynamodb

import (
    "errors"
    "fmt"
    "strings"
)

// -----------------------------------------------------------------------------
// Sentinel errors
// -----------------------------------------------------------------------------

var (
    // ErrMissingRequiredField is returned when a spec or protobuf message is
    // missing a mandatory field.
    ErrMissingRequiredField = errors.New("missing required field")

    // ErrInvalidEnumValue is returned when an enum value is outside the range
    // of the allowed values defined in the protobuf specification.
    ErrInvalidEnumValue = errors.New("invalid enum value")

    // ErrUnsupportedConversion is returned when a conversion between two
    // representations cannot be performed (for example, from a protobuf enum to
    // the AWS SDK representation).
    ErrUnsupportedConversion = errors.New("unsupported conversion")
)

// -----------------------------------------------------------------------------
// ValidationError
// -----------------------------------------------------------------------------

// ValidationError represents a failure of a single validation rule against the
// user supplied specification.  It implements the standard error interface so
// it can be returned directly from provider entry points and will surface in
// Pulumi diagnostics.
//
// The structure holds enough context so that the caller can decide whether it
// should be displayed raw, or pretty-printed / aggregated together with other
// errors.
//
// NOTE: we purposefully do not embed proto-validation errors directly to avoid
// importing the generated code throughout the project (keeping import graphs
// minimal).
// -----------------------------------------------------------------------------

type ValidationError struct {
    // Field is the dotted-path location of the field that failed validation
    // (e.g. "ttl_specification.attribute_name").  When the error pertains to a
    // repeated element, callers are encouraged to suffix the path with the
    // index (e.g. "global_secondary_indexes[1].projection").
    Field string

    // Reason is a human-friendly explanation of why the value is invalid.
    Reason string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
    if e == nil {
        return "<nil>"
    }
    if e.Field == "" {
        return fmt.Sprintf("validation error: %s", e.Reason)
    }
    return fmt.Sprintf("validation error on %q: %s", e.Field, e.Reason)
}

// Is enables errors.Is(err, target) usage so that the caller can match
// *ValidationError against an arbitrary error value.
func (e *ValidationError) Is(target error) bool {
    _, ok := target.(*ValidationError)
    if ok {
        return true
    }
    return false
}

// NewValidationError is a small convenience wrapper that mirrors the built-in
// errors.New function but produces a *ValidationError.
func NewValidationError(field, reason string) *ValidationError {
    return &ValidationError{
        Field:  field,
        Reason: reason,
    }
}

// -----------------------------------------------------------------------------
// ConversionError
// -----------------------------------------------------------------------------

// ConversionError wraps failures that occur while converting the protobuf spec
// (or StackInputs / StackOutputs) into their AWS SDK representations and vice
// versa.
//
// Having a specific error allows the call-site to differentiate between user
// validation issues (ValidationError) and programmer / runtime issues
// (ConversionError) which may require a bug-fix instead of a user action.
// -----------------------------------------------------------------------------

type ConversionError struct {
    // From identifies the origin type or value (e.g. "proto.BillingMode").
    From string
    // To identifies the destination type or value (e.g. "types.BillingMode").
    To string
    // Reason explains why the conversion failed.
    Reason string
    // Err optionally wraps the original error.
    Err error
}

func (e *ConversionError) Error() string {
    b := strings.Builder{}
    b.WriteString("conversion error")
    if e.From != "" || e.To != "" {
        b.WriteString(fmt.Sprintf(" (%s → %s)", e.From, e.To))
    }
    if e.Reason != "" {
        b.WriteString(": ")
        b.WriteString(e.Reason)
    }
    if e.Err != nil {
        // Only add the underlying error if Reason did not already contain it.
        if e.Reason == "" {
            b.WriteString(": ")
            b.WriteString(e.Err.Error())
        } else {
            b.WriteString(" – ")
            b.WriteString(e.Err.Error())
        }
    }
    return b.String()
}

func (e *ConversionError) Unwrap() error { return e.Err }

// Is lets errors.Is detect a *ConversionError regardless of the instance.
func (e *ConversionError) Is(target error) bool {
    _, ok := target.(*ConversionError)
    if ok {
        return true
    }
    return false
}

// WrapConversionError constructs a *ConversionError while preserving the
// original error so that callers can use errors.Is / errors.As on the wrapped
// value as well.
func WrapConversionError(err error, from, to, reason string) *ConversionError {
    return &ConversionError{
        From:   from,
        To:     to,
        Reason: reason,
        Err:    err,
    }
}

// -----------------------------------------------------------------------------
// MultiError (utility)
// -----------------------------------------------------------------------------

// MultiError is a thin wrapper around a slice of errors that implements the
// error interface.  It is mainly used to aggregate multiple ValidationError or
// other independent issues discovered during pre-flight checks.
//
// By default, the Error() method concatenates the contained errors using a
// newline separator, which Pulumi will display nicely in its diagnostics.
//
// The zero value is ready to use.
// -----------------------------------------------------------------------------

type MultiError struct {
    Errors []error
}

// Append adds one or more errors to the collection.  nil values are ignored so
// callers can safely pass the result of a function that may return nil.
func (m *MultiError) Append(errs ...error) {
    for _, err := range errs {
        if err != nil {
            m.Errors = append(m.Errors, err)
        }
    }
}

// HasErrors returns true when the container holds at least one error.
func (m *MultiError) HasErrors() bool { return len(m.Errors) > 0 }

// Error implements the error interface.
func (m *MultiError) Error() string {
    if !m.HasErrors() {
        return "<nil>"
    }

    var sb strings.Builder
    for i, err := range m.Errors {
        if i > 0 {
            sb.WriteString("\n")
        }
        sb.WriteString(err.Error())
    }
    return sb.String()
}

// Unwrap enables errors.Unwrap / errors.Is / errors.As on the first underlying
// error (this behaviour matches fmt.Errorf with %w).  It is intentionally kept
// simple – callers that need the full list can type-assert to *MultiError.
func (m *MultiError) Unwrap() error {
    if len(m.Errors) == 0 {
        return nil
    }
    return m.Errors[0]
}

// -----------------------------------------------------------------------------
// Helper constructors
// -----------------------------------------------------------------------------

// NewMissingFieldError returns a *ValidationError pre-filled with
// ErrMissingRequiredField so that call-sites can create consistent messages.
func NewMissingFieldError(field string) *ValidationError {
    return &ValidationError{
        Field:  field,
        Reason: ErrMissingRequiredField.Error(),
    }
}

// NewInvalidEnumValueError returns a *ValidationError pre-filled with
// ErrInvalidEnumValue.
func NewInvalidEnumValueError(field string, value interface{}) *ValidationError {
    return &ValidationError{
        Field:  field,
        Reason: fmt.Sprintf("%s: %v", ErrInvalidEnumValue, value),
    }
}
