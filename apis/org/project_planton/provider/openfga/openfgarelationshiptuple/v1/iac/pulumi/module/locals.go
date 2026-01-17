package module

import (
	openfgarelationshiptuplev1 "github.com/plantonhq/project-planton/apis/org/project_planton/provider/openfga/openfgarelationshiptuple/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals contains computed values derived from the stack input.
// This is a placeholder for consistency - OpenFGA has no Pulumi provider.
type Locals struct {
	// StoreId is the ID of the OpenFGA store this tuple belongs to.
	StoreId string
	// AuthorizationModelId is the optional ID of the authorization model.
	AuthorizationModelId string
	// User is the subject of the relationship tuple.
	User string
	// Relation is the relationship type.
	Relation string
	// Object is the resource being accessed.
	Object string
}

// initializeLocals creates and returns the computed local values from stack input.
// This is a placeholder for consistency - OpenFGA has no Pulumi provider.
func initializeLocals(_ *pulumi.Context, stackInput *openfgarelationshiptuplev1.OpenFgaRelationshipTupleStackInput) *Locals {
	target := stackInput.Target
	spec := target.Spec

	return &Locals{
		StoreId:              spec.StoreId,
		AuthorizationModelId: spec.AuthorizationModelId,
		User:                 spec.User,
		Relation:             spec.Relation,
		Object:               spec.Object,
	}
}
