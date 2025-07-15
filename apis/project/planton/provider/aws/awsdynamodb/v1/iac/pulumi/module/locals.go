package module

import (
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals bundles together all the in-module variables that are shared across
// resources. It follows the Project Planton conventions where the main spec is
// exposed as <Cloud><ResourceKind> (here AwsDynamodb) and common tags/labels
// are also kept in the struct.
//
// NOTE: additional transient/local fields can be appended later if the module
// needs to share derived data between helper functions.

type Locals struct {
	AwsDynamodb *awsdynamodbv1.AwsDynamodbSpec // Parsed spec from the stack input.
	Tags        map[string]string             // Consolidated/merged tags for AWS resources.
}

// initializeLocals extracts the spec and the user-supplied tags from the stack
// input. Extra default tags can be injected here if required by your
// organisation. The function keeps the implementation extremely simple so that
// it is easy to augment later without touching the call-sites.
func initializeLocals(input *awsdynamodbv1.AwsDynamodbStackInput) *Locals {
	var spec *awsdynamodbv1.AwsDynamodbSpec
	if input != nil {
		spec = input.GetSpec()
	}

	// Start with the tags coming from the spec (if any).
	tags := make(map[string]string)
	if spec != nil {
		for k, v := range spec.Tags {
			tags[k] = v
		}
	}

	return &Locals{
		AwsDynamodb: spec,
		Tags:        tags,
	}
}
