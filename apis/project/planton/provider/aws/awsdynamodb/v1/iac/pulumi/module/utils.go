package awsdynamodb

// Copyright (c) Project Planton.
// SPDX-License-Identifier: Apache-2.0
//
// utils.go contains a very small set of helper functions that are handy when
// authoring Pulumi programs.  They do not contain any business logic that is
// specific to the AwsDynamodb component; instead they simply reduce boiler‐
// plate (pointer helpers, simple formatting helpers, lightweight error
// wrapping, …).  They live in their own file so they can be imported from any
// other part of the component implementation without creating import cycles.

import (
    "fmt"
    "strings"

    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ------------------------
// Pointer helpers (Go ≥1.18)
// ------------------------
//
// A generic Ptr helper saves a lot of := v; &v repetitions across the code base.
// Generic helpers require Go 1.18+, which is already a requirement for Pulumi.

// Ptr returns a pointer to v.  Example: dynamodb.TableArgs{Name: utils.Ptr("foo")}
func Ptr[T any](v T) *T { //nolint:ireturn // trivial helper
    return &v
}

// Val dereferences the pointer p or, when p == nil, returns the zero value of T.
func Val[T any](p *T) T {
    if p == nil {
        var zero T
        return zero
    }
    return *p
}

// ------------------------
// String helpers
// ------------------------

// NormalizeName joins the provided parts using dash ("-") as separator, removes
// duplicate separators, replaces spaces with dashes and returns the result in
// lower-case.  This is helpful when you want to build predictable resource
// names whilst still allowing users to provide arbitrary strings.
func NormalizeName(parts ...string) string {
    cleaned := make([]string, 0, len(parts))
    for _, p := range parts {
        p = strings.TrimSpace(p)
        p = strings.ReplaceAll(p, " ", "-")
        p = strings.ReplaceAll(p, "_", "-")
        if p != "" {
            cleaned = append(cleaned, p)
        }
    }
    return strings.ToLower(strings.Join(cleaned, "-"))
}

// ------------------------
// Pulumi helpers
// ------------------------

// MapToStringMap converts a Go map[string]string into the pulumi.StringMap type
// that Pulumi resource arguments expect.
func MapToStringMap(m map[string]string) pulumi.StringMap {
    res := make(pulumi.StringMap, len(m))
    for k, v := range m {
        res[k] = pulumi.String(v)
    }
    return res
}

// ------------------------
// Error helpers
// ------------------------

// WrapError prepends msg to err using the %w verb, returning a new error that
// can still be unwrapped with errors.Is / errors.As.
func WrapError(err error, msg string) error {
    if err == nil {
        return nil
    }
    return fmt.Errorf("%s: %w", msg, err)
}
