package module

import (
	"github.com/google/uuid"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// guidNamespace scopes the GUIDs this module derives, so an assignment's GUID
// never equals one another module derives from the same identity.
var guidNamespace = uuid.MustParse("eb279405-325a-4537-8afc-3a1bd6b24fbc")

// stableGuid derives the assignment's GUID from its identity: the metadata id
// the platform assigns once and keeps, or -- for a deploy that carries none --
// the Pulumi stack that holds its state. The same identity derives the same
// GUID on every run, so an unchanged assignment plans no change.
func stableGuid(ctx *pulumi.Context, metadataId string) string {
	identity := metadataId
	if identity == "" {
		identity = ctx.Organization() + "/" + ctx.Project() + "/" + ctx.Stack()
	}
	return uuid.NewSHA1(guidNamespace, []byte(identity)).String()
}
